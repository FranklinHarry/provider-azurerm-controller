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

type Circuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CircuitSpec   `json:"spec,omitempty"`
	Status            CircuitStatus `json:"status,omitempty"`
}

type CircuitSpecSku struct {
	Family *string `json:"family" tf:"family"`
	Tier   *string `json:"tier" tf:"tier"`
}

type CircuitSpec struct {
	State *CircuitSpecResource `json:"state,omitempty" tf:"-"`

	Resource CircuitSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CircuitSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations"`
	// +optional
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps"`
	// +optional
	BandwidthInMbps *int64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps"`
	// +optional
	ExpressRoutePortID *string `json:"expressRoutePortID,omitempty" tf:"express_route_port_id"`
	Location           *string `json:"location" tf:"location"`
	Name               *string `json:"name" tf:"name"`
	// +optional
	PeeringLocation   *string `json:"peeringLocation,omitempty" tf:"peering_location"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ServiceKey *string `json:"-" sensitive:"true" tf:"service_key"`
	// +optional
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name"`
	// +optional
	ServiceProviderProvisioningState *string         `json:"serviceProviderProvisioningState,omitempty" tf:"service_provider_provisioning_state"`
	Sku                              *CircuitSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type CircuitStatus struct {
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

// CircuitList is a list of Circuits
type CircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Circuit CRD objects
	Items []Circuit `json:"items,omitempty"`
}
