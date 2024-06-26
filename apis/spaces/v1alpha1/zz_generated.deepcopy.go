//go:build !ignore_autogenerated

/*
Copyright 2020 The Upbound Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCredentials) DeepCopyInto(out *BackupCredentials) {
	*out = *in
	in.LocalCommonCredentialSelectors.DeepCopyInto(&out.LocalCommonCredentialSelectors)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCredentials.
func (in *BackupCredentials) DeepCopy() *BackupCredentials {
	if in == nil {
		return nil
	}
	out := new(BackupCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupDefinition) DeepCopyInto(out *BackupDefinition) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ConfigRef.DeepCopyInto(&out.ConfigRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupDefinition.
func (in *BackupDefinition) DeepCopy() *BackupDefinition {
	if in == nil {
		return nil
	}
	out := new(BackupDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupObjectStorage) DeepCopyInto(out *BackupObjectStorage) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	in.Credentials.DeepCopyInto(&out.Credentials)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupObjectStorage.
func (in *BackupObjectStorage) DeepCopy() *BackupObjectStorage {
	if in == nil {
		return nil
	}
	out := new(BackupObjectStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSchedule) DeepCopyInto(out *BackupSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSchedule.
func (in *BackupSchedule) DeepCopy() *BackupSchedule {
	if in == nil {
		return nil
	}
	out := new(BackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleDefinition) DeepCopyInto(out *BackupScheduleDefinition) {
	*out = *in
	in.BackupDefinition.DeepCopyInto(&out.BackupDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleDefinition.
func (in *BackupScheduleDefinition) DeepCopy() *BackupScheduleDefinition {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleList) DeepCopyInto(out *BackupScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleList.
func (in *BackupScheduleList) DeepCopy() *BackupScheduleList {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleSpec) DeepCopyInto(out *BackupScheduleSpec) {
	*out = *in
	in.BackupScheduleDefinition.DeepCopyInto(&out.BackupScheduleDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleSpec.
func (in *BackupScheduleSpec) DeepCopy() *BackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupScheduleStatus) DeepCopyInto(out *BackupScheduleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	if in.LastBackup != nil {
		in, out := &in.LastBackup, &out.LastBackup
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupScheduleStatus.
func (in *BackupScheduleStatus) DeepCopy() *BackupScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(BackupScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	in.BackupDefinition.DeepCopyInto(&out.BackupDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalCommonCredentialSelectors) DeepCopyInto(out *LocalCommonCredentialSelectors) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalSecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalCommonCredentialSelectors.
func (in *LocalCommonCredentialSelectors) DeepCopy() *LocalCommonCredentialSelectors {
	if in == nil {
		return nil
	}
	out := new(LocalCommonCredentialSelectors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSecretKeySelector) DeepCopyInto(out *LocalSecretKeySelector) {
	*out = *in
	out.LocalSecretReference = in.LocalSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSecretKeySelector.
func (in *LocalSecretKeySelector) DeepCopy() *LocalSecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(LocalSecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreciseLocalObjectReference) DeepCopyInto(out *PreciseLocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreciseLocalObjectReference.
func (in *PreciseLocalObjectReference) DeepCopy() *PreciseLocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(PreciseLocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetadata) DeepCopyInto(out *ResourceMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetadata.
func (in *ResourceMetadata) DeepCopy() *ResourceMetadata {
	if in == nil {
		return nil
	}
	out := new(ResourceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvisioningFailure) DeepCopyInto(out *SecretStoreProvisioningFailure) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.SecretStoreStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvisioningFailure.
func (in *SecretStoreProvisioningFailure) DeepCopy() *SecretStoreProvisioningFailure {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvisioningFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvisioningSuccess) DeepCopyInto(out *SecretStoreProvisioningSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvisioningSuccess.
func (in *SecretStoreProvisioningSuccess) DeepCopy() *SecretStoreProvisioningSuccess {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvisioningSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackup) DeepCopyInto(out *SharedBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackup.
func (in *SharedBackup) DeepCopy() *SharedBackup {
	if in == nil {
		return nil
	}
	out := new(SharedBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupConfig) DeepCopyInto(out *SharedBackupConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupConfig.
func (in *SharedBackupConfig) DeepCopy() *SharedBackupConfig {
	if in == nil {
		return nil
	}
	out := new(SharedBackupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackupConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupConfigList) DeepCopyInto(out *SharedBackupConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedBackupConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupConfigList.
func (in *SharedBackupConfigList) DeepCopy() *SharedBackupConfigList {
	if in == nil {
		return nil
	}
	out := new(SharedBackupConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackupConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupConfigSpec) DeepCopyInto(out *SharedBackupConfigSpec) {
	*out = *in
	in.ObjectStorage.DeepCopyInto(&out.ObjectStorage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupConfigSpec.
func (in *SharedBackupConfigSpec) DeepCopy() *SharedBackupConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SharedBackupConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupList) DeepCopyInto(out *SharedBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupList.
func (in *SharedBackupList) DeepCopy() *SharedBackupList {
	if in == nil {
		return nil
	}
	out := new(SharedBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupSchedule) DeepCopyInto(out *SharedBackupSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupSchedule.
func (in *SharedBackupSchedule) DeepCopy() *SharedBackupSchedule {
	if in == nil {
		return nil
	}
	out := new(SharedBackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackupSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupScheduleList) DeepCopyInto(out *SharedBackupScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedBackupSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupScheduleList.
func (in *SharedBackupScheduleList) DeepCopy() *SharedBackupScheduleList {
	if in == nil {
		return nil
	}
	out := new(SharedBackupScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedBackupScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupScheduleSpec) DeepCopyInto(out *SharedBackupScheduleSpec) {
	*out = *in
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.BackupScheduleDefinition.DeepCopyInto(&out.BackupScheduleDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupScheduleSpec.
func (in *SharedBackupScheduleSpec) DeepCopy() *SharedBackupScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(SharedBackupScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupScheduleStatus) DeepCopyInto(out *SharedBackupScheduleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	if in.SelectedControlPlanes != nil {
		in, out := &in.SelectedControlPlanes, &out.SelectedControlPlanes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupScheduleStatus.
func (in *SharedBackupScheduleStatus) DeepCopy() *SharedBackupScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(SharedBackupScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupSpec) DeepCopyInto(out *SharedBackupSpec) {
	*out = *in
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.BackupDefinition.DeepCopyInto(&out.BackupDefinition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupSpec.
func (in *SharedBackupSpec) DeepCopy() *SharedBackupSpec {
	if in == nil {
		return nil
	}
	out := new(SharedBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedBackupStatus) DeepCopyInto(out *SharedBackupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	if in.SelectedControlPlanes != nil {
		in, out := &in.SelectedControlPlanes, &out.SelectedControlPlanes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedBackupStatus.
func (in *SharedBackupStatus) DeepCopy() *SharedBackupStatus {
	if in == nil {
		return nil
	}
	out := new(SharedBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecret) DeepCopyInto(out *SharedExternalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecret.
func (in *SharedExternalSecret) DeepCopy() *SharedExternalSecret {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedExternalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretList) DeepCopyInto(out *SharedExternalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedExternalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretList.
func (in *SharedExternalSecretList) DeepCopy() *SharedExternalSecretList {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedExternalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretProvisioningFailure) DeepCopyInto(out *SharedExternalSecretProvisioningFailure) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.ClusterExternalSecretStatusCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretProvisioningFailure.
func (in *SharedExternalSecretProvisioningFailure) DeepCopy() *SharedExternalSecretProvisioningFailure {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretProvisioningFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretProvisioningSuccess) DeepCopyInto(out *SharedExternalSecretProvisioningSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretProvisioningSuccess.
func (in *SharedExternalSecretProvisioningSuccess) DeepCopy() *SharedExternalSecretProvisioningSuccess {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretProvisioningSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretSpec) DeepCopyInto(out *SharedExternalSecretSpec) {
	*out = *in
	if in.ExternalSecretMetadata != nil {
		in, out := &in.ExternalSecretMetadata, &out.ExternalSecretMetadata
		*out = new(ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.ExternalSecretSpec.DeepCopyInto(&out.ExternalSecretSpec)
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretSpec.
func (in *SharedExternalSecretSpec) DeepCopy() *SharedExternalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretStatus) DeepCopyInto(out *SharedExternalSecretStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]SharedExternalSecretProvisioningFailure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provisioned != nil {
		in, out := &in.Provisioned, &out.Provisioned
		*out = make([]SharedExternalSecretProvisioningSuccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretStatus.
func (in *SharedExternalSecretStatus) DeepCopy() *SharedExternalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStore) DeepCopyInto(out *SharedSecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStore.
func (in *SharedSecretStore) DeepCopy() *SharedSecretStore {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedSecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreList) DeepCopyInto(out *SharedSecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedSecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreList.
func (in *SharedSecretStoreList) DeepCopy() *SharedSecretStoreList {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedSecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreSpec) DeepCopyInto(out *SharedSecretStoreSpec) {
	*out = *in
	if in.SecretStoreMetadata != nil {
		in, out := &in.SecretStoreMetadata, &out.SecretStoreMetadata
		*out = new(ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.Provider.DeepCopyInto(&out.Provider)
	if in.RetrySettings != nil {
		in, out := &in.RetrySettings, &out.RetrySettings
		*out = new(v1beta1.SecretStoreRetrySettings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreSpec.
func (in *SharedSecretStoreSpec) DeepCopy() *SharedSecretStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreStatus) DeepCopyInto(out *SharedSecretStoreStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]SecretStoreProvisioningFailure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provisioned != nil {
		in, out := &in.Provisioned, &out.Provisioned
		*out = make([]SecretStoreProvisioningSuccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreStatus.
func (in *SharedSecretStoreStatus) DeepCopy() *SharedSecretStoreStatus {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreStatus)
	in.DeepCopyInto(out)
	return out
}
