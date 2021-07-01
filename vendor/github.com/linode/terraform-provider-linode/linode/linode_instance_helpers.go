package azurerm

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/azurerm/azurermgo"
	"golang.org/x/crypto/sha3"
)

var (
	boolFalse = false
	boolTrue  = true
)

type flattenedProfileReferrals map[string]interface{}

type diskSpec map[string]interface{}

// getDeadlineSeconds gets the seconds remaining until deadline is met.
func getDeadlineSeconds(ctx context.Context, d *schema.ResourceData) int {
	duration := d.Timeout(schema.TimeoutUpdate)
	if deadline, ok := ctx.Deadline(); ok {
		duration = time.Until(deadline)
	}
	return int(duration.Seconds())
}

func flattenProfileReferrals(referrals azurermgo.ProfileReferrals) []flattenedProfileReferrals {
	return []flattenedProfileReferrals{{
		"code":      referrals.Code,
		"url":       referrals.URL,
		"total":     referrals.Total,
		"completed": referrals.Completed,
		"pending":   referrals.Pending,
		"credit":    referrals.Credit,
	}}
}

func flattenInstanceSpecs(instance azurermgo.Instance) []map[string]int {
	return []map[string]int{{
		"vcpus":    instance.Specs.VCPUs,
		"disk":     instance.Specs.Disk,
		"memory":   instance.Specs.Memory,
		"transfer": instance.Specs.Transfer,
	}}
}

func flattenInstanceAlerts(instance azurermgo.Instance) []map[string]int {
	return []map[string]int{{
		"cpu":            instance.Alerts.CPU,
		"io":             instance.Alerts.IO,
		"network_in":     instance.Alerts.NetworkIn,
		"network_out":    instance.Alerts.NetworkOut,
		"transfer_quota": instance.Alerts.TransferQuota,
	}}
}

func flattenInstanceBackups(instance azurermgo.Instance) []map[string]interface{} {
	return []map[string]interface{}{{
		"enabled": instance.Backups.Enabled,
		"schedule": []map[string]interface{}{{
			"day":    instance.Backups.Schedule.Day,
			"window": instance.Backups.Schedule.Window,
		}},
	}}
}

func flattenInstanceDisks(instanceDisks []azurermgo.InstanceDisk) (disks []map[string]interface{}, swapSize int) {
	for _, disk := range instanceDisks {
		// Determine if swap exists and the size.  If it does not exist, swap_size=0
		if disk.Filesystem == "swap" {
			swapSize += disk.Size
		}
		disks = append(disks, map[string]interface{}{
			"id":         disk.ID,
			"size":       disk.Size,
			"label":      disk.Label,
			"filesystem": string(disk.Filesystem),
		})
	}
	return
}

func flattenInstanceConfigs(
	instanceConfigs []azurermgo.InstanceConfig,
	diskLabelIDMap map[int]string,
) (configs []map[string]interface{}) {
	for _, config := range instanceConfigs {

		devices := []map[string]interface{}{{
			"sda": flattenInstanceConfigDevice(config.Devices.SDA, diskLabelIDMap),
			"sdb": flattenInstanceConfigDevice(config.Devices.SDB, diskLabelIDMap),
			"sdc": flattenInstanceConfigDevice(config.Devices.SDC, diskLabelIDMap),
			"sdd": flattenInstanceConfigDevice(config.Devices.SDD, diskLabelIDMap),
			"sde": flattenInstanceConfigDevice(config.Devices.SDE, diskLabelIDMap),
			"sdf": flattenInstanceConfigDevice(config.Devices.SDF, diskLabelIDMap),
			"sdg": flattenInstanceConfigDevice(config.Devices.SDG, diskLabelIDMap),
			"sdh": flattenInstanceConfigDevice(config.Devices.SDH, diskLabelIDMap),
		}}

		interfaces := make([]interface{}, len(config.Interfaces))
		for i, ni := range config.Interfaces {
			interfaces[i] = flattenLinodeConfigInterface(ni)
		}

		// Determine if swap exists and the size.  If it does not exist, swap_size=0
		c := map[string]interface{}{
			"root_device":  config.RootDevice,
			"kernel":       config.Kernel,
			"run_level":    string(config.RunLevel),
			"virt_mode":    string(config.VirtMode),
			"comments":     config.Comments,
			"memory_limit": config.MemoryLimit,
			"label":        config.Label,
			"helpers": []map[string]bool{{
				"updatedb_disabled":  config.Helpers.UpdateDBDisabled,
				"distro":             config.Helpers.Distro,
				"modules_dep":        config.Helpers.ModulesDep,
				"network":            config.Helpers.Network,
				"devtmpfs_automount": config.Helpers.DevTmpFsAutomount,
			}},
			"devices":   devices,
			"interface": interfaces,
		}

		configs = append(configs, c)
	}
	return
}

