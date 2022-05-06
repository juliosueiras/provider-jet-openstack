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

type VipV1Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`
}

type VipV1Parameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// +kubebuilder:validation:Optional
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FloatingIP *string `json:"floatingIp,omitempty" tf:"floating_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Persistence map[string]*string `json:"persistence,omitempty" tf:"persistence,omitempty"`

	// +kubebuilder:validation:Required
	PoolID *string `json:"poolId" tf:"pool_id,omitempty"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

// VipV1Spec defines the desired state of VipV1
type VipV1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VipV1Parameters `json:"forProvider"`
}

// VipV1Status defines the observed state of VipV1.
type VipV1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VipV1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VipV1 is the Schema for the VipV1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstackjet}
type VipV1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VipV1Spec   `json:"spec"`
	Status            VipV1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VipV1List contains a list of VipV1s
type VipV1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VipV1 `json:"items"`
}

// Repository type metadata.
var (
	VipV1_Kind             = "VipV1"
	VipV1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VipV1_Kind}.String()
	VipV1_KindAPIVersion   = VipV1_Kind + "." + CRDGroupVersion.String()
	VipV1_GroupVersionKind = CRDGroupVersion.WithKind(VipV1_Kind)
)

func init() {
	SchemeBuilder.Register(&VipV1{}, &VipV1List{})
}