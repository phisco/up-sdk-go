//go:build !ignore_autogenerated

/*
Copyright 2025 The Upbound Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeReference) DeepCopyInto(out *CompositeReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeReference.
func (in *CompositeReference) DeepCopy() *CompositeReference {
	if in == nil {
		return nil
	}
	out := new(CompositeReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompositeReferencePath) DeepCopyInto(out *CompositeReferencePath) {
	*out = *in
	out.CompositeReference = in.CompositeReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompositeReferencePath.
func (in *CompositeReferencePath) DeepCopy() *CompositeReferencePath {
	if in == nil {
		return nil
	}
	out := new(CompositeReferencePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectObservation) DeepCopyInto(out *ObjectObservation) {
	*out = *in
	in.Manifest.DeepCopyInto(&out.Manifest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectObservation.
func (in *ObjectObservation) DeepCopy() *ObjectObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectParameters) DeepCopyInto(out *ObjectParameters) {
	*out = *in
	in.Manifest.DeepCopyInto(&out.Manifest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectParameters.
func (in *ObjectParameters) DeepCopy() *ObjectParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	out.TypedReference = in.TypedReference
	if in.Grants != nil {
		in, out := &in.Grants, &out.Grants
		*out = make([]v1.ManagementAction, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSpec) DeepCopyInto(out *ObjectSpec) {
	*out = *in
	if in.ManagementPolicies != nil {
		in, out := &in.ManagementPolicies, &out.ManagementPolicies
		*out = make(v1.ManagementPolicies, len(*in))
		copy(*out, *in)
	}
	out.Readiness = in.Readiness
	out.Composite = in.Composite
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSpec.
func (in *ObjectSpec) DeepCopy() *ObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectsReference) DeepCopyInto(out *ObjectsReference) {
	*out = *in
	in.MatchLabels.DeepCopyInto(&out.MatchLabels)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectsReference.
func (in *ObjectsReference) DeepCopy() *ObjectsReference {
	if in == nil {
		return nil
	}
	out := new(ObjectsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Readiness) DeepCopyInto(out *Readiness) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Readiness.
func (in *Readiness) DeepCopy() *Readiness {
	if in == nil {
		return nil
	}
	out := new(Readiness)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencableKind) DeepCopyInto(out *ReferencableKind) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencableKind.
func (in *ReferencableKind) DeepCopy() *ReferencableKind {
	if in == nil {
		return nil
	}
	out := new(ReferencableKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencePath) DeepCopyInto(out *ReferencePath) {
	*out = *in
	if in.Kinds != nil {
		in, out := &in.Kinds, &out.Kinds
		*out = make([]ReferencableKind, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencePath.
func (in *ReferencePath) DeepCopy() *ReferencePath {
	if in == nil {
		return nil
	}
	out := new(ReferencePath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferenceSchema) DeepCopyInto(out *ReferenceSchema) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.References != nil {
		in, out := &in.References, &out.References
		*out = make([]ReferencePath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferenceSchema.
func (in *ReferenceSchema) DeepCopy() *ReferenceSchema {
	if in == nil {
		return nil
	}
	out := new(ReferenceSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReferenceSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencedObject) DeepCopyInto(out *ReferencedObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencedObject.
func (in *ReferencedObject) DeepCopy() *ReferencedObject {
	if in == nil {
		return nil
	}
	out := new(ReferencedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReferencedObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReferencedObjectList) DeepCopyInto(out *ReferencedObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReferencedObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReferencedObjectList.
func (in *ReferencedObjectList) DeepCopy() *ReferencedObjectList {
	if in == nil {
		return nil
	}
	out := new(ReferencedObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReferencedObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
