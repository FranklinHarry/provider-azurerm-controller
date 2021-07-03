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

type ActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActivityLogAlertSpec   `json:"spec,omitempty"`
	Status            ActivityLogAlertStatus `json:"status,omitempty"`
}

type ActivityLogAlertSpecAction struct {
	ActionGroupID *string `json:"actionGroupID" tf:"action_group_id"`
	// +optional
	WebhookProperties *map[string]string `json:"webhookProperties,omitempty" tf:"webhook_properties"`
}

type ActivityLogAlertSpecCriteriaServiceHealth struct {
	// +optional
	Events []string `json:"events,omitempty" tf:"events"`
	// +optional
	Locations []string `json:"locations,omitempty" tf:"locations"`
	// +optional
	Services []string `json:"services,omitempty" tf:"services"`
}

type ActivityLogAlertSpecCriteria struct {
	// +optional
	Caller   *string `json:"caller,omitempty" tf:"caller"`
	Category *string `json:"category" tf:"category"`
	// +optional
	Level *string `json:"level,omitempty" tf:"level"`
	// +optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name"`
	// +optional
	RecommendationCategory *string `json:"recommendationCategory,omitempty" tf:"recommendation_category"`
	// +optional
	RecommendationImpact *string `json:"recommendationImpact,omitempty" tf:"recommendation_impact"`
	// +optional
	RecommendationType *string `json:"recommendationType,omitempty" tf:"recommendation_type"`
	// +optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group"`
	// +optional
	ResourceID *string `json:"resourceID,omitempty" tf:"resource_id"`
	// +optional
	ResourceProvider *string `json:"resourceProvider,omitempty" tf:"resource_provider"`
	// +optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type"`
	// +optional
	ServiceHealth []ActivityLogAlertSpecCriteriaServiceHealth `json:"serviceHealth,omitempty" tf:"service_health"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status"`
}

type ActivityLogAlertSpec struct {
	KubeformOutput *ActivityLogAlertSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ActivityLogAlertSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ActivityLogAlertSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Action   []ActivityLogAlertSpecAction  `json:"action,omitempty" tf:"action"`
	Criteria *ActivityLogAlertSpecCriteria `json:"criteria" tf:"criteria"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Enabled           *bool   `json:"enabled,omitempty" tf:"enabled"`
	Name              *string `json:"name" tf:"name"`
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	Scopes []string `json:"scopes" tf:"scopes"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type ActivityLogAlertStatus struct {
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

// ActivityLogAlertList is a list of ActivityLogAlerts
type ActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ActivityLogAlert CRD objects
	Items []ActivityLogAlert `json:"items,omitempty"`
}
