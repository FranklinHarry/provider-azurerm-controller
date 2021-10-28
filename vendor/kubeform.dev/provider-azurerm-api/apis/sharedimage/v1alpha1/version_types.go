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

type Version struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VersionSpec   `json:"spec,omitempty"`
	Status            VersionStatus `json:"status,omitempty"`
}

type VersionSpecTargetRegion struct {
	Name                 *string `json:"name" tf:"name"`
	RegionalReplicaCount *int64  `json:"regionalReplicaCount" tf:"regional_replica_count"`
	// +optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type"`
}

type VersionSpec struct {
	State *VersionSpecResource `json:"state,omitempty" tf:"-"`

	Resource VersionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VersionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ExcludeFromLatest *bool   `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest"`
	GalleryName       *string `json:"galleryName" tf:"gallery_name"`
	ImageName         *string `json:"imageName" tf:"image_name"`
	Location          *string `json:"location" tf:"location"`
	// +optional
	ManagedImageID *string `json:"managedImageID,omitempty" tf:"managed_image_id"`
	Name           *string `json:"name" tf:"name"`
	// +optional
	OsDiskSnapshotID  *string `json:"osDiskSnapshotID,omitempty" tf:"os_disk_snapshot_id"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags         *map[string]string        `json:"tags,omitempty" tf:"tags"`
	TargetRegion []VersionSpecTargetRegion `json:"targetRegion" tf:"target_region"`
}

type VersionStatus struct {
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

// VersionList is a list of Versions
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Version CRD objects
	Items []Version `json:"items,omitempty"`
}
