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

type ChannelDirectline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelDirectlineSpec   `json:"spec,omitempty"`
	Status            ChannelDirectlineStatus `json:"status,omitempty"`
}

type ChannelDirectlineSpecSite struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	EnhancedAuthenticationEnabled *bool `json:"enhancedAuthenticationEnabled,omitempty" tf:"enhanced_authentication_enabled"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Key *string `json:"-" sensitive:"true" tf:"key"`
	// +optional
	Key2 *string `json:"-" sensitive:"true" tf:"key2"`
	Name *string `json:"name" tf:"name"`
	// +optional
	TrustedOrigins []string `json:"trustedOrigins,omitempty" tf:"trusted_origins"`
	// +optional
	V1Allowed *bool `json:"v1Allowed,omitempty" tf:"v1_allowed"`
	// +optional
	V3Allowed *bool `json:"v3Allowed,omitempty" tf:"v3_allowed"`
}

type ChannelDirectlineSpec struct {
	KubeformOutput *ChannelDirectlineSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ChannelDirectlineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ChannelDirectlineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BotName           *string                     `json:"botName" tf:"bot_name"`
	Location          *string                     `json:"location" tf:"location"`
	ResourceGroupName *string                     `json:"resourceGroupName" tf:"resource_group_name"`
	Site              []ChannelDirectlineSpecSite `json:"site" tf:"site"`
}

type ChannelDirectlineStatus struct {
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

// ChannelDirectlineList is a list of ChannelDirectlines
type ChannelDirectlineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ChannelDirectline CRD objects
	Items []ChannelDirectline `json:"items,omitempty"`
}