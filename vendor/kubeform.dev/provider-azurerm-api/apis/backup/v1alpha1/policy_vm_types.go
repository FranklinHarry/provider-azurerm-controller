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

type PolicyVm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyVmSpec   `json:"spec,omitempty"`
	Status            PolicyVmStatus `json:"status,omitempty"`
}

type PolicyVmSpecBackup struct {
	Frequency *string `json:"frequency" tf:"frequency"`
	Time      *string `json:"time" tf:"time"`
	// +optional
	Weekdays []string `json:"weekdays,omitempty" tf:"weekdays"`
}

type PolicyVmSpecRetentionDaily struct {
	Count *int64 `json:"count" tf:"count"`
}

type PolicyVmSpecRetentionMonthly struct {
	Count    *int64   `json:"count" tf:"count"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
	Weeks    []string `json:"weeks" tf:"weeks"`
}

type PolicyVmSpecRetentionWeekly struct {
	Count    *int64   `json:"count" tf:"count"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
}

type PolicyVmSpecRetentionYearly struct {
	Count    *int64   `json:"count" tf:"count"`
	Months   []string `json:"months" tf:"months"`
	Weekdays []string `json:"weekdays" tf:"weekdays"`
	Weeks    []string `json:"weeks" tf:"weeks"`
}

type PolicyVmSpec struct {
	State *PolicyVmSpecResource `json:"state,omitempty" tf:"-"`

	Resource PolicyVmSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type PolicyVmSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Backup *PolicyVmSpecBackup `json:"backup" tf:"backup"`
	// +optional
	InstantRestoreRetentionDays *int64  `json:"instantRestoreRetentionDays,omitempty" tf:"instant_restore_retention_days"`
	Name                        *string `json:"name" tf:"name"`
	RecoveryVaultName           *string `json:"recoveryVaultName" tf:"recovery_vault_name"`
	ResourceGroupName           *string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RetentionDaily *PolicyVmSpecRetentionDaily `json:"retentionDaily,omitempty" tf:"retention_daily"`
	// +optional
	RetentionMonthly *PolicyVmSpecRetentionMonthly `json:"retentionMonthly,omitempty" tf:"retention_monthly"`
	// +optional
	RetentionWeekly *PolicyVmSpecRetentionWeekly `json:"retentionWeekly,omitempty" tf:"retention_weekly"`
	// +optional
	RetentionYearly *PolicyVmSpecRetentionYearly `json:"retentionYearly,omitempty" tf:"retention_yearly"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone"`
}

type PolicyVmStatus struct {
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

// PolicyVmList is a list of PolicyVms
type PolicyVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyVm CRD objects
	Items []PolicyVm `json:"items,omitempty"`
}