func createInstanceConfigsFromSet(
	ctx context.Context,
	client azurermgo.Client,
	instanceID int,
	cset []interface{},
	diskIDLabelMap map[string]int,
	detacher volumeDetacher,
) (map[int]azurermgo.InstanceConfig, error) {
	configIDMap := make(map[int]azurermgo.InstanceConfig, len(cset))
	for _, v := range cset {
		config := v.(map[string]interface{})

		configOpts := azurermgo.InstanceConfigCreateOptions{}
		configOpts.Kernel = config["kernel"].(string)
		configOpts.Label = config["label"].(string)
		configOpts.Comments = config["comments"].(string)

		if helpers, ok := config["helpers"].([]interface{}); ok {
			for _, helper := range helpers {
				if helperMap, helperMapOk := helper.(map[string]interface{}); helperMapOk {
					configOpts.Helpers = &azurermgo.InstanceConfigHelpers{}
					if updateDBDisabled, ok := helperMap["updatedb_disabled"]; ok {
						configOpts.Helpers.UpdateDBDisabled = updateDBDisabled.(bool)
					}
					if distro, ok := helperMap["distro"]; ok {
						configOpts.Helpers.Distro = distro.(bool)
					}
					if modulesDep, ok := helperMap["modules_dep"]; ok {
						configOpts.Helpers.ModulesDep = modulesDep.(bool)
					}
					if network, ok := helperMap["network"]; ok {
						configOpts.Helpers.Network = network.(bool)
					}
					if devTmpFsAutomount, ok := helperMap["devtmpfs_automount"]; ok {
						configOpts.Helpers.DevTmpFsAutomount = devTmpFsAutomount.(bool)
					}
				}
			}
		}

		if interfaces, ok := config["interface"]; ok {
			interfaces := interfaces.([]interface{})
			configOpts.Interfaces = make([]azurermgo.InstanceConfigInterface, len(interfaces))

			for i, ni := range interfaces {
				configOpts.Interfaces[i] = expandLinodeConfigInterface(ni.(map[string]interface{}))
			}
		}

		rootDevice := config["root_device"].(string)
		if rootDevice != "" {
			configOpts.RootDevice = &rootDevice
		}
		// configOpts.InitRD = config["initrd"].(string)
		// TODO(displague) need a disk_label to initrd lookup?
		devices, ok := config["devices"].([]interface{})
		if !ok {
			return configIDMap, fmt.Errorf("Error converting config devices")
		}
		for _, device := range devices {
			deviceMap := device.(map[string]interface{})
			confDevices, err := expandInstanceConfigDeviceMap(deviceMap, diskIDLabelMap)
			if err != nil {
				return configIDMap, err
			}
			configOpts.Devices = *confDevices
		}

		if err := detachConfigVolumes(ctx, configOpts.Devices, detacher); err != nil {
			return configIDMap, err
		}

		instanceConfig, err := client.CreateInstanceConfig(ctx, instanceID, configOpts)
		if err != nil {
			return configIDMap, fmt.Errorf("Error creating Instance Config: %s", err)
		}
		configIDMap[instanceConfig.ID] = *instanceConfig
	}
	return configIDMap, nil
}

