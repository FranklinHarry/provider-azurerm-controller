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

type MachineScaleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineScaleSetSpec   `json:"spec,omitempty"`
	Status            MachineScaleSetStatus `json:"status,omitempty"`
}

type MachineScaleSetSpecBootDiagnostics struct {
	// +optional
	Enabled    *bool   `json:"enabled,omitempty" tf:"enabled"`
	StorageURI *string `json:"storageURI" tf:"storage_uri"`
}

type MachineScaleSetSpecExtension struct {
	// +optional
	AutoUpgradeMinorVersion *bool   `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version"`
	Name                    *string `json:"name" tf:"name"`
	// +optional
	ProtectedSettings *string `json:"-" sensitive:"true" tf:"protected_settings"`
	// +optional
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions"`
	Publisher                *string  `json:"publisher" tf:"publisher"`
	// +optional
	Settings           *string `json:"settings,omitempty" tf:"settings"`
	Type               *string `json:"type" tf:"type"`
	TypeHandlerVersion *string `json:"typeHandlerVersion" tf:"type_handler_version"`
}

type MachineScaleSetSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	Type        *string `json:"type" tf:"type"`
}

type MachineScaleSetSpecNetworkProfileDnsSettings struct {
	DnsServers []string `json:"dnsServers" tf:"dns_servers"`
}

type MachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration struct {
	DomainNameLabel *string `json:"domainNameLabel" tf:"domain_name_label"`
	IdleTimeout     *int64  `json:"idleTimeout" tf:"idle_timeout"`
	Name            *string `json:"name" tf:"name"`
}

type MachineScaleSetSpecNetworkProfileIpConfiguration struct {
	// +optional
	ApplicationGatewayBackendAddressPoolIDS []string `json:"applicationGatewayBackendAddressPoolIDS,omitempty" tf:"application_gateway_backend_address_pool_ids"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ApplicationSecurityGroupIDS []string `json:"applicationSecurityGroupIDS,omitempty" tf:"application_security_group_ids"`
	// +optional
	LoadBalancerBackendAddressPoolIDS []string `json:"loadBalancerBackendAddressPoolIDS,omitempty" tf:"load_balancer_backend_address_pool_ids"`
	// +optional
	LoadBalancerInboundNATRulesIDS []string `json:"loadBalancerInboundNATRulesIDS,omitempty" tf:"load_balancer_inbound_nat_rules_ids"`
	Name                           *string  `json:"name" tf:"name"`
	Primary                        *bool    `json:"primary" tf:"primary"`
	// +optional
	PublicIPAddressConfiguration *MachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration `json:"publicIPAddressConfiguration,omitempty" tf:"public_ip_address_configuration"`
	SubnetID                     *string                                                                       `json:"subnetID" tf:"subnet_id"`
}

type MachineScaleSetSpecNetworkProfile struct {
	// +optional
	AcceleratedNetworking *bool `json:"acceleratedNetworking,omitempty" tf:"accelerated_networking"`
	// +optional
	DnsSettings     *MachineScaleSetSpecNetworkProfileDnsSettings      `json:"dnsSettings,omitempty" tf:"dns_settings"`
	IpConfiguration []MachineScaleSetSpecNetworkProfileIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	// +optional
	IpForwarding *bool   `json:"ipForwarding,omitempty" tf:"ip_forwarding"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id"`
	Primary                *bool   `json:"primary" tf:"primary"`
}

type MachineScaleSetSpecOsProfile struct {
	// +optional
	AdminPassword      *string `json:"-" sensitive:"true" tf:"admin_password"`
	AdminUsername      *string `json:"adminUsername" tf:"admin_username"`
	ComputerNamePrefix *string `json:"computerNamePrefix" tf:"computer_name_prefix"`
	// +optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data"`
}

type MachineScaleSetSpecOsProfileLinuxConfigSshKeys struct {
	// +optional
	KeyData *string `json:"keyData,omitempty" tf:"key_data"`
	Path    *string `json:"path" tf:"path"`
}

type MachineScaleSetSpecOsProfileLinuxConfig struct {
	// +optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication"`
	// +optional
	SshKeys []MachineScaleSetSpecOsProfileLinuxConfigSshKeys `json:"sshKeys,omitempty" tf:"ssh_keys"`
}

type MachineScaleSetSpecOsProfileSecretsVaultCertificates struct {
	// +optional
	CertificateStore *string `json:"certificateStore,omitempty" tf:"certificate_store"`
	CertificateURL   *string `json:"certificateURL" tf:"certificate_url"`
}

type MachineScaleSetSpecOsProfileSecrets struct {
	SourceVaultID *string `json:"sourceVaultID" tf:"source_vault_id"`
	// +optional
	VaultCertificates []MachineScaleSetSpecOsProfileSecretsVaultCertificates `json:"vaultCertificates,omitempty" tf:"vault_certificates"`
}

type MachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig struct {
	Component   *string `json:"component" tf:"component"`
	Content     *string `json:"-" sensitive:"true" tf:"content"`
	Pass        *string `json:"pass" tf:"pass"`
	SettingName *string `json:"settingName" tf:"setting_name"`
}

type MachineScaleSetSpecOsProfileWindowsConfigWinrm struct {
	// +optional
	CertificateURL *string `json:"certificateURL,omitempty" tf:"certificate_url"`
	Protocol       *string `json:"protocol" tf:"protocol"`
}

