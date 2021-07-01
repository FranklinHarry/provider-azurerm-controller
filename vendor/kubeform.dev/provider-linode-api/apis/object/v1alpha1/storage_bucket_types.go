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

type StorageBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketSpec   `json:"spec,omitempty"`
	Status            StorageBucketStatus `json:"status,omitempty"`
}

type StorageBucketSpecCert struct {
	// The Base64 encoded and PEM formatted SSL certificate.
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// The private key associated with the TLS/SSL certificate.
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
}

type StorageBucketSpecLifecycleRuleExpiration struct {
	// Specifies the date after which you want the corresponding action to take effect.
	// +optional
	Date *string `json:"date,omitempty" tf:"date"`
	// Specifies the number of days after object creation when the specific rule action takes effect.
	// +optional
	Days *int64 `json:"days,omitempty" tf:"days"`
	// Directs Linode Object Storage to remove expired deleted markers.
	// +optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	// Specifies the number of days non-current object versions expire.
	Days *int64 `json:"days" tf:"days"`
}

type StorageBucketSpecLifecycleRule struct {
	// Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	// +optional
	AbortIncompleteMultipartUploadDays *int64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days"`
	// Specifies whether the lifecycle rule is active.
	Enabled *bool `json:"enabled" tf:"enabled"`
	// Specifies a period in the object's expire.
	// +optional
	Expiration *StorageBucketSpecLifecycleRuleExpiration `json:"expiration,omitempty" tf:"expiration"`
	// The unique identifier for the rule.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// Specifies when non-current object versions expire.
	// +optional
	NoncurrentVersionExpiration *StorageBucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration"`
	// The object key prefix identifying one or more objects to which the rule applies.
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
}

type StorageBucketSpec struct {
	KubeformOutput *StorageBucketSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource StorageBucketSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type StorageBucketSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key"`
	// The Access Control Level of the bucket using a canned ACL string.
	// +optional
	Acl *string `json:"acl,omitempty" tf:"acl"`
	// +optional
	Cert *StorageBucketSpecCert `json:"cert,omitempty" tf:"cert"`
	// The cluster of the Linode Object Storage Bucket.
	Cluster *string `json:"cluster" tf:"cluster"`
	// If true, the bucket will be created with CORS enabled for all origins.
	// +optional
	CorsEnabled *bool `json:"corsEnabled,omitempty" tf:"cors_enabled"`
	// The label of the Linode Object Storage Bucket.
	Label *string `json:"label" tf:"label"`
	// Lifecycle rules to be applied to the bucket.
	// +optional
	LifecycleRule []StorageBucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule"`
	// +optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key"`
	// Whether to enable versioning.
	// +optional
	Versioning *bool `json:"versioning,omitempty" tf:"versioning"`
}

type StorageBucketStatus struct {
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

// StorageBucketList is a list of StorageBuckets
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucket CRD objects
	Items []StorageBucket `json:"items,omitempty"`
}