func updateInstanceConfigs(
	ctx context.Context,
	client azurermgo.Client,
	d *schema.ResourceData,
	instance azurermgo.Instance,
	tfConfigsOld, tfConfigsNew interface{},
	diskIDLabelMap map[string]int,
) (bool, map[string]int, []*azurermgo.InstanceConfig, error) {
	var updatedConfigMap map[string]int
	var rebootInstance bool
	var updatedConfigs []*azurermgo.InstanceConfig

	configs, err := client.ListInstanceConfigs(ctx, int(instance.ID), nil)
	if err != nil {
		return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf(
			"Error fetching the config for Instance %d: %s", instance.ID, err)
	}

	configMap := make(map[string]azurermgo.InstanceConfig, len(configs))
	for _, config := range configs {
		if _, duplicate := configMap[config.Label]; duplicate {
			return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf(
				"Error indexing Instance %d Configs: Label '%s' is assigned to multiple configs", instance.ID, config.Label)
		}
		configMap[config.Label] = config
	}

	oldConfigLabels := make([]string, len(tfConfigsOld.([]interface{})))

	for _, tfConfigOld := range tfConfigsOld.([]interface{}) {
		if oldConfig, ok := tfConfigOld.(map[string]interface{}); ok {
			oldConfigLabels = append(oldConfigLabels, oldConfig["label"].(string))
		}
	}
	tfConfigs := tfConfigsNew.([]interface{})
	updatedConfigs = make([]*azurermgo.InstanceConfig, len(tfConfigs))
	updatedConfigMap = make(map[string]int, len(tfConfigs))
	for _, tfConfig := range tfConfigs {
		tfc, _ := tfConfig.(map[string]interface{})
		label, _ := tfc["label"].(string)
		rootDevice, _ := tfc["root_device"].(string)
		if existingConfig, existing := configMap[label]; existing {
			configUpdateOpts := existingConfig.GetUpdateOptions()
			configUpdateOpts.Kernel = tfc["kernel"].(string)
			configUpdateOpts.RunLevel = tfc["run_level"].(string)
			configUpdateOpts.VirtMode = tfc["virt_mode"].(string)
			configUpdateOpts.RootDevice = rootDevice
			configUpdateOpts.Comments = tfc["comments"].(string)
			configUpdateOpts.MemoryLimit = tfc["memory_limit"].(int)

			tfcHelpersRaw, helpersFound := tfc["helpers"]
			if tfcHelpers, ok := tfcHelpersRaw.([]interface{}); helpersFound && ok {
				helpersMap := tfcHelpers[0].(map[string]interface{})
				configUpdateOpts.Helpers = &azurermgo.InstanceConfigHelpers{
					UpdateDBDisabled:  helpersMap["updatedb_disabled"].(bool),
					Distro:            helpersMap["distro"].(bool),
					ModulesDep:        helpersMap["modules_dep"].(bool),
					Network:           helpersMap["network"].(bool),
					DevTmpFsAutomount: helpersMap["devtmpfs_automount"].(bool),
				}

			}

			configUpdateOpts.Interfaces = make([]azurermgo.InstanceConfigInterface, 0)

			if interfaces, ok := tfc["interface"]; ok {
				interfaces := interfaces.([]interface{})

				configUpdateOpts.Interfaces = make([]azurermgo.InstanceConfigInterface, len(interfaces))

				for i, ni := range interfaces {
					configUpdateOpts.Interfaces[i] = expandLinodeConfigInterface(ni.(map[string]interface{}))
				}
			}

			tfcDevicesRaw, devicesFound := tfc["devices"]
			if tfcDevices, ok := tfcDevicesRaw.([]interface{}); devicesFound && ok {
				devices := tfcDevices[0].(map[string]interface{})

				configUpdateOpts.Devices, err = expandInstanceConfigDeviceMap(devices, diskIDLabelMap)

				if err != nil {
					return rebootInstance, updatedConfigMap, updatedConfigs, err
				}
				if configUpdateOpts.Devices != nil && emptyConfigDeviceMap(*configUpdateOpts.Devices) {
					configUpdateOpts.Devices = nil
				}
			} else {
				configUpdateOpts.Devices = nil
			}

			if configUpdateOpts.Devices != nil {
				detacher := makeVolumeDetacher(client, d)

				if detachErr := detachConfigVolumes(ctx, *configUpdateOpts.Devices, detacher); detachErr != nil {
					return rebootInstance, updatedConfigMap, updatedConfigs, detachErr
				}
			}

			updatedConfig, err := client.UpdateInstanceConfig(ctx, instance.ID, existingConfig.ID, configUpdateOpts)
			if err != nil {
				return rebootInstance, updatedConfigMap, updatedConfigs, fmt.Errorf(
					"Error updating Instance %d Config %d: %s", instance.ID, existingConfig.ID, err)
			}

			updatedConfigMap[updatedConfig.Label] = updatedConfig.ID
		} else {
			detacher := makeVolumeDetacher(client, d)

			configIDMap, err := createInstanceConfigsFromSet(
				ctx, client, instance.ID, []interface{}{tfc}, diskIDLabelMap, detacher)
			if err != nil {
				return rebootInstance, updatedConfigMap, updatedConfigs, err
			}
			for _, config := range configIDMap {
				updatedConfigMap[config.Label] = config.ID
				updatedConfigs = append(updatedConfigs, &config)
			}
		}
	}

	updatedConfigMap, err = deleteInstanceConfigs(ctx, client, instance.ID, oldConfigLabels, updatedConfigMap, configMap)
	if err != nil {
		return rebootInstance, updatedConfigMap, updatedConfigs, err
	}

	return rebootInstance, updatedConfigMap, updatedConfigs, nil
}

func deleteInstanceConfigs(
	ctx context.Context,
	client azurermgo.Client,
	instanceID int,
	oldConfigLabels []string,
	newConfigLabels map[string]int,
	configMap map[string]azurermgo.InstanceConfig,
) (map[string]int, error) {
	for _, oldLabel := range oldConfigLabels {
		if _, found := newConfigLabels[oldLabel]; !found {
			if listedConfig, found := configMap[oldLabel]; found {
				if err := client.DeleteInstanceConfig(ctx, instanceID, listedConfig.ID); err != nil {
					return newConfigLabels, err
				}
				delete(newConfigLabels, oldLabel)
			}
		}
	}
	return newConfigLabels, nil
}

func flattenInstanceConfigDevice(
	dev *azurermgo.InstanceConfigDevice, diskLabelIDMap map[int]string) []map[string]interface{} {
	if dev == nil || emptyInstanceConfigDevice(*dev) {
		return nil
	}

	if dev.DiskID > 0 {
		ret := map[string]interface{}{
			"disk_id": dev.DiskID,
		}
		if label, found := diskLabelIDMap[dev.DiskID]; found {
			ret["disk_label"] = label
		}
		return []map[string]interface{}{ret}
	}
	return []map[string]interface{}{{
		"volume_id": dev.VolumeID,
	}}
}

