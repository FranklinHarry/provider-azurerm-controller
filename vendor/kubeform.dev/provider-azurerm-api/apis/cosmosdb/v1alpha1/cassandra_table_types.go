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

type CassandraTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CassandraTableSpec   `json:"spec,omitempty"`
	Status            CassandraTableStatus `json:"status,omitempty"`
}

type CassandraTableSpecAutoscaleSettings struct {
	// +optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput"`
}

type CassandraTableSpecSchemaClusterKey struct {
	Name    *string `json:"name" tf:"name"`
	OrderBy *string `json:"orderBy" tf:"order_by"`
}

type CassandraTableSpecSchemaColumn struct {
	Name *string `json:"name" tf:"name"`
	Type *string `json:"type" tf:"type"`
}

type CassandraTableSpecSchemaPartitionKey struct {
	Name *string `json:"name" tf:"name"`
}

type CassandraTableSpecSchema struct {
	// +optional
	ClusterKey []CassandraTableSpecSchemaClusterKey `json:"clusterKey,omitempty" tf:"cluster_key"`
	// +kubebuilder:validation:MinItems=1
	Column       []CassandraTableSpecSchemaColumn       `json:"column" tf:"column"`
	PartitionKey []CassandraTableSpecSchemaPartitionKey `json:"partitionKey" tf:"partition_key"`
}

type CassandraTableSpec struct {
	KubeformOutput *CassandraTableSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CassandraTableSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CassandraTableSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AnalyticalStorageTtl *int64 `json:"analyticalStorageTtl,omitempty" tf:"analytical_storage_ttl"`
	// +optional
	AutoscaleSettings   *CassandraTableSpecAutoscaleSettings `json:"autoscaleSettings,omitempty" tf:"autoscale_settings"`
	CassandraKeyspaceID *string                              `json:"cassandraKeyspaceID" tf:"cassandra_keyspace_id"`
	// +optional
	DefaultTtl *int64                    `json:"defaultTtl,omitempty" tf:"default_ttl"`
	Name       *string                   `json:"name" tf:"name"`
	Schema     *CassandraTableSpecSchema `json:"schema" tf:"schema"`
	// +optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput"`
}

type CassandraTableStatus struct {
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

// CassandraTableList is a list of CassandraTables
type CassandraTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CassandraTable CRD objects
	Items []CassandraTable `json:"items,omitempty"`
}
