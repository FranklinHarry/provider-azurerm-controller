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

type AnalyticsStreamInputEventhub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsStreamInputEventhubSpec   `json:"spec,omitempty"`
	Status            AnalyticsStreamInputEventhubStatus `json:"status,omitempty"`
}

type AnalyticsStreamInputEventhubSpecSerialization struct {
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	// +optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter"`
	Type           *string `json:"type" tf:"type"`
}

type AnalyticsStreamInputEventhubSpec struct {
	State *AnalyticsStreamInputEventhubSpecResource `json:"state,omitempty" tf:"-"`

	Resource AnalyticsStreamInputEventhubSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type AnalyticsStreamInputEventhubSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	EventhubConsumerGroupName *string                                        `json:"eventhubConsumerGroupName" tf:"eventhub_consumer_group_name"`
	EventhubName              *string                                        `json:"eventhubName" tf:"eventhub_name"`
	Name                      *string                                        `json:"name" tf:"name"`
	ResourceGroupName         *string                                        `json:"resourceGroupName" tf:"resource_group_name"`
	Serialization             *AnalyticsStreamInputEventhubSpecSerialization `json:"serialization" tf:"serialization"`
	ServicebusNamespace       *string                                        `json:"servicebusNamespace" tf:"servicebus_namespace"`
	SharedAccessPolicyKey     *string                                        `json:"-" sensitive:"true" tf:"shared_access_policy_key"`
	SharedAccessPolicyName    *string                                        `json:"sharedAccessPolicyName" tf:"shared_access_policy_name"`
	StreamAnalyticsJobName    *string                                        `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type AnalyticsStreamInputEventhubStatus struct {
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

// AnalyticsStreamInputEventhubList is a list of AnalyticsStreamInputEventhubs
type AnalyticsStreamInputEventhubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AnalyticsStreamInputEventhub CRD objects
	Items []AnalyticsStreamInputEventhub `json:"items,omitempty"`
}
