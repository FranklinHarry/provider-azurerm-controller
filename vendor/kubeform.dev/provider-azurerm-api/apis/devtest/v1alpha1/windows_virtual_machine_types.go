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

type WindowsVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WindowsVirtualMachineSpec   `json:"spec,omitempty"`
	Status            WindowsVirtualMachineStatus `json:"status,omitempty"`
}

type WindowsVirtualMachineSpecGalleryImageReference struct {
	Offer     *string `json:"offer" tf:"offer"`
	Publisher *string `json:"publisher" tf:"publisher"`
	Sku       *string `json:"sku" tf:"sku"`
	Version   *string `json:"version" tf:"version"`
}

type WindowsVirtualMachineSpecInboundNATRule struct {
	BackendPort *int64 `json:"backendPort" tf:"backend_port"`
	// +optional
	FrontendPort *int64  `json:"frontendPort,omitempty" tf:"frontend_port"`
	Protocol     *string `json:"protocol" tf:"protocol"`
}

type WindowsVirtualMachineSpec struct {
	KubeformOutput *WindowsVirtualMachineSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource WindowsVirtualMachineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type WindowsVirtualMachineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowClaim *bool `json:"allowClaim,omitempty" tf:"allow_claim"`
	// +optional
	DisallowPublicIPAddress *bool `json:"disallowPublicIPAddress,omitempty" tf:"disallow_public_ip_address"`
	// +optional
	Fqdn                  *string                                         `json:"fqdn,omitempty" tf:"fqdn"`
	GalleryImageReference *WindowsVirtualMachineSpecGalleryImageReference `json:"galleryImageReference" tf:"gallery_image_reference"`
	// +optional
	InboundNATRule      []WindowsVirtualMachineSpecInboundNATRule `json:"inboundNATRule,omitempty" tf:"inbound_nat_rule"`
	LabName             *string                                   `json:"labName" tf:"lab_name"`
	LabSubnetName       *string                                   `json:"labSubnetName" tf:"lab_subnet_name"`
	LabVirtualNetworkID *string                                   `json:"labVirtualNetworkID" tf:"lab_virtual_network_id"`
	Location            *string                                   `json:"location" tf:"location"`
	Name                *string                                   `json:"name" tf:"name"`
	// +optional
	Notes             *string `json:"notes,omitempty" tf:"notes"`
	Password          *string `json:"password" tf:"password"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	Size              *string `json:"size" tf:"size"`
	StorageType       *string `json:"storageType" tf:"storage_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UniqueIdentifier *string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier"`
	Username         *string `json:"username" tf:"username"`
}

type WindowsVirtualMachineStatus struct {
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

// WindowsVirtualMachineList is a list of WindowsVirtualMachines
type WindowsVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WindowsVirtualMachine CRD objects
	Items []WindowsVirtualMachine `json:"items,omitempty"`
}
