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

type Insights struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InsightsSpec   `json:"spec,omitempty"`
	Status            InsightsStatus `json:"status,omitempty"`
}

type InsightsSpec struct {
	KubeformOutput *InsightsSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource InsightsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type InsightsSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppID           *string `json:"appID,omitempty" tf:"app_id"`
	ApplicationType *string `json:"applicationType" tf:"application_type"`
	// +optional
	ConnectionString *string `json:"-" sensitive:"true" tf:"connection_string"`
	// +optional
	DailyDataCapInGb *float64 `json:"dailyDataCapInGb,omitempty" tf:"daily_data_cap_in_gb"`
	// +optional
	DailyDataCapNotificationsDisabled *bool `json:"dailyDataCapNotificationsDisabled,omitempty" tf:"daily_data_cap_notifications_disabled"`
	// +optional
	DisableIPMasking *bool `json:"disableIPMasking,omitempty" tf:"disable_ip_masking"`
	// +optional
	InstrumentationKey *string `json:"-" sensitive:"true" tf:"instrumentation_key"`
	Location           *string `json:"location" tf:"location"`
	Name               *string `json:"name" tf:"name"`
	ResourceGroupName  *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`
	// +optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
}

type InsightsStatus struct {
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

// InsightsList is a list of Insightss
type InsightsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Insights CRD objects
	Items []Insights `json:"items,omitempty"`
}