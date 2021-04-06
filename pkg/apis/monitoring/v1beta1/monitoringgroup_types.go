// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MonitoringGroupSpec struct {
	/* A user-assigned name for this group, used only for display purposes. */
	DisplayName string `json:"displayName,omitempty"`
	/* The filter used to determine which monitored resources belong to this group. */
	Filter string `json:"filter,omitempty"`
	/* If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters. */
	IsCluster bool `json:"isCluster,omitempty"`
	/*  */
	ParentRef v1alpha1.ResourceRef `json:"parentRef,omitempty"`
	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	ResourceID string `json:"resourceID,omitempty"`
}

type MonitoringGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringGroup is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringGroupSpec   `json:"spec,omitempty"`
	Status MonitoringGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringGroupList contains a list of MonitoringGroup
type MonitoringGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringGroup{}, &MonitoringGroupList{})
}