// expandInstanceConfigDeviceMap converts a terraform azurerm_instance config.*.devices map to a InstanceConfigDeviceMap
// for the Linode API.
func expandInstanceConfigDeviceMap(
	m map[string]interface{}, diskIDLabelMap map[string]int) (deviceMap *azurermgo.InstanceConfigDeviceMap, err error) {
	if len(m) == 0 {
		return nil, nil
	}
	deviceMap = &azurermgo.InstanceConfigDeviceMap{}
	for k, rdev := range m {
		devSlots := rdev.([]interface{})
		for _, rrdev := range devSlots {
			dev := rrdev.(map[string]interface{})
			tDevice := new(azurermgo.InstanceConfigDevice)
			if err := assignConfigDevice(tDevice, dev, diskIDLabelMap); err != nil {
				return nil, err
			}

			*deviceMap = changeInstanceConfigDevice(*deviceMap, k, tDevice)
		}
	}
	return deviceMap, nil
}

// changeInstanceConfigDevice returns a copy of a config device map with the specified disk slot changed to the
// provided device.
func changeInstanceConfigDevice(
	deviceMap azurermgo.InstanceConfigDeviceMap,
	namedSlot string,
	device *azurermgo.InstanceConfigDevice,
) azurermgo.InstanceConfigDeviceMap {
	tDevice := device
	if tDevice != nil && emptyInstanceConfigDevice(*tDevice) {
		tDevice = nil
	}
	switch namedSlot {
	case "sda":
		deviceMap.SDA = tDevice
	case "sdb":
		deviceMap.SDB = tDevice
	case "sdc":
		deviceMap.SDC = tDevice
	case "sdd":
		deviceMap.SDD = tDevice
	case "sde":
		deviceMap.SDE = tDevice
	case "sdf":
		deviceMap.SDF = tDevice
	case "sdg":
		deviceMap.SDG = tDevice
	case "sdh":
		deviceMap.SDH = tDevice
	}

	return deviceMap
}

// emptyInstanceConfigDevice returns true only when neither the disk or volume have been assigned to a config device.
func emptyInstanceConfigDevice(dev azurermgo.InstanceConfigDevice) bool {
	return (dev.DiskID == 0 && dev.VolumeID == 0)
}

// emptyConfigDeviceMap returns true only when none of the disks in a config device map have been assigned.
func emptyConfigDeviceMap(dmap azurermgo.InstanceConfigDeviceMap) bool {
	drives := []*azurermgo.InstanceConfigDevice{
		dmap.SDA, dmap.SDB, dmap.SDC, dmap.SDD, dmap.SDE, dmap.SDF, dmap.SDG, dmap.SDH,
	}
	empty := true
	for _, drive := range drives {
		if drive != nil && !emptyInstanceConfigDevice(*drive) {
			empty = false
			break
		}
	}
	return empty
}

type volumeDetacher func(context.Context, int, string) error

func makeVolumeDetacher(client azurermgo.Client, d *schema.ResourceData) volumeDetacher {
	return func(ctx context.Context, volumeID int, reason string) error {
		log.Printf("[INFO] Detaching Linode Volume %d %s", volumeID, reason)
		if err := client.DetachVolume(ctx, volumeID); err != nil {
			return err
		}

		log.Printf("[INFO] Waiting for Linode Volume %d to detach ...", volumeID)
		if _, err := client.WaitForVolumeLinodeID(ctx, volumeID, nil, getDeadlineSeconds(ctx, d)); err != nil {
			return err
		}
		return nil
	}
}

func expandInstanceConfigDevice(m map[string]interface{}) *azurermgo.InstanceConfigDevice {
	var dev *azurermgo.InstanceConfigDevice
	// be careful of `disk_label string` in m
	if diskID, ok := m["disk_id"]; ok && diskID.(int) > 0 {
		dev = &azurermgo.InstanceConfigDevice{
			DiskID: diskID.(int),
		}
	} else if volumeID, ok := m["volume_id"]; ok && volumeID.(int) > 0 {
		dev = &azurermgo.InstanceConfigDevice{
			VolumeID: m["volume_id"].(int),
		}
	}
	return dev
}

