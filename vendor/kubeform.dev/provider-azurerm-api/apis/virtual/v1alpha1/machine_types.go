/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Machine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineSpec   `json:"spec,omitempty"`
	Status            MachineStatus `json:"status,omitempty"`
}

type MachineSpecAdditionalCapabilities struct {
	UltraSsdEnabled *bool `json:"ultraSsdEnabled" tf:"ultra_ssd_enabled"`
}

type MachineSpecBootDiagnostics struct {
	Enabled    *bool   `json:"enabled" tf:"enabled"`
	StorageURI *string `json:"storageURI" tf:"storage_uri"`
}

type MachineSpecIdentity struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	Type        *string `json:"type" tf:"type"`
}

type MachineSpecOsProfile struct {
	// +optional
	AdminPassword *string `json:"-" sensitive:"true" tf:"admin_password"`
	AdminUsername *string `json:"adminUsername" tf:"admin_username"`
	ComputerName  *string `json:"computerName" tf:"computer_name"`
	// +optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data"`
}

type MachineSpecOsProfileLinuxConfigSshKeys struct {
	KeyData *string `json:"keyData" tf:"key_data"`
	Path    *string `json:"path" tf:"path"`
}

type MachineSpecOsProfileLinuxConfig struct {
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication" tf:"disable_password_authentication"`
	// +optional
	SshKeys []MachineSpecOsProfileLinuxConfigSshKeys `json:"sshKeys,omitempty" tf:"ssh_keys"`
}

type MachineSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore *string `json:"certificateStore,omitempty" tf:"certificate_store"`
	CertificateURL   *string `json:"certificateURL" tf:"certificate_url"`
}

type MachineSpecOsProfileSecrets struct {
	SourceVaultID *string `json:"sourceVaultID" tf:"source_vault_id"`
	// +optional
	VaultCertificates []MachineSpecOsProfileSecretsVaultCertificates `json:"vaultCertificates,omitempty" tf:"vault_certificates"`
}

type MachineSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   *string `json:"component" tf:"component"`
	Content     *string `json:"-" sensitive:"true" tf:"content"`
	Pass        *string `json:"pass" tf:"pass"`
	SettingName *string `json:"settingName" tf:"setting_name"`
}

type MachineSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateURL *string `json:"certificateURL,omitempty" tf:"certificate_url"`
	Protocol       *string `json:"protocol" tf:"protocol"`
}

type MachineSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig []MachineSpecOsProfileWindowsConfigAdditionalUnattendConfig `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config"`
	// +optional
	EnableAutomaticUpgrades *bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades"`
	// +optional
	ProvisionVmAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
	// +optional
	Winrm []MachineSpecOsProfileWindowsConfigWinrm `json:"winrm,omitempty" tf:"winrm"`
}

type MachineSpecPlan struct {
	Name      *string `json:"name" tf:"name"`
	Product   *string `json:"product" tf:"product"`
	Publisher *string `json:"publisher" tf:"publisher"`
}

type MachineSpecStorageDataDisk struct {
	// +optional
	Caching      *string `json:"caching,omitempty" tf:"caching"`
	CreateOption *string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	Lun        *int64 `json:"lun" tf:"lun"`
	// +optional
	ManagedDiskID *string `json:"managedDiskID,omitempty" tf:"managed_disk_id"`
	// +optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`
	Name            *string `json:"name" tf:"name"`
	// +optional
	VhdURI *string `json:"vhdURI,omitempty" tf:"vhd_uri"`
	// +optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type MachineSpecStorageImageReference struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Offer *string `json:"offer,omitempty" tf:"offer"`
	// +optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher"`
	// +optional
	Sku *string `json:"sku,omitempty" tf:"sku"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type MachineSpecStorageOsDisk struct {
	// +optional
	Caching      *string `json:"caching,omitempty" tf:"caching"`
	CreateOption *string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// +optional
	ImageURI *string `json:"imageURI,omitempty" tf:"image_uri"`
	// +optional
	ManagedDiskID *string `json:"managedDiskID,omitempty" tf:"managed_disk_id"`
	// +optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`
	Name            *string `json:"name" tf:"name"`
	// +optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`
	// +optional
	VhdURI *string `json:"vhdURI,omitempty" tf:"vhd_uri"`
	// +optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type MachineSpec struct {
	State *MachineSpecResource `json:"state,omitempty" tf:"-"`

	Resource MachineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MachineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalCapabilities *MachineSpecAdditionalCapabilities `json:"additionalCapabilities,omitempty" tf:"additional_capabilities"`
	// +optional
	AvailabilitySetID *string `json:"availabilitySetID,omitempty" tf:"availability_set_id"`
	// +optional
	BootDiagnostics *MachineSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`
	// +optional
	DeleteDataDisksOnTermination *bool `json:"deleteDataDisksOnTermination,omitempty" tf:"delete_data_disks_on_termination"`
	// +optional
	DeleteOsDiskOnTermination *bool `json:"deleteOsDiskOnTermination,omitempty" tf:"delete_os_disk_on_termination"`
	// +optional
	Identity *MachineSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LicenseType         *string  `json:"licenseType,omitempty" tf:"license_type"`
	Location            *string  `json:"location" tf:"location"`
	Name                *string  `json:"name" tf:"name"`
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	// +optional
	OsProfile *MachineSpecOsProfile `json:"osProfile,omitempty" tf:"os_profile"`
	// +optional
	OsProfileLinuxConfig *MachineSpecOsProfileLinuxConfig `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config"`
	// +optional
	OsProfileSecrets []MachineSpecOsProfileSecrets `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets"`
	// +optional
	OsProfileWindowsConfig *MachineSpecOsProfileWindowsConfig `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config"`
	// +optional
	Plan *MachineSpecPlan `json:"plan,omitempty" tf:"plan"`
	// +optional
	PrimaryNetworkInterfaceID *string `json:"primaryNetworkInterfaceID,omitempty" tf:"primary_network_interface_id"`
	// +optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id"`
	ResourceGroupName         *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	StorageDataDisk []MachineSpecStorageDataDisk `json:"storageDataDisk,omitempty" tf:"storage_data_disk"`
	// +optional
	StorageImageReference *MachineSpecStorageImageReference `json:"storageImageReference,omitempty" tf:"storage_image_reference"`
	StorageOsDisk         *MachineSpecStorageOsDisk         `json:"storageOsDisk" tf:"storage_os_disk"`
	// +optional
	Tags   *map[string]string `json:"tags,omitempty" tf:"tags"`
	VmSize *string            `json:"vmSize" tf:"vm_size"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type MachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MachineList is a list of Machines
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Machine CRD objects
	Items []Machine `json:"items,omitempty"`
}