type MachineScaleSetSpecOsProfileWindowsConfig struct {
	// +optional
	AdditionalUnattendConfig []MachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config"`
	// +optional
	EnableAutomaticUpgrades *bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades"`
	// +optional
	ProvisionVmAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`
	// +optional
	Winrm []MachineScaleSetSpecOsProfileWindowsConfigWinrm `json:"winrm,omitempty" tf:"winrm"`
}

type MachineScaleSetSpecPlan struct {
	Name      *string `json:"name" tf:"name"`
	Product   *string `json:"product" tf:"product"`
	Publisher *string `json:"publisher" tf:"publisher"`
}

type MachineScaleSetSpecRollingUpgradePolicy struct {
	// +optional
	MaxBatchInstancePercent *int64 `json:"maxBatchInstancePercent,omitempty" tf:"max_batch_instance_percent"`
	// +optional
	MaxUnhealthyInstancePercent *int64 `json:"maxUnhealthyInstancePercent,omitempty" tf:"max_unhealthy_instance_percent"`
	// +optional
	MaxUnhealthyUpgradedInstancePercent *int64 `json:"maxUnhealthyUpgradedInstancePercent,omitempty" tf:"max_unhealthy_upgraded_instance_percent"`
	// +optional
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches,omitempty" tf:"pause_time_between_batches"`
}

type MachineScaleSetSpecSku struct {
	Capacity *int64  `json:"capacity" tf:"capacity"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	Tier *string `json:"tier,omitempty" tf:"tier"`
}

type MachineScaleSetSpecStorageProfileDataDisk struct {
	// +optional
	Caching      *string `json:"caching,omitempty" tf:"caching"`
	CreateOption *string `json:"createOption" tf:"create_option"`
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	Lun        *int64 `json:"lun" tf:"lun"`
	// +optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`
}

type MachineScaleSetSpecStorageProfileImageReference struct {
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

type MachineScaleSetSpecStorageProfileOsDisk struct {
	// +optional
	Caching      *string `json:"caching,omitempty" tf:"caching"`
	CreateOption *string `json:"createOption" tf:"create_option"`
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// +optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`
	// +optional
	VhdContainers []string `json:"vhdContainers,omitempty" tf:"vhd_containers"`
}

type MachineScaleSetSpec struct {
	KubeformOutput *MachineScaleSetSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource MachineScaleSetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type MachineScaleSetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutomaticOsUpgrade *bool `json:"automaticOsUpgrade,omitempty" tf:"automatic_os_upgrade"`
	// +optional
	BootDiagnostics *MachineScaleSetSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`
	// +optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`
	// +optional
	Extension []MachineScaleSetSpecExtension `json:"extension,omitempty" tf:"extension"`
	// +optional
	HealthProbeID *string `json:"healthProbeID,omitempty" tf:"health_probe_id"`
	// +optional
	Identity *MachineScaleSetSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LicenseType    *string                             `json:"licenseType,omitempty" tf:"license_type"`
	Location       *string                             `json:"location" tf:"location"`
	Name           *string                             `json:"name" tf:"name"`
	NetworkProfile []MachineScaleSetSpecNetworkProfile `json:"networkProfile" tf:"network_profile"`
	OsProfile      *MachineScaleSetSpecOsProfile       `json:"osProfile" tf:"os_profile"`
	// +optional
	OsProfileLinuxConfig *MachineScaleSetSpecOsProfileLinuxConfig `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config"`
	// +optional
	OsProfileSecrets []MachineScaleSetSpecOsProfileSecrets `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets"`
	// +optional
	OsProfileWindowsConfig *MachineScaleSetSpecOsProfileWindowsConfig `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config"`
	// +optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision"`
	// +optional
	Plan *MachineScaleSetSpecPlan `json:"plan,omitempty" tf:"plan"`
	// +optional
	Priority *string `json:"priority,omitempty" tf:"priority"`
	// +optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id"`
	ResourceGroupName         *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RollingUpgradePolicy *MachineScaleSetSpecRollingUpgradePolicy `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy"`
	// +optional
	SinglePlacementGroup *bool                   `json:"singlePlacementGroup,omitempty" tf:"single_placement_group"`
	Sku                  *MachineScaleSetSpecSku `json:"sku" tf:"sku"`
	// +optional
	StorageProfileDataDisk []MachineScaleSetSpecStorageProfileDataDisk `json:"storageProfileDataDisk,omitempty" tf:"storage_profile_data_disk"`
	// +optional
	StorageProfileImageReference *MachineScaleSetSpecStorageProfileImageReference `json:"storageProfileImageReference,omitempty" tf:"storage_profile_image_reference"`
	StorageProfileOsDisk         *MachineScaleSetSpecStorageProfileOsDisk         `json:"storageProfileOsDisk" tf:"storage_profile_os_disk"`
	// +optional
	Tags              *map[string]string `json:"tags,omitempty" tf:"tags"`
	UpgradePolicyMode *string            `json:"upgradePolicyMode" tf:"upgrade_policy_mode"`
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type MachineScaleSetStatus struct {
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

// MachineScaleSetList is a list of MachineScaleSets
type MachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MachineScaleSet CRD objects
	Items []MachineScaleSet `json:"items,omitempty"`
}