func createInstanceDisk(
	ctx context.Context,
	client azurermgo.Client,
	instance azurermgo.Instance,
	disk diskSpec,
	d *schema.ResourceData,
) (*azurermgo.InstanceDisk, error) {
	diskOpts := azurermgo.InstanceDiskCreateOptions{
		Label:      disk["label"].(string),
		Filesystem: disk["filesystem"].(string),
		Size:       disk["size"].(int),
	}

	if image, ok := disk["image"]; ok {
		diskOpts.Image = image.(string)

		if rootPass, ok := disk["root_pass"]; ok && rootPass != "" {
			diskOpts.RootPass = rootPass.(string)
		} else {
			var err error
			diskOpts.RootPass, err = createRandomRootPassword()
			if err != nil {
				return nil, err
			}
		}

		if authorizedKeys, ok := disk["authorized_keys"]; ok {
			for _, sshKey := range authorizedKeys.([]interface{}) {
				diskOpts.AuthorizedKeys = append(diskOpts.AuthorizedKeys, sshKey.(string))
			}
		}

		if authorizedUsers, ok := disk["authorized_users"]; ok {
			for _, user := range authorizedUsers.([]interface{}) {
				diskOpts.AuthorizedUsers = append(diskOpts.AuthorizedUsers, user.(string))
			}
		}

		if stackscriptID, ok := disk["stackscript_id"]; ok {
			diskOpts.StackscriptID = stackscriptID.(int)
		}

		if stackscriptDataRaw, ok := disk["stackscript_data"]; ok {
			stackscriptData, ok := stackscriptDataRaw.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("Error parsing stackscript_data: expected map[string]interface{}")
			}
			diskOpts.StackscriptData = make(map[string]string, len(stackscriptData))
			for name, value := range stackscriptData {
				diskOpts.StackscriptData[name] = value.(string)
			}
		}
	}

	instanceDisk, err := client.CreateInstanceDisk(ctx, instance.ID, diskOpts)
	if err != nil {
		return nil, fmt.Errorf("Error creating Linode instance %d disk: %s", instance.ID, err)
	}

	_, err = client.WaitForEventFinished(ctx, instance.ID, azurermgo.EntityLinode,
		azurermgo.ActionDiskCreate, *instanceDisk.Created, getDeadlineSeconds(ctx, d))
	if err != nil {
		return nil, fmt.Errorf("Error waiting for Linode instance %d disk: %s", instanceDisk.ID, err)
	}

	return instanceDisk, err
}

// getInstanceDisks returns a map of disks for a given instance that is indexed by label.
func getInstanceDisks(
	ctx context.Context, client azurermgo.Client, instanceID int) (map[string]azurermgo.InstanceDisk, error) {
	disks, err := client.ListInstanceDisks(ctx, instanceID, nil)
	if err != nil {
		return nil, fmt.Errorf("Error fetching the disks for Instance %d: %s", instanceID, err)
	}

	diskMap := make(map[string]azurermgo.InstanceDisk)
	for _, disk := range disks {
		if _, duplicate := diskMap[disk.Label]; duplicate {
			return nil, fmt.Errorf("Error indexing Instance %d Disks: Label '%s' is assigned to multiple disks",
				instanceID, disk.Label)
		}
		diskMap[disk.Label] = disk
	}
	return diskMap, nil
}

// getInstanceDiskLabelIDMap returns a map of an instances disk labels to their corresponding IDs.
func getInstanceDiskLabelIDMap(
	ctx context.Context, client azurermgo.Client, d *schema.ResourceData, instanceID int) (map[string]int, error) {
	diskSpecs := d.Get("disk").([]interface{})
	disks, err := getInstanceDisks(ctx, client, instanceID)
	if err != nil {
		return nil, err
	}

	labelIDMap := make(map[string]int)
	for _, spec := range diskSpecs {
		label := spec.(map[string]interface{})["label"].(string)
		disk, ok := disks[label]
		if !ok {
			return nil, fmt.Errorf(`could not map disk label "%s" to an ID; not found`, label)
		}
		labelIDMap[label] = disk.ID
	}
	return labelIDMap, nil
}

// getInstanceDiskSpecChange returns a map of disk specs indexed by label.
func getInstanceDiskSpecChange(d *schema.ResourceData) (oldDiskSpecs, newDiskSpecs map[string]diskSpec) {
	old, new := d.GetChange("disk")
	oldDisk := old.([]interface{})
	newDisk := new.([]interface{})
	oldDiskSpecs = make(map[string]diskSpec)
	newDiskSpecs = make(map[string]diskSpec)

	for _, spec := range oldDisk {
		spec := spec.(map[string]interface{})
		oldDiskSpecs[spec["label"].(string)] = spec
	}
	for _, spec := range newDisk {
		spec := spec.(map[string]interface{})
		newDiskSpecs[spec["label"].(string)] = spec
	}
	return
}

// getInstanceDiskSpecDiffs sorts the disk specs by added, removed, and existing.
func getInstanceDiskSpecDiffs(
	oldDiskSpecs, newDiskSpecs map[string]diskSpec) (added, removed, existing map[string]diskSpec) {
	added = make(map[string]diskSpec)
	removed = make(map[string]diskSpec)
	existing = make(map[string]diskSpec)

	placed := make(map[string]struct{})
	for label, spec := range newDiskSpecs {
		_, exists := oldDiskSpecs[label]
		if exists {
			existing[label] = spec
		} else {
			added[label] = spec
		}
		placed[label] = struct{}{}
	}

	for label, spec := range oldDiskSpecs {
		if _, ok := placed[label]; !ok {
			removed[label] = spec
		}
	}
	return
}

