/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VolumeV3AttachmentObservation struct {
	Device *string `json:"device,omitempty" tf:"device,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`
}

type VolumeV3AttachmentParameters struct {
}

type VolumeV3Observation struct {
	Attachment []VolumeV3AttachmentObservation `json:"attachment,omitempty" tf:"attachment,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VolumeV3Parameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	ConsistencyGroupID *string `json:"consistencyGroupId,omitempty" tf:"consistency_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnableOnlineResize *bool `json:"enableOnlineResize,omitempty" tf:"enable_online_resize,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	Multiattach *bool `json:"multiattach,omitempty" tf:"multiattach,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulerHints []VolumeV3SchedulerHintsParameters `json:"schedulerHints,omitempty" tf:"scheduler_hints,omitempty"`

	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceReplica *string `json:"sourceReplica,omitempty" tf:"source_replica,omitempty"`

	// +kubebuilder:validation:Optional
	SourceVolID *string `json:"sourceVolId,omitempty" tf:"source_vol_id,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type VolumeV3SchedulerHintsObservation struct {
}

type VolumeV3SchedulerHintsParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	DifferentHost []*string `json:"differentHost,omitempty" tf:"different_host,omitempty"`

	// +kubebuilder:validation:Optional
	LocalToInstance *string `json:"localToInstance,omitempty" tf:"local_to_instance,omitempty"`

	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// +kubebuilder:validation:Optional
	SameHost []*string `json:"sameHost,omitempty" tf:"same_host,omitempty"`
}

// VolumeV3Spec defines the desired state of VolumeV3
type VolumeV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VolumeV3Parameters `json:"forProvider"`
}

// VolumeV3Status defines the observed state of VolumeV3.
type VolumeV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VolumeV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3 is the Schema for the VolumeV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type VolumeV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeV3Spec   `json:"spec"`
	Status            VolumeV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeV3List contains a list of VolumeV3s
type VolumeV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeV3 `json:"items"`
}

// Repository type metadata.
var (
	VolumeV3_Kind             = "VolumeV3"
	VolumeV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VolumeV3_Kind}.String()
	VolumeV3_KindAPIVersion   = VolumeV3_Kind + "." + CRDGroupVersion.String()
	VolumeV3_GroupVersionKind = CRDGroupVersion.WithKind(VolumeV3_Kind)
)

func init() {
	SchemeBuilder.Register(&VolumeV3{}, &VolumeV3List{})
}
