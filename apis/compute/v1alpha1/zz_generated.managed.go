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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AggregateV2.
func (mg *AggregateV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AggregateV2.
func (mg *AggregateV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AggregateV2.
func (mg *AggregateV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AggregateV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AggregateV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AggregateV2.
func (mg *AggregateV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AggregateV2.
func (mg *AggregateV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AggregateV2.
func (mg *AggregateV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AggregateV2.
func (mg *AggregateV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AggregateV2.
func (mg *AggregateV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AggregateV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AggregateV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AggregateV2.
func (mg *AggregateV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AggregateV2.
func (mg *AggregateV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FlavorAccessV2.
func (mg *FlavorAccessV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FlavorAccessV2.
func (mg *FlavorAccessV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FlavorAccessV2.
func (mg *FlavorAccessV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FlavorAccessV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FlavorAccessV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FlavorAccessV2.
func (mg *FlavorAccessV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FlavorAccessV2.
func (mg *FlavorAccessV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FlavorAccessV2.
func (mg *FlavorAccessV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FlavorAccessV2.
func (mg *FlavorAccessV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FlavorAccessV2.
func (mg *FlavorAccessV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FlavorAccessV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FlavorAccessV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FlavorAccessV2.
func (mg *FlavorAccessV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FlavorAccessV2.
func (mg *FlavorAccessV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FlavorV2.
func (mg *FlavorV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FlavorV2.
func (mg *FlavorV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FlavorV2.
func (mg *FlavorV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FlavorV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FlavorV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FlavorV2.
func (mg *FlavorV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FlavorV2.
func (mg *FlavorV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FlavorV2.
func (mg *FlavorV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FlavorV2.
func (mg *FlavorV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FlavorV2.
func (mg *FlavorV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FlavorV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FlavorV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FlavorV2.
func (mg *FlavorV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FlavorV2.
func (mg *FlavorV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FloatingipAssociateV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FloatingipAssociateV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FloatingipAssociateV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FloatingipAssociateV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FloatingipAssociateV2.
func (mg *FloatingipAssociateV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FloatingipV2.
func (mg *FloatingipV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FloatingipV2.
func (mg *FloatingipV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FloatingipV2.
func (mg *FloatingipV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FloatingipV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FloatingipV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FloatingipV2.
func (mg *FloatingipV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FloatingipV2.
func (mg *FloatingipV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FloatingipV2.
func (mg *FloatingipV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FloatingipV2.
func (mg *FloatingipV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FloatingipV2.
func (mg *FloatingipV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FloatingipV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FloatingipV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FloatingipV2.
func (mg *FloatingipV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FloatingipV2.
func (mg *FloatingipV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this InstanceV2.
func (mg *InstanceV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this InstanceV2.
func (mg *InstanceV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this InstanceV2.
func (mg *InstanceV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this InstanceV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *InstanceV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this InstanceV2.
func (mg *InstanceV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this InstanceV2.
func (mg *InstanceV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this InstanceV2.
func (mg *InstanceV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this InstanceV2.
func (mg *InstanceV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this InstanceV2.
func (mg *InstanceV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this InstanceV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *InstanceV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this InstanceV2.
func (mg *InstanceV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this InstanceV2.
func (mg *InstanceV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this InterfaceAttachV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *InterfaceAttachV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this InterfaceAttachV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *InterfaceAttachV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this InterfaceAttachV2.
func (mg *InterfaceAttachV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KeypairV2.
func (mg *KeypairV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KeypairV2.
func (mg *KeypairV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KeypairV2.
func (mg *KeypairV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KeypairV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KeypairV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this KeypairV2.
func (mg *KeypairV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this KeypairV2.
func (mg *KeypairV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KeypairV2.
func (mg *KeypairV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KeypairV2.
func (mg *KeypairV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KeypairV2.
func (mg *KeypairV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KeypairV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KeypairV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this KeypairV2.
func (mg *KeypairV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this KeypairV2.
func (mg *KeypairV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this QuotasetV2.
func (mg *QuotasetV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this QuotasetV2.
func (mg *QuotasetV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this QuotasetV2.
func (mg *QuotasetV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this QuotasetV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *QuotasetV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this QuotasetV2.
func (mg *QuotasetV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this QuotasetV2.
func (mg *QuotasetV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this QuotasetV2.
func (mg *QuotasetV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this QuotasetV2.
func (mg *QuotasetV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this QuotasetV2.
func (mg *QuotasetV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this QuotasetV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *QuotasetV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this QuotasetV2.
func (mg *QuotasetV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this QuotasetV2.
func (mg *QuotasetV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SecgroupV2.
func (mg *SecgroupV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SecgroupV2.
func (mg *SecgroupV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this SecgroupV2.
func (mg *SecgroupV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this SecgroupV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *SecgroupV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this SecgroupV2.
func (mg *SecgroupV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SecgroupV2.
func (mg *SecgroupV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SecgroupV2.
func (mg *SecgroupV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SecgroupV2.
func (mg *SecgroupV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this SecgroupV2.
func (mg *SecgroupV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this SecgroupV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *SecgroupV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this SecgroupV2.
func (mg *SecgroupV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SecgroupV2.
func (mg *SecgroupV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ServergroupV2.
func (mg *ServergroupV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ServergroupV2.
func (mg *ServergroupV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ServergroupV2.
func (mg *ServergroupV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ServergroupV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ServergroupV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this ServergroupV2.
func (mg *ServergroupV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ServergroupV2.
func (mg *ServergroupV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ServergroupV2.
func (mg *ServergroupV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ServergroupV2.
func (mg *ServergroupV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ServergroupV2.
func (mg *ServergroupV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ServergroupV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ServergroupV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this ServergroupV2.
func (mg *ServergroupV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ServergroupV2.
func (mg *ServergroupV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VolumeAttachV2.
func (mg *VolumeAttachV2) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VolumeAttachV2.
func (mg *VolumeAttachV2) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VolumeAttachV2.
func (mg *VolumeAttachV2) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VolumeAttachV2.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VolumeAttachV2) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VolumeAttachV2.
func (mg *VolumeAttachV2) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VolumeAttachV2.
func (mg *VolumeAttachV2) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VolumeAttachV2.
func (mg *VolumeAttachV2) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VolumeAttachV2.
func (mg *VolumeAttachV2) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VolumeAttachV2.
func (mg *VolumeAttachV2) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VolumeAttachV2.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VolumeAttachV2) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VolumeAttachV2.
func (mg *VolumeAttachV2) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VolumeAttachV2.
func (mg *VolumeAttachV2) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