// updateInstanceDisks ensures the disk specification matches the instance disk state. This means creating,
// updating, and deleting disks as needed.
//
// This function will also warn when there are disks attached to an instance which are not managed by
// terraform.
func updateInstanceDisks(
	ctx context.Context, client azurermgo.Client, d *schema.ResourceData, instance azurermgo.Instance) (bool, error) {
	oldDisk, newDisk := getInstanceDiskSpecChange(d)
	added, removed, existing := getInstanceDiskSpecDiffs(oldDisk, newDisk)
	disks, err := getInstanceDisks(ctx, client, instance.ID)
	if err != nil {
		return false, err
	}
	hasChanges := len(added)+len(removed) > 0

	// ensure all old disks which have not been removed are present on the instance
	for label := range oldDisk {
		_, wasRemoved := removed[label]
		if _, ok := disks[label]; !ok && !wasRemoved {
			return hasChanges, fmt.Errorf(`disk "%s" was not found on instance %d`, label, instance.ID)
		}
	}
	// keep track of all disks visited for accounting
	visited := make(map[string]struct{})

	// remove disks staged for removal
	for label := range removed {
		disk, ok := disks[label]
		if !ok {
			// It's ok if a removed disk is not found
			continue
		}
		if err := client.DeleteInstanceDisk(ctx, instance.ID, disk.ID); err != nil {
			return hasChanges, err
		}
		_, err = client.WaitForEventFinished(ctx, instance.ID, azurermgo.EntityLinode,
			azurermgo.ActionDiskDelete, *instance.Created, getDeadlineSeconds(ctx, d))
		if err != nil {
			return hasChanges, fmt.Errorf(
				"error waiting for Instance %d Disk %d to finish deleting: %s", instance.ID, disk.ID, err)
		}
		visited[label] = struct{}{}
	}

	// ensure state is consistent with existing disks specs
	for label, spec := range existing {
		existingDisk := disks[label]
		// The only non-destructive change supported is resize.
		// Label renames are not supported because this TF provider relies on the label as an identifier.
		if spec["size"].(int) != existingDisk.Size {
			if err := changeInstanceDiskSize(ctx, &client, instance, existingDisk, spec["size"].(int), d); err != nil {
				return hasChanges, err
			}
			hasChanges = true
		}
		if spec["filesystem"].(string) != string(existingDisk.Filesystem) {
			return hasChanges, fmt.Errorf("failed to update disk %d; filesystems can not be changed", existingDisk.ID)
		}
		visited[label] = struct{}{}
	}

	// create disks staged for creation
	for _, spec := range added {
		if _, err := createInstanceDisk(ctx, client, instance, spec, d); err != nil {
			return hasChanges, err
		}
	}

	// ensure all disks visited
	for label, disk := range disks {
		if _, ok := visited[label]; !ok {
			// warn if disk found that is not in terraform state
			fmt.Printf("[WARN] found disk %s (%d) on instance %d not found in state", label, disk.ID, instance.ID)
		}
	}

	return hasChanges, nil
}

// sshKeyState hashes a string passed in as an interface.
func sshKeyState(val interface{}) string {
	return hashString(strings.Join(val.([]string), "\n"))
}

// rootPasswordState hashes a string passed in as an interface.
func rootPasswordState(val interface{}) string {
	return hashString(val.(string))
}

// hashString hashes a string.
func hashString(key string) string {
	hash := sha3.Sum512([]byte(key))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func createRandomRootPassword() (string, error) {
	rawRootPass := make([]byte, 50)
	_, err := rand.Read(rawRootPass)
	if err != nil {
		return "", fmt.Errorf("Failed to generate random password")
	}
	rootPass := base64.StdEncoding.EncodeToString(rawRootPass)
	return rootPass, nil
}

// ensureInstanceOffline ensures that a given instance is offline.
func ensureInstanceOffline(
	ctx context.Context, client *azurermgo.Client, instanceID, timeout int) (instance *azurermgo.Instance, err error) {
	if instance, err = client.GetInstance(ctx, instanceID); err != nil {
		return
	}

	if instance.Status == azurermgo.InstanceOffline {
		return
	} else if instance.Status != azurermgo.InstanceShuttingDown {
		err = client.ShutdownInstance(ctx, instanceID)
	}

	if err != nil {
		return
	}
	return client.WaitForInstanceStatus(ctx, instanceID, azurermgo.InstanceOffline, timeout)
}

// changeInstanceType resizes the Linode Instance.
func changeInstanceType(
	ctx context.Context,
	client *azurermgo.Client,
	instanceID int,
	targetType string,
	d *schema.ResourceData,
) (*azurermgo.Instance, error) {
	instance, err := ensureInstanceOffline(ctx, client, instanceID, getDeadlineSeconds(ctx, d))
	if err != nil {
		return nil, err
	}

	diskResize := false
	resizeOpts := azurermgo.InstanceResizeOptions{
		AllowAutoDiskResize: &diskResize,
		Type:                targetType,
	}

	if err := client.ResizeInstance(ctx, instance.ID, resizeOpts); err != nil {
		return nil, fmt.Errorf("Error resizing Instance %d: %s", instance.ID, err)
	}
	_, err = client.WaitForEventFinished(ctx, instance.ID, azurermgo.EntityLinode, azurermgo.ActionLinodeResize,
		*instance.Created, getDeadlineSeconds(ctx, d))
	if err != nil {
		return nil, fmt.Errorf("Error waiting for instance %d to finish resizing: %s", instance.ID, err)
	}

	// Wait for instance status to go back to idle, offline state
	if instance, err = client.WaitForInstanceStatus(
		ctx, instance.ID, azurermgo.InstanceOffline, getDeadlineSeconds(ctx, d),
	); err != nil {
		return nil, fmt.Errorf("Error waiting for Instance %d to enter offline state: %s", instance.ID, err)
	}
	return instance, nil
}

// returns the amount of disk space used by the new plan and old plan.
func getDiskSizeChange(oldDisk interface{}, newDisk interface{}) (int, int) {
	tfDisksOldInterface := oldDisk.([]interface{})
	tfDisksNewInterface := newDisk.([]interface{})

	oldDiskSize := 0
	newDiskSize := 0
	// Get total amount of disk usage before and after
	for _, disk := range tfDisksOldInterface {
		oldDiskSize += disk.(map[string]interface{})["size"].(int)
	}

	for _, disk := range tfDisksNewInterface {
		newDiskSize += disk.(map[string]interface{})["size"].(int)
	}
	return oldDiskSize, newDiskSize
}

func changeInstanceDiskSize(
	ctx context.Context,
	client *azurermgo.Client,
	instance azurermgo.Instance,
	disk azurermgo.InstanceDisk,
	targetSize int,
	d *schema.ResourceData,
) error {
	if instance.Specs.Disk < targetSize {
		return fmt.Errorf("Error resizing disk %d: size exceeds disk size for Instance %d", disk.ID, instance.ID)
	}

	switch instance.Status {
	case azurermgo.InstanceShuttingDown:
		if _, err := client.WaitForInstanceStatus(
			ctx, instance.ID, azurermgo.InstanceOffline, getDeadlineSeconds(ctx, d),
		); err != nil {
			return fmt.Errorf("Error waiting for Instance %d to go offline: %s", instance.ID, err)
		}
	case azurermgo.InstanceOffline:
	default:
		if err := client.ShutdownInstance(ctx, instance.ID); err != nil {
			return err
		}
	}

	// Wait for instance to go offline. Resize the disk once Linode is shut down.
	if _, err := client.WaitForInstanceStatus(
		ctx, instance.ID, azurermgo.InstanceOffline, getDeadlineSeconds(ctx, d),
	); err != nil {
		return fmt.Errorf("Error waiting for Instance %d to go offline: %s", instance.ID, err)
	}

	if err := client.ResizeInstanceDisk(ctx, instance.ID, disk.ID, targetSize); err != nil {
		return fmt.Errorf("Error resizing disk %d for Instance %d: %s", disk.ID, instance.ID, err)
	}

	// Wait for the disk resize operation to complete, and boot instance.
	_, err := client.WaitForEventFinished(ctx, instance.ID, azurermgo.EntityLinode, azurermgo.ActionDiskResize,
		*disk.Updated, getDeadlineSeconds(ctx, d))
	if err != nil {
		return fmt.Errorf("Error waiting for resize of Instance %d Disk %d: %s", instance.ID, disk.ID, err)
	}

	// Check to see if the resize operation worked
	if updatedDisk, err := client.WaitForInstanceDiskStatus(ctx, instance.ID, disk.ID, azurermgo.DiskReady,
		getDeadlineSeconds(ctx, d)); err != nil {
		return fmt.Errorf("Error waiting disk %d on instance %d to be ready: %s", disk.ID, instance.ID, err)
	} else if updatedDisk.Size != targetSize {
		return fmt.Errorf(
			"Error resizing disk %d on instance %d from %d to %d", disk.ID, instance.ID, disk.Size, targetSize)
	}

	return nil
}

// privateIP determines if an IP is for private use (RFC1918)
// https://stackoverflow.com/a/41273687
func privateIP(ip net.IP) bool {
	_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
	_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
	_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
	private := private24BitBlock.Contains(ip) || private20BitBlock.Contains(ip) || private16BitBlock.Contains(ip)
	return private
}

func assignConfigDevice(
	device *azurermgo.InstanceConfigDevice, dev map[string]interface{}, diskIDLabelMap map[string]int) error {
	if label, ok := dev["disk_label"].(string); ok && len(label) > 0 {
		if dev["disk_id"], ok = diskIDLabelMap[label]; !ok {
			return fmt.Errorf(`Error mapping disk label "%s" to ID`, dev["disk_label"])
		}
	}
	expanded := expandInstanceConfigDevice(dev)
	if expanded != nil {
		*device = *expanded
	}
	return nil
}

// getInstanceDefaultDisks gets the boot and swap disk for an instance which has implicit, default disks.
func getInstanceDefaultDisks(
	ctx context.Context,
	instanceID int,
	client *azurermgo.Client,
) (bootDisk, swapDisk *azurermgo.InstanceDisk, err error) {
	var disks []azurermgo.InstanceDisk
	disks, err = client.ListInstanceDisks(ctx, instanceID, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Error getting instance managed disks: %s", err)
	}
	bootDisk = findDiskByFS(disks, azurermgo.FilesystemExt4)
	swapDisk = findDiskByFS(disks, azurermgo.FilesystemSwap)
	return
}

// getInstanceTypeChange checks to see if the azurerm itself was resized.
func getInstanceTypeChange(
	ctx context.Context,
	d *schema.ResourceData,
	client *azurermgo.Client,
) (oldSpec, newSpec *azurermgo.LinodeType, err error) {
	old, new := d.GetChange("type")
	oldType, newType := old.(string), new.(string)

	oldSpec, err = client.GetType(ctx, oldType)
	if err != nil {
		return
	}
	newSpec, err = client.GetType(ctx, newType)
	if err != nil {
		return
	}
	return
}

// applyInstanceDiskSpec checks to see if the staged disk changes can be supported by the instance specification's
// capacity. If there is sufficient space, it attempts to update the disks.
//
// returns bool describing whether change has occurred.
func applyInstanceDiskSpec(
	ctx context.Context,
	d *schema.ResourceData,
	client *azurermgo.Client,
	instance *azurermgo.Instance,
	typ *azurermgo.LinodeType,
) (bool, error) {
	if err := assertDiskConfigFitsInstanceType(d, typ); err != nil {
		return false, err
	}
	return updateInstanceDisks(ctx, *client, d, *instance)
}

// assertDiskConfigFitsInstanceType asserts that the cumulative disk space used by a given disk config fits a given
// azurerm type spec for disk capacity.
func assertDiskConfigFitsInstanceType(d *schema.ResourceData, typ *azurermgo.LinodeType) error {
	oldDisks, newDisks := d.GetChange("disk")
	_, newDiskSize := getDiskSizeChange(oldDisks, newDisks)
	if typ.Disk < newDiskSize {
		return fmt.Errorf(
			"Linode type %s has insufficient disk capacity for the config. Have %d; want %d",
			typ.Label, typ.Disk, newDiskSize)
	}
	return nil
}

// applyInstanceTypeChange checks to see if the staged disk changes can be supported by the new instance
// specification. If there is sufficient space, it attempts to update the instance type.
func applyInstanceTypeChange(
	ctx context.Context,
	d *schema.ResourceData,
	client *azurermgo.Client,
	instance *azurermgo.Instance,
	typ *azurermgo.LinodeType,
) (*azurermgo.Instance, error) {
	if err := assertDiskConfigFitsInstanceType(d, typ); err != nil {
		return nil, err
	}
	return changeInstanceType(ctx, client, instance.ID, typ.ID, d)
}

// detachConfigVolumes detaches any volumes associated with an InstanceConfig.Devices struct.
func detachConfigVolumes(
	ctx context.Context, dmap azurermgo.InstanceConfigDeviceMap, detacher volumeDetacher) error {
	// Preallocate our slice of config devices
	drives := []*azurermgo.InstanceConfigDevice{
		dmap.SDA, dmap.SDB, dmap.SDC, dmap.SDD, dmap.SDE, dmap.SDF, dmap.SDG, dmap.SDH,
	}

	// Make a buffered error channel for our goroutines to send error values back on
	errCh := make(chan error, len(drives))

	// Make a sync.WaitGroup so our devices can signal they're finished
	var wg sync.WaitGroup
	wg.Add(len(drives))

	// For each drive, spawn a goroutine to detach the volume, send an error on the err channel
	// if one exists, and signal the worker process is done
	for _, d := range drives {
		go func(dev *azurermgo.InstanceConfigDevice) {
			defer wg.Done()

			if dev != nil && dev.VolumeID > 0 {
				err := detacher(ctx, dev.VolumeID, "for config attachment")
				if err != nil {
					errCh <- err
				}
			}
		}(d)
	}

	// Wait until all processes are finished and close the error channel so we can range over it
	wg.Wait()
	close(errCh)

	// Build the error from the errors in the channel and return the combined error if any exist
	var errStr string
	for err := range errCh {
		if len(errStr) == 0 {
			errStr += ", "
		}

		errStr += err.Error()
	}

	if len(errStr) > 0 {
		return fmt.Errorf("Error detaching volumes: %s", errStr)
	}

	return nil
}

func expandLinodeConfigInterface(i map[string]interface{}) azurermgo.InstanceConfigInterface {
	result := azurermgo.InstanceConfigInterface{}

	result.Label = i["label"].(string)
	result.Purpose = azurermgo.ConfigInterfacePurpose(i["purpose"].(string))
	result.IPAMAddress = i["ipam_address"].(string)

	return result
}

func flattenLinodeConfigInterface(i azurermgo.InstanceConfigInterface) map[string]interface{} {
	result := make(map[string]interface{})

	result["label"] = i.Label
	result["purpose"] = i.Purpose
	result["ipam_address"] = i.IPAMAddress

	return result
}
