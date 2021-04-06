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
func (in *ContainerAnalysisNote) DeepCopyInto(out *ContainerAnalysisNote) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAnalysisNote.
func (in *ContainerAnalysisNote) DeepCopy() *ContainerAnalysisNote {
	if in == nil {
		return nil
	}
	out := new(ContainerAnalysisNote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerAnalysisNote) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAnalysisNoteList) DeepCopyInto(out *ContainerAnalysisNoteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerAnalysisNote, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAnalysisNoteList.
func (in *ContainerAnalysisNoteList) DeepCopy() *ContainerAnalysisNoteList {
	if in == nil {
		return nil
	}
	out := new(ContainerAnalysisNoteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerAnalysisNoteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAnalysisNoteSpec) DeepCopyInto(out *ContainerAnalysisNoteSpec) {
	*out = *in
	out.Attestation = in.Attestation
	out.Build = in.Build
	in.Deployment.DeepCopyInto(&out.Deployment)
	out.Discovery = in.Discovery
	in.Image.DeepCopyInto(&out.Image)
	in.Package.DeepCopyInto(&out.Package)
	if in.RelatedNoteNames != nil {
		in, out := &in.RelatedNoteNames, &out.RelatedNoteNames
		*out = make([]v1alpha1.ResourceRef, len(*in))
		copy(*out, *in)
	}
	if in.RelatedUrl != nil {
		in, out := &in.RelatedUrl, &out.RelatedUrl
		*out = make([]NoteRelatedUrl, len(*in))
		copy(*out, *in)
	}
	in.Vulnerability.DeepCopyInto(&out.Vulnerability)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAnalysisNoteSpec.
func (in *ContainerAnalysisNoteSpec) DeepCopy() *ContainerAnalysisNoteSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerAnalysisNoteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerAnalysisNoteStatus) DeepCopyInto(out *ContainerAnalysisNoteStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerAnalysisNoteStatus.
func (in *ContainerAnalysisNoteStatus) DeepCopy() *ContainerAnalysisNoteStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerAnalysisNoteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteAffectedVersionEnd) DeepCopyInto(out *NoteAffectedVersionEnd) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteAffectedVersionEnd.
func (in *NoteAffectedVersionEnd) DeepCopy() *NoteAffectedVersionEnd {
	if in == nil {
		return nil
	}
	out := new(NoteAffectedVersionEnd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteAffectedVersionStart) DeepCopyInto(out *NoteAffectedVersionStart) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteAffectedVersionStart.
func (in *NoteAffectedVersionStart) DeepCopy() *NoteAffectedVersionStart {
	if in == nil {
		return nil
	}
	out := new(NoteAffectedVersionStart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteAttestation) DeepCopyInto(out *NoteAttestation) {
	*out = *in
	out.Hint = in.Hint
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteAttestation.
func (in *NoteAttestation) DeepCopy() *NoteAttestation {
	if in == nil {
		return nil
	}
	out := new(NoteAttestation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteBuild) DeepCopyInto(out *NoteBuild) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteBuild.
func (in *NoteBuild) DeepCopy() *NoteBuild {
	if in == nil {
		return nil
	}
	out := new(NoteBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteCvssV3) DeepCopyInto(out *NoteCvssV3) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteCvssV3.
func (in *NoteCvssV3) DeepCopy() *NoteCvssV3 {
	if in == nil {
		return nil
	}
	out := new(NoteCvssV3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteDeployment) DeepCopyInto(out *NoteDeployment) {
	*out = *in
	if in.ResourceUri != nil {
		in, out := &in.ResourceUri, &out.ResourceUri
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteDeployment.
func (in *NoteDeployment) DeepCopy() *NoteDeployment {
	if in == nil {
		return nil
	}
	out := new(NoteDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteDetails) DeepCopyInto(out *NoteDetails) {
	*out = *in
	out.AffectedVersionEnd = in.AffectedVersionEnd
	out.AffectedVersionStart = in.AffectedVersionStart
	out.FixedVersion = in.FixedVersion
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteDetails.
func (in *NoteDetails) DeepCopy() *NoteDetails {
	if in == nil {
		return nil
	}
	out := new(NoteDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteDiscovery) DeepCopyInto(out *NoteDiscovery) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteDiscovery.
func (in *NoteDiscovery) DeepCopy() *NoteDiscovery {
	if in == nil {
		return nil
	}
	out := new(NoteDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteDistribution) DeepCopyInto(out *NoteDistribution) {
	*out = *in
	out.LatestVersion = in.LatestVersion
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteDistribution.
func (in *NoteDistribution) DeepCopy() *NoteDistribution {
	if in == nil {
		return nil
	}
	out := new(NoteDistribution)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteFingerprint) DeepCopyInto(out *NoteFingerprint) {
	*out = *in
	if in.V2Blob != nil {
		in, out := &in.V2Blob, &out.V2Blob
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteFingerprint.
func (in *NoteFingerprint) DeepCopy() *NoteFingerprint {
	if in == nil {
		return nil
	}
	out := new(NoteFingerprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteFingerprintStatus) DeepCopyInto(out *NoteFingerprintStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteFingerprintStatus.
func (in *NoteFingerprintStatus) DeepCopy() *NoteFingerprintStatus {
	if in == nil {
		return nil
	}
	out := new(NoteFingerprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteFixedVersion) DeepCopyInto(out *NoteFixedVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteFixedVersion.
func (in *NoteFixedVersion) DeepCopy() *NoteFixedVersion {
	if in == nil {
		return nil
	}
	out := new(NoteFixedVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteFixingKbs) DeepCopyInto(out *NoteFixingKbs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteFixingKbs.
func (in *NoteFixingKbs) DeepCopy() *NoteFixingKbs {
	if in == nil {
		return nil
	}
	out := new(NoteFixingKbs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteHint) DeepCopyInto(out *NoteHint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteHint.
func (in *NoteHint) DeepCopy() *NoteHint {
	if in == nil {
		return nil
	}
	out := new(NoteHint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteImage) DeepCopyInto(out *NoteImage) {
	*out = *in
	in.Fingerprint.DeepCopyInto(&out.Fingerprint)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteImage.
func (in *NoteImage) DeepCopy() *NoteImage {
	if in == nil {
		return nil
	}
	out := new(NoteImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteImageStatus) DeepCopyInto(out *NoteImageStatus) {
	*out = *in
	out.Fingerprint = in.Fingerprint
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteImageStatus.
func (in *NoteImageStatus) DeepCopy() *NoteImageStatus {
	if in == nil {
		return nil
	}
	out := new(NoteImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteLatestVersion) DeepCopyInto(out *NoteLatestVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteLatestVersion.
func (in *NoteLatestVersion) DeepCopy() *NoteLatestVersion {
	if in == nil {
		return nil
	}
	out := new(NoteLatestVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotePackage) DeepCopyInto(out *NotePackage) {
	*out = *in
	if in.Distribution != nil {
		in, out := &in.Distribution, &out.Distribution
		*out = make([]NoteDistribution, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotePackage.
func (in *NotePackage) DeepCopy() *NotePackage {
	if in == nil {
		return nil
	}
	out := new(NotePackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteRelatedUrl) DeepCopyInto(out *NoteRelatedUrl) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteRelatedUrl.
func (in *NoteRelatedUrl) DeepCopy() *NoteRelatedUrl {
	if in == nil {
		return nil
	}
	out := new(NoteRelatedUrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteVulnerability) DeepCopyInto(out *NoteVulnerability) {
	*out = *in
	out.CvssV3 = in.CvssV3
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]NoteDetails, len(*in))
		copy(*out, *in)
	}
	if in.WindowsDetails != nil {
		in, out := &in.WindowsDetails, &out.WindowsDetails
		*out = make([]NoteWindowsDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteVulnerability.
func (in *NoteVulnerability) DeepCopy() *NoteVulnerability {
	if in == nil {
		return nil
	}
	out := new(NoteVulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NoteWindowsDetails) DeepCopyInto(out *NoteWindowsDetails) {
	*out = *in
	if in.FixingKbs != nil {
		in, out := &in.FixingKbs, &out.FixingKbs
		*out = make([]NoteFixingKbs, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NoteWindowsDetails.
func (in *NoteWindowsDetails) DeepCopy() *NoteWindowsDetails {
	if in == nil {
		return nil
	}
	out := new(NoteWindowsDetails)
	in.DeepCopyInto(out)
	return out
}