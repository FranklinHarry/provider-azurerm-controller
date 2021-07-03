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

type BudgetResourceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetResourceGroupSpec   `json:"spec,omitempty"`
	Status            BudgetResourceGroupStatus `json:"status,omitempty"`
}

type BudgetResourceGroupSpecFilterDimension struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values" tf:"values"`
}

type BudgetResourceGroupSpecFilterNotDimension struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values" tf:"values"`
}

type BudgetResourceGroupSpecFilterNotTag struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Operator *string  `json:"operator,omitempty" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type BudgetResourceGroupSpecFilterNot struct {
	// +optional
	Dimension *BudgetResourceGroupSpecFilterNotDimension `json:"dimension,omitempty" tf:"dimension"`
	// +optional
	Tag *BudgetResourceGroupSpecFilterNotTag `json:"tag,omitempty" tf:"tag"`
}

type BudgetResourceGroupSpecFilterTag struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Operator *string  `json:"operator,omitempty" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type BudgetResourceGroupSpecFilter struct {
	// +optional
	Dimension []BudgetResourceGroupSpecFilterDimension `json:"dimension,omitempty" tf:"dimension"`
	// +optional
	Not *BudgetResourceGroupSpecFilterNot `json:"not,omitempty" tf:"not"`
	// +optional
	Tag []BudgetResourceGroupSpecFilterTag `json:"tag,omitempty" tf:"tag"`
}

type BudgetResourceGroupSpecNotification struct {
	// +optional
	ContactEmails []string `json:"contactEmails,omitempty" tf:"contact_emails"`
	// +optional
	ContactGroups []string `json:"contactGroups,omitempty" tf:"contact_groups"`
	// +optional
	ContactRoles []string `json:"contactRoles,omitempty" tf:"contact_roles"`
	// +optional
	Enabled   *bool   `json:"enabled,omitempty" tf:"enabled"`
	Operator  *string `json:"operator" tf:"operator"`
	Threshold *int64  `json:"threshold" tf:"threshold"`
}

type BudgetResourceGroupSpecTimePeriod struct {
	// +optional
	EndDate   *string `json:"endDate,omitempty" tf:"end_date"`
	StartDate *string `json:"startDate" tf:"start_date"`
}

type BudgetResourceGroupSpec struct {
	KubeformOutput *BudgetResourceGroupSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource BudgetResourceGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BudgetResourceGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Amount *float64 `json:"amount" tf:"amount"`
	// +optional
	Filter *BudgetResourceGroupSpecFilter `json:"filter,omitempty" tf:"filter"`
	Name   *string                        `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:MinItems=1
	Notification    []BudgetResourceGroupSpecNotification `json:"notification" tf:"notification"`
	ResourceGroupID *string                               `json:"resourceGroupID" tf:"resource_group_id"`
	// +optional
	TimeGrain  *string                            `json:"timeGrain,omitempty" tf:"time_grain"`
	TimePeriod *BudgetResourceGroupSpecTimePeriod `json:"timePeriod" tf:"time_period"`
}

type BudgetResourceGroupStatus struct {
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

// BudgetResourceGroupList is a list of BudgetResourceGroups
type BudgetResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BudgetResourceGroup CRD objects
	Items []BudgetResourceGroup `json:"items,omitempty"`
}
