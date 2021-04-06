// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogSink) DeepCopyInto(out *LoggingLogSink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogSink.
func (in *LoggingLogSink) DeepCopy() *LoggingLogSink {
	if in == nil {
		return nil
	}
	out := new(LoggingLogSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingLogSink) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogSinkList) DeepCopyInto(out *LoggingLogSinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoggingLogSink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogSinkList.
func (in *LoggingLogSinkList) DeepCopy() *LoggingLogSinkList {
	if in == nil {
		return nil
	}
	out := new(LoggingLogSinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoggingLogSinkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogSinkSpec) DeepCopyInto(out *LoggingLogSinkSpec) {
	*out = *in
	out.BigqueryOptions = in.BigqueryOptions
	out.Destination = in.Destination
	if in.Exclusions != nil {
		in, out := &in.Exclusions, &out.Exclusions
		*out = make([]LogsinkExclusions, len(*in))
		copy(*out, *in)
	}
	out.FolderRef = in.FolderRef
	out.OrganizationRef = in.OrganizationRef
	out.ProjectRef = in.ProjectRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogSinkSpec.
func (in *LoggingLogSinkSpec) DeepCopy() *LoggingLogSinkSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingLogSinkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingLogSinkStatus) DeepCopyInto(out *LoggingLogSinkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingLogSinkStatus.
func (in *LoggingLogSinkStatus) DeepCopy() *LoggingLogSinkStatus {
	if in == nil {
		return nil
	}
	out := new(LoggingLogSinkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsinkBigqueryOptions) DeepCopyInto(out *LogsinkBigqueryOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsinkBigqueryOptions.
func (in *LogsinkBigqueryOptions) DeepCopy() *LogsinkBigqueryOptions {
	if in == nil {
		return nil
	}
	out := new(LogsinkBigqueryOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsinkDestination) DeepCopyInto(out *LogsinkDestination) {
	*out = *in
	out.BigQueryDatasetRef = in.BigQueryDatasetRef
	out.PubSubTopicRef = in.PubSubTopicRef
	out.StorageBucketRef = in.StorageBucketRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsinkDestination.
func (in *LogsinkDestination) DeepCopy() *LogsinkDestination {
	if in == nil {
		return nil
	}
	out := new(LogsinkDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogsinkExclusions) DeepCopyInto(out *LogsinkExclusions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogsinkExclusions.
func (in *LogsinkExclusions) DeepCopy() *LogsinkExclusions {
	if in == nil {
		return nil
	}
	out := new(LogsinkExclusions)
	in.DeepCopyInto(out)
	return out
}