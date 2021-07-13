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

type Elasticpool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticpoolSpec   `json:"spec,omitempty"`
	Status            ElasticpoolStatus `json:"status,omitempty"`
}

type ElasticpoolSpecPerDatabaseSettings struct {
	MaxCapacity *float64 `json:"maxCapacity" tf:"max_capacity"`
	MinCapacity *float64 `json:"minCapacity" tf:"min_capacity"`
}

type ElasticpoolSpecSku struct {
	Capacity *int64 `json:"capacity" tf:"capacity"`
	// +optional
	Family *string `json:"family,omitempty" tf:"family"`
	Name   *string `json:"name" tf:"name"`
	Tier   *string `json:"tier" tf:"tier"`
}

type ElasticpoolSpec struct {
	State *ElasticpoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource ElasticpoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElasticpoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`
	Location    *string `json:"location" tf:"location"`
	// +optional
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes"`
	// +optional
	MaxSizeGb           *float64                            `json:"maxSizeGb,omitempty" tf:"max_size_gb"`
	Name                *string                             `json:"name" tf:"name"`
	PerDatabaseSettings *ElasticpoolSpecPerDatabaseSettings `json:"perDatabaseSettings" tf:"per_database_settings"`
	ResourceGroupName   *string                             `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName          *string                             `json:"serverName" tf:"server_name"`
	Sku                 *ElasticpoolSpecSku                 `json:"sku" tf:"sku"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant"`
}

type ElasticpoolStatus struct {
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

// ElasticpoolList is a list of Elasticpools
type ElasticpoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Elasticpool CRD objects
	Items []Elasticpool `json:"items,omitempty"`
}
