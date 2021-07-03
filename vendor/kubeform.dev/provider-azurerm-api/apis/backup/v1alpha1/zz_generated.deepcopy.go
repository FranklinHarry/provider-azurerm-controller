// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageAccount) DeepCopyInto(out *ContainerStorageAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageAccount.
func (in *ContainerStorageAccount) DeepCopy() *ContainerStorageAccount {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerStorageAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageAccountList) DeepCopyInto(out *ContainerStorageAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerStorageAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageAccountList.
func (in *ContainerStorageAccountList) DeepCopy() *ContainerStorageAccountList {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerStorageAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageAccountSpec) DeepCopyInto(out *ContainerStorageAccountSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ContainerStorageAccountSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageAccountSpec.
func (in *ContainerStorageAccountSpec) DeepCopy() *ContainerStorageAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageAccountSpecResource) DeepCopyInto(out *ContainerStorageAccountSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.StorageAccountID != nil {
		in, out := &in.StorageAccountID, &out.StorageAccountID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageAccountSpecResource.
func (in *ContainerStorageAccountSpecResource) DeepCopy() *ContainerStorageAccountSpecResource {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageAccountSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageAccountStatus) DeepCopyInto(out *ContainerStorageAccountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageAccountStatus.
func (in *ContainerStorageAccountStatus) DeepCopy() *ContainerStorageAccountStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShare) DeepCopyInto(out *PolicyFileShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShare.
func (in *PolicyFileShare) DeepCopy() *PolicyFileShare {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyFileShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareList) DeepCopyInto(out *PolicyFileShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyFileShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareList.
func (in *PolicyFileShareList) DeepCopy() *PolicyFileShareList {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyFileShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpec) DeepCopyInto(out *PolicyFileShareSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(PolicyFileShareSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpec.
func (in *PolicyFileShareSpec) DeepCopy() *PolicyFileShareSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecBackup) DeepCopyInto(out *PolicyFileShareSpecBackup) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecBackup.
func (in *PolicyFileShareSpecBackup) DeepCopy() *PolicyFileShareSpecBackup {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecResource) DeepCopyInto(out *PolicyFileShareSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(PolicyFileShareSpecBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.RetentionDaily != nil {
		in, out := &in.RetentionDaily, &out.RetentionDaily
		*out = new(PolicyFileShareSpecRetentionDaily)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionMonthly != nil {
		in, out := &in.RetentionMonthly, &out.RetentionMonthly
		*out = new(PolicyFileShareSpecRetentionMonthly)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionWeekly != nil {
		in, out := &in.RetentionWeekly, &out.RetentionWeekly
		*out = new(PolicyFileShareSpecRetentionWeekly)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionYearly != nil {
		in, out := &in.RetentionYearly, &out.RetentionYearly
		*out = new(PolicyFileShareSpecRetentionYearly)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecResource.
func (in *PolicyFileShareSpecResource) DeepCopy() *PolicyFileShareSpecResource {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecRetentionDaily) DeepCopyInto(out *PolicyFileShareSpecRetentionDaily) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecRetentionDaily.
func (in *PolicyFileShareSpecRetentionDaily) DeepCopy() *PolicyFileShareSpecRetentionDaily {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecRetentionDaily)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecRetentionMonthly) DeepCopyInto(out *PolicyFileShareSpecRetentionMonthly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecRetentionMonthly.
func (in *PolicyFileShareSpecRetentionMonthly) DeepCopy() *PolicyFileShareSpecRetentionMonthly {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecRetentionMonthly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecRetentionWeekly) DeepCopyInto(out *PolicyFileShareSpecRetentionWeekly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecRetentionWeekly.
func (in *PolicyFileShareSpecRetentionWeekly) DeepCopy() *PolicyFileShareSpecRetentionWeekly {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecRetentionWeekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareSpecRetentionYearly) DeepCopyInto(out *PolicyFileShareSpecRetentionYearly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Months != nil {
		in, out := &in.Months, &out.Months
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareSpecRetentionYearly.
func (in *PolicyFileShareSpecRetentionYearly) DeepCopy() *PolicyFileShareSpecRetentionYearly {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareSpecRetentionYearly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyFileShareStatus) DeepCopyInto(out *PolicyFileShareStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyFileShareStatus.
func (in *PolicyFileShareStatus) DeepCopy() *PolicyFileShareStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyFileShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVm) DeepCopyInto(out *PolicyVm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVm.
func (in *PolicyVm) DeepCopy() *PolicyVm {
	if in == nil {
		return nil
	}
	out := new(PolicyVm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyVm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmList) DeepCopyInto(out *PolicyVmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyVm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmList.
func (in *PolicyVmList) DeepCopy() *PolicyVmList {
	if in == nil {
		return nil
	}
	out := new(PolicyVmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyVmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpec) DeepCopyInto(out *PolicyVmSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(PolicyVmSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpec.
func (in *PolicyVmSpec) DeepCopy() *PolicyVmSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecBackup) DeepCopyInto(out *PolicyVmSpecBackup) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecBackup.
func (in *PolicyVmSpecBackup) DeepCopy() *PolicyVmSpecBackup {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecResource) DeepCopyInto(out *PolicyVmSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(PolicyVmSpecBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.InstantRestoreRetentionDays != nil {
		in, out := &in.InstantRestoreRetentionDays, &out.InstantRestoreRetentionDays
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.RetentionDaily != nil {
		in, out := &in.RetentionDaily, &out.RetentionDaily
		*out = new(PolicyVmSpecRetentionDaily)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionMonthly != nil {
		in, out := &in.RetentionMonthly, &out.RetentionMonthly
		*out = new(PolicyVmSpecRetentionMonthly)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionWeekly != nil {
		in, out := &in.RetentionWeekly, &out.RetentionWeekly
		*out = new(PolicyVmSpecRetentionWeekly)
		(*in).DeepCopyInto(*out)
	}
	if in.RetentionYearly != nil {
		in, out := &in.RetentionYearly, &out.RetentionYearly
		*out = new(PolicyVmSpecRetentionYearly)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecResource.
func (in *PolicyVmSpecResource) DeepCopy() *PolicyVmSpecResource {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecRetentionDaily) DeepCopyInto(out *PolicyVmSpecRetentionDaily) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecRetentionDaily.
func (in *PolicyVmSpecRetentionDaily) DeepCopy() *PolicyVmSpecRetentionDaily {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecRetentionDaily)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecRetentionMonthly) DeepCopyInto(out *PolicyVmSpecRetentionMonthly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecRetentionMonthly.
func (in *PolicyVmSpecRetentionMonthly) DeepCopy() *PolicyVmSpecRetentionMonthly {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecRetentionMonthly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecRetentionWeekly) DeepCopyInto(out *PolicyVmSpecRetentionWeekly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecRetentionWeekly.
func (in *PolicyVmSpecRetentionWeekly) DeepCopy() *PolicyVmSpecRetentionWeekly {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecRetentionWeekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmSpecRetentionYearly) DeepCopyInto(out *PolicyVmSpecRetentionYearly) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Months != nil {
		in, out := &in.Months, &out.Months
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Weeks != nil {
		in, out := &in.Weeks, &out.Weeks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmSpecRetentionYearly.
func (in *PolicyVmSpecRetentionYearly) DeepCopy() *PolicyVmSpecRetentionYearly {
	if in == nil {
		return nil
	}
	out := new(PolicyVmSpecRetentionYearly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVmStatus) DeepCopyInto(out *PolicyVmStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVmStatus.
func (in *PolicyVmStatus) DeepCopy() *PolicyVmStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyVmStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedFileShare) DeepCopyInto(out *ProtectedFileShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedFileShare.
func (in *ProtectedFileShare) DeepCopy() *ProtectedFileShare {
	if in == nil {
		return nil
	}
	out := new(ProtectedFileShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectedFileShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedFileShareList) DeepCopyInto(out *ProtectedFileShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProtectedFileShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedFileShareList.
func (in *ProtectedFileShareList) DeepCopy() *ProtectedFileShareList {
	if in == nil {
		return nil
	}
	out := new(ProtectedFileShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectedFileShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedFileShareSpec) DeepCopyInto(out *ProtectedFileShareSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ProtectedFileShareSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedFileShareSpec.
func (in *ProtectedFileShareSpec) DeepCopy() *ProtectedFileShareSpec {
	if in == nil {
		return nil
	}
	out := new(ProtectedFileShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedFileShareSpecResource) DeepCopyInto(out *ProtectedFileShareSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupPolicyID != nil {
		in, out := &in.BackupPolicyID, &out.BackupPolicyID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceFileShareName != nil {
		in, out := &in.SourceFileShareName, &out.SourceFileShareName
		*out = new(string)
		**out = **in
	}
	if in.SourceStorageAccountID != nil {
		in, out := &in.SourceStorageAccountID, &out.SourceStorageAccountID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedFileShareSpecResource.
func (in *ProtectedFileShareSpecResource) DeepCopy() *ProtectedFileShareSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProtectedFileShareSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedFileShareStatus) DeepCopyInto(out *ProtectedFileShareStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedFileShareStatus.
func (in *ProtectedFileShareStatus) DeepCopy() *ProtectedFileShareStatus {
	if in == nil {
		return nil
	}
	out := new(ProtectedFileShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedVm) DeepCopyInto(out *ProtectedVm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedVm.
func (in *ProtectedVm) DeepCopy() *ProtectedVm {
	if in == nil {
		return nil
	}
	out := new(ProtectedVm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectedVm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedVmList) DeepCopyInto(out *ProtectedVmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProtectedVm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedVmList.
func (in *ProtectedVmList) DeepCopy() *ProtectedVmList {
	if in == nil {
		return nil
	}
	out := new(ProtectedVmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProtectedVmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedVmSpec) DeepCopyInto(out *ProtectedVmSpec) {
	*out = *in
	if in.KubeformOutput != nil {
		in, out := &in.KubeformOutput, &out.KubeformOutput
		*out = new(ProtectedVmSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedVmSpec.
func (in *ProtectedVmSpec) DeepCopy() *ProtectedVmSpec {
	if in == nil {
		return nil
	}
	out := new(ProtectedVmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedVmSpecResource) DeepCopyInto(out *ProtectedVmSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupPolicyID != nil {
		in, out := &in.BackupPolicyID, &out.BackupPolicyID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryVaultName != nil {
		in, out := &in.RecoveryVaultName, &out.RecoveryVaultName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SourceVmID != nil {
		in, out := &in.SourceVmID, &out.SourceVmID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedVmSpecResource.
func (in *ProtectedVmSpecResource) DeepCopy() *ProtectedVmSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProtectedVmSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProtectedVmStatus) DeepCopyInto(out *ProtectedVmStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProtectedVmStatus.
func (in *ProtectedVmStatus) DeepCopy() *ProtectedVmStatus {
	if in == nil {
		return nil
	}
	out := new(ProtectedVmStatus)
	in.DeepCopyInto(out)
	return out
}
