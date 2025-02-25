//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) 2023 Bankdata (bankdata@bankdata.dk)

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedMapping) DeepCopyInto(out *AllowedMapping) {
	*out = *in
	if in.Expected != nil {
		in, out := &in.Expected, &out.Expected
		*out = new(Expected)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedMapping.
func (in *AllowedMapping) DeepCopy() *AllowedMapping {
	if in == nil {
		return nil
	}
	out := new(AllowedMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ColumnMapping) DeepCopyInto(out *ColumnMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ColumnMapping.
func (in *ColumnMapping) DeepCopy() *ColumnMapping {
	if in == nil {
		return nil
	}
	out := new(ColumnMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datasource) DeepCopyInto(out *Datasource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datasource.
func (in *Datasource) DeepCopy() *Datasource {
	if in == nil {
		return nil
	}
	out := new(Datasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecisionMapping) DeepCopyInto(out *DecisionMapping) {
	*out = *in
	if in.Columns != nil {
		in, out := &in.Columns, &out.Columns
		*out = make([]ColumnMapping, len(*in))
		copy(*out, *in)
	}
	out.Reason = in.Reason
	if in.Allowed != nil {
		in, out := &in.Allowed, &out.Allowed
		*out = new(AllowedMapping)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecisionMapping.
func (in *DecisionMapping) DeepCopy() *DecisionMapping {
	if in == nil {
		return nil
	}
	out := new(DecisionMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Expected) DeepCopyInto(out *Expected) {
	*out = *in
	if in.String != nil {
		in, out := &in.String, &out.String
		*out = new(string)
		**out = **in
	}
	if in.Boolean != nil {
		in, out := &in.Boolean, &out.Boolean
		*out = new(bool)
		**out = **in
	}
	if in.Integer != nil {
		in, out := &in.Integer, &out.Integer
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Expected.
func (in *Expected) DeepCopy() *Expected {
	if in == nil {
		return nil
	}
	out := new(Expected)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepo) DeepCopyInto(out *GitRepo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepo.
func (in *GitRepo) DeepCopy() *GitRepo {
	if in == nil {
		return nil
	}
	out := new(GitRepo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalPlane) DeepCopyInto(out *LocalPlane) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalPlane.
func (in *LocalPlane) DeepCopy() *LocalPlane {
	if in == nil {
		return nil
	}
	out := new(LocalPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReasonMapping) DeepCopyInto(out *ReasonMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReasonMapping.
func (in *ReasonMapping) DeepCopy() *ReasonMapping {
	if in == nil {
		return nil
	}
	out := new(ReasonMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceControl) DeepCopyInto(out *SourceControl) {
	*out = *in
	out.Origin = in.Origin
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceControl.
func (in *SourceControl) DeepCopy() *SourceControl {
	if in == nil {
		return nil
	}
	out := new(SourceControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subject) DeepCopyInto(out *Subject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subject.
func (in *Subject) DeepCopy() *Subject {
	if in == nil {
		return nil
	}
	out := new(Subject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *System) DeepCopyInto(out *System) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new System.
func (in *System) DeepCopy() *System {
	if in == nil {
		return nil
	}
	out := new(System)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *System) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemList) DeepCopyInto(out *SystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]System, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemList.
func (in *SystemList) DeepCopy() *SystemList {
	if in == nil {
		return nil
	}
	out := new(SystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemSpec) DeepCopyInto(out *SystemSpec) {
	*out = *in
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]Subject, len(*in))
		copy(*out, *in)
	}
	if in.DecisionMappings != nil {
		in, out := &in.DecisionMappings, &out.DecisionMappings
		*out = make([]DecisionMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Datasources != nil {
		in, out := &in.Datasources, &out.Datasources
		*out = make([]Datasource, len(*in))
		copy(*out, *in)
	}
	if in.SourceControl != nil {
		in, out := &in.SourceControl, &out.SourceControl
		*out = new(SourceControl)
		**out = **in
	}
	if in.LocalPlane != nil {
		in, out := &in.LocalPlane, &out.LocalPlane
		*out = new(LocalPlane)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemSpec.
func (in *SystemSpec) DeepCopy() *SystemSpec {
	if in == nil {
		return nil
	}
	out := new(SystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemStatus) DeepCopyInto(out *SystemStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemStatus.
func (in *SystemStatus) DeepCopy() *SystemStatus {
	if in == nil {
		return nil
	}
	out := new(SystemStatus)
	in.DeepCopyInto(out)
	return out
}
