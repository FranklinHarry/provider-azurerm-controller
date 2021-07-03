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

type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineSpec   `json:"spec,omitempty"`
	Status            VirtualMachineStatus `json:"status,omitempty"`
}

type VirtualMachineSpecAdditionalCapabilities struct {
	// +optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled"`
}

type VirtualMachineSpecAdminSSHKey struct {
	PublicKey *string `json:"publicKey" tf:"public_key"`
	Username  *string `json:"username" tf:"username"`
}

type VirtualMachineSpecBootDiagnostics struct {
	// +optional
	StorageAccountURI *string `json:"storageAccountURI,omitempty" tf:"storage_account_uri"`
}

type VirtualMachineSpecIdentity struct {
	// +optional
	IdentityIDS []string `json:"identityIDS,omitempty" tf:"identity_ids"`
	// +optional
	PrincipalID *string `json:"principalID,omitempty" tf:"principal_id"`
	// +optional
	TenantID *string `json:"tenantID,omitempty" tf:"tenant_id"`
	Type     *string `json:"type" tf:"type"`
}

type VirtualMachineSpecOsDiskDiffDiskSettings struct {
	Option *string `json:"option" tf:"option"`
}

type VirtualMachineSpecOsDisk struct {
	Caching *string `json:"caching" tf:"caching"`
	// +optional
	DiffDiskSettings *VirtualMachineSpecOsDiskDiffDiskSettings `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings"`
	// +optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetID,omitempty" tf:"disk_encryption_set_id"`
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// +optional
	Name               *string `json:"name,omitempty" tf:"name"`
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type"`
	// +optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled"`
}

type VirtualMachineSpecPlan struct {
	Name      *string `json:"name" tf:"name"`
	Product   *string `json:"product" tf:"product"`
	Publisher *string `json:"publisher" tf:"publisher"`
}

type VirtualMachineSpecSecretCertificate struct {
	Url *string `json:"url" tf:"url"`
}

type VirtualMachineSpecSecret struct {
	// +kubebuilder:validation:MinItems=1
	Certificate []VirtualMachineSpecSecretCertificate `json:"certificate" tf:"certificate"`
	KeyVaultID  *string                               `json:"keyVaultID" tf:"key_vault_id"`
}

type VirtualMachineSpecSourceImageReference struct {
	Offer     *string `json:"offer" tf:"offer"`
	Publisher *string `json:"publisher" tf:"publisher"`
	Sku       *string `json:"sku" tf:"sku"`
	Version   *string `json:"version" tf:"version"`
}

type VirtualMachineSpec struct {
	KubeformOutput *VirtualMachineSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource VirtualMachineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type VirtualMachineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalCapabilities *VirtualMachineSpecAdditionalCapabilities `json:"additionalCapabilities,omitempty" tf:"additional_capabilities"`
	// +optional
	AdminPassword *string `json:"-" sensitive:"true" tf:"admin_password"`
	// +optional
	AdminSSHKey   []VirtualMachineSpecAdminSSHKey `json:"adminSSHKey,omitempty" tf:"admin_ssh_key"`
	AdminUsername *string                         `json:"adminUsername" tf:"admin_username"`
	// +optional
	AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty" tf:"allow_extension_operations"`
	// +optional
	AvailabilitySetID *string `json:"availabilitySetID,omitempty" tf:"availability_set_id"`
	// +optional
	BootDiagnostics *VirtualMachineSpecBootDiagnostics `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`
	// +optional
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name"`
	// +optional
	CustomData *string `json:"-" sensitive:"true" tf:"custom_data"`
	// +optional
	DedicatedHostID *string `json:"dedicatedHostID,omitempty" tf:"dedicated_host_id"`
	// +optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication"`
	// +optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled"`
	// +optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`
	// +optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget"`
	// +optional
	Identity *VirtualMachineSpecIdentity `json:"identity,omitempty" tf:"identity"`
	// +optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`
	Location    *string `json:"location" tf:"location"`
	// +optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price"`
	Name        *string  `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	NetworkInterfaceIDS []string                  `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	OsDisk              *VirtualMachineSpecOsDisk `json:"osDisk" tf:"os_disk"`
	// +optional
	Plan *VirtualMachineSpecPlan `json:"plan,omitempty" tf:"plan"`
	// +optional
	PlatformFaultDomain *int64 `json:"platformFaultDomain,omitempty" tf:"platform_fault_domain"`
	// +optional
	Priority *string `json:"priority,omitempty" tf:"priority"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	// +optional
	ProvisionVmAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`
	// +optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty" tf:"proximity_placement_group_id"`
	// +optional
	PublicIPAddress *string `json:"publicIPAddress,omitempty" tf:"public_ip_address"`
	// +optional
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty" tf:"public_ip_addresses"`
	ResourceGroupName *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Secret []VirtualMachineSpecSecret `json:"secret,omitempty" tf:"secret"`
	Size   *string                    `json:"size" tf:"size"`
	// +optional
	SourceImageID *string `json:"sourceImageID,omitempty" tf:"source_image_id"`
	// +optional
	SourceImageReference *VirtualMachineSpecSourceImageReference `json:"sourceImageReference,omitempty" tf:"source_image_reference"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VirtualMachineID *string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id"`
	// +optional
	VirtualMachineScaleSetID *string `json:"virtualMachineScaleSetID,omitempty" tf:"virtual_machine_scale_set_id"`
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type VirtualMachineStatus struct {
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

// VirtualMachineList is a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachine CRD objects
	Items []VirtualMachine `json:"items,omitempty"`
}
