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

type ApplicationDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationDefinitionSpec   `json:"spec,omitempty"`
	Status            ApplicationDefinitionStatus `json:"status,omitempty"`
}

type ApplicationDefinitionSpecAuthorization struct {
	RoleDefinitionID   *string `json:"roleDefinitionID" tf:"role_definition_id"`
	ServicePrincipalID *string `json:"servicePrincipalID" tf:"service_principal_id"`
}

type ApplicationDefinitionSpec struct {
	State *ApplicationDefinitionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApplicationDefinitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApplicationDefinitionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	Authorization []ApplicationDefinitionSpecAuthorization `json:"authorization,omitempty" tf:"authorization"`
	// +optional
	CreateUiDefinition *string `json:"createUiDefinition,omitempty" tf:"create_ui_definition"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	Location    *string `json:"location" tf:"location"`
	LockLevel   *string `json:"lockLevel" tf:"lock_level"`
	// +optional
	MainTemplate *string `json:"mainTemplate,omitempty" tf:"main_template"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	PackageEnabled *bool `json:"packageEnabled,omitempty" tf:"package_enabled"`
	// +optional
	PackageFileURI    *string `json:"packageFileURI,omitempty" tf:"package_file_uri"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ApplicationDefinitionStatus struct {
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

// ApplicationDefinitionList is a list of ApplicationDefinitions
type ApplicationDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationDefinition CRD objects
	Items []ApplicationDefinition `json:"items,omitempty"`
}
