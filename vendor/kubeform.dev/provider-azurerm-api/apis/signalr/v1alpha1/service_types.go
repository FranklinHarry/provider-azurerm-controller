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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecCors struct {
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
}

type ServiceSpecFeatures struct {
	Flag  *string `json:"flag" tf:"flag"`
	Value *string `json:"value" tf:"value"`
}

type ServiceSpecSku struct {
	Capacity *int64  `json:"capacity" tf:"capacity"`
	Name     *string `json:"name" tf:"name"`
}

type ServiceSpecUpstreamEndpoint struct {
	CategoryPattern []string `json:"categoryPattern" tf:"category_pattern"`
	EventPattern    []string `json:"eventPattern" tf:"event_pattern"`
	HubPattern      []string `json:"hubPattern" tf:"hub_pattern"`
	UrlTemplate     *string  `json:"urlTemplate" tf:"url_template"`
}

type ServiceSpec struct {
	KubeformOutput *ServiceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Cors []ServiceSpecCors `json:"cors,omitempty" tf:"cors"`
	// +optional
	Features []ServiceSpecFeatures `json:"features,omitempty" tf:"features"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	Location  *string `json:"location" tf:"location"`
	Name      *string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey *string `json:"-" sensitive:"true" tf:"primary_access_key"`
	// +optional
	PrimaryConnectionString *string `json:"-" sensitive:"true" tf:"primary_connection_string"`
	// +optional
	PublicPort        *int64  `json:"publicPort,omitempty" tf:"public_port"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey *string `json:"-" sensitive:"true" tf:"secondary_access_key"`
	// +optional
	SecondaryConnectionString *string `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	// +optional
	ServerPort *int64          `json:"serverPort,omitempty" tf:"server_port"`
	Sku        *ServiceSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UpstreamEndpoint []ServiceSpecUpstreamEndpoint `json:"upstreamEndpoint,omitempty" tf:"upstream_endpoint"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}