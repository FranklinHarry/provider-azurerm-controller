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

type LinkService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkServiceSpec   `json:"spec,omitempty"`
	Status            LinkServiceStatus `json:"status,omitempty"`
}

type LinkServiceSpecNatIPConfiguration struct {
	Name    *string `json:"name" tf:"name"`
	Primary *bool   `json:"primary" tf:"primary"`
	// +optional
	PrivateIPAddress *string `json:"privateIPAddress,omitempty" tf:"private_ip_address"`
	// +optional
	PrivateIPAddressVersion *string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version"`
	SubnetID                *string `json:"subnetID" tf:"subnet_id"`
}

type LinkServiceSpec struct {
	State *LinkServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource LinkServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LinkServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Alias *string `json:"alias,omitempty" tf:"alias"`
	// +optional
	AutoApprovalSubscriptionIDS []string `json:"autoApprovalSubscriptionIDS,omitempty" tf:"auto_approval_subscription_ids"`
	// +optional
	EnableProxyProtocol                    *bool    `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol"`
	LoadBalancerFrontendIPConfigurationIDS []string `json:"loadBalancerFrontendIPConfigurationIDS" tf:"load_balancer_frontend_ip_configuration_ids"`
	Location                               *string  `json:"location" tf:"location"`
	Name                                   *string  `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=8
	NatIPConfiguration []LinkServiceSpecNatIPConfiguration `json:"natIPConfiguration" tf:"nat_ip_configuration"`
	ResourceGroupName  *string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	VisibilitySubscriptionIDS []string `json:"visibilitySubscriptionIDS,omitempty" tf:"visibility_subscription_ids"`
}

type LinkServiceStatus struct {
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

// LinkServiceList is a list of LinkServices
type LinkServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinkService CRD objects
	Items []LinkService `json:"items,omitempty"`
}
