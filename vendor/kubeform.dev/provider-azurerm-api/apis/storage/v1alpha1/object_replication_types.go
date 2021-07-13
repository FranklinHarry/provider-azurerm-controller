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

type ObjectReplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectReplicationSpec   `json:"spec,omitempty"`
	Status            ObjectReplicationStatus `json:"status,omitempty"`
}

type ObjectReplicationSpecRules struct {
	// +optional
	CopyBlobsCreatedAfter    *string `json:"copyBlobsCreatedAfter,omitempty" tf:"copy_blobs_created_after"`
	DestinationContainerName *string `json:"destinationContainerName" tf:"destination_container_name"`
	// +optional
	FilterOutBlobsWithPrefix []string `json:"filterOutBlobsWithPrefix,omitempty" tf:"filter_out_blobs_with_prefix"`
	// +optional
	Name                *string `json:"name,omitempty" tf:"name"`
	SourceContainerName *string `json:"sourceContainerName" tf:"source_container_name"`
}

type ObjectReplicationSpec struct {
	State *ObjectReplicationSpecResource `json:"state,omitempty" tf:"-"`

	Resource ObjectReplicationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ObjectReplicationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DestinationObjectReplicationID *string                      `json:"destinationObjectReplicationID,omitempty" tf:"destination_object_replication_id"`
	DestinationStorageAccountID    *string                      `json:"destinationStorageAccountID" tf:"destination_storage_account_id"`
	Rules                          []ObjectReplicationSpecRules `json:"rules" tf:"rules"`
	// +optional
	SourceObjectReplicationID *string `json:"sourceObjectReplicationID,omitempty" tf:"source_object_replication_id"`
	SourceStorageAccountID    *string `json:"sourceStorageAccountID" tf:"source_storage_account_id"`
}

type ObjectReplicationStatus struct {
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

// ObjectReplicationList is a list of ObjectReplications
type ObjectReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ObjectReplication CRD objects
	Items []ObjectReplication `json:"items,omitempty"`
}
