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

type Lb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSpec   `json:"spec,omitempty"`
	Status            LbStatus `json:"status,omitempty"`
}

type LbSpecFrontendIPConfiguration struct {
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InboundNATRules []string `json:"inboundNATRules,omitempty" tf:"inbound_nat_rules"`
	// +optional
	LoadBalancerRules []string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules"`
	Name              *string  `json:"name" tf:"name"`
	// +optional
	OutboundRules []string `json:"outboundRules,omitempty" tf:"outbound_rules"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateIPAddressAllocation *string `json:"privateIPAddressAllocation,omitempty" tf:"private_ip_address_allocation"`
	// +optional
	PrivateIPAddressVersion *string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version"`
	// +optional
	PublicIPAddressID *string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id"`
	// +optional
	PublicIPPrefixID *string `json:"publicIPPrefixID,omitempty" tf:"public_ip_prefix_id"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	// Deprecated
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type LbSpec struct {
	State *LbSpecResource `json:"state,omitempty" tf:"-"`

	Resource LbSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LbSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	FrontendIPConfiguration []LbSpecFrontendIPConfiguration `json:"frontendIPConfiguration,omitempty" tf:"frontend_ip_configuration"`
	Location                *string                         `json:"location" tf:"location"`
	Name                    *string                         `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses"`
	ResourceGroupName  *string  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Sku *string `json:"sku,omitempty" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type LbStatus struct {
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

// LbList is a list of Lbs
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Lb CRD objects
	Items []Lb `json:"items,omitempty"`
}
