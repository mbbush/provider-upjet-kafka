// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ACLInitParameters struct {
}

type ACLObservation struct {
	ACLHost *string `json:"aclHost,omitempty" tf:"acl_host,omitempty"`

	ACLOperation *string `json:"aclOperation,omitempty" tf:"acl_operation,omitempty"`

	ACLPermissionType *string `json:"aclPermissionType,omitempty" tf:"acl_permission_type,omitempty"`

	ACLPrincipal *string `json:"aclPrincipal,omitempty" tf:"acl_principal,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name,omitempty"`

	ResourcePatternTypeFilter *string `json:"resourcePatternTypeFilter,omitempty" tf:"resource_pattern_type_filter,omitempty"`

	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type ACLParameters struct {

	// +kubebuilder:validation:Required
	ACLHost *string `json:"aclHost" tf:"acl_host,omitempty"`

	// +kubebuilder:validation:Required
	ACLOperation *string `json:"aclOperation" tf:"acl_operation,omitempty"`

	// +kubebuilder:validation:Required
	ACLPermissionType *string `json:"aclPermissionType" tf:"acl_permission_type,omitempty"`

	// +kubebuilder:validation:Required
	ACLPrincipal *string `json:"aclPrincipal" tf:"acl_principal,omitempty"`

	// The name of the resource
	// +kubebuilder:validation:Required
	ResourceName *string `json:"resourceName" tf:"resource_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourcePatternTypeFilter *string `json:"resourcePatternTypeFilter,omitempty" tf:"resource_pattern_type_filter,omitempty"`

	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`
}

// ACLSpec defines the desired state of ACL
type ACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ACLParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ACLInitParameters `json:"initProvider,omitempty"`
}

// ACLStatus defines the observed state of ACL.
type ACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ACL is the Schema for the ACLs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,kafka}
type ACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ACLSpec   `json:"spec"`
	Status            ACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ACLList contains a list of ACLs
type ACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ACL `json:"items"`
}

// Repository type metadata.
var (
	ACL_Kind             = "ACL"
	ACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ACL_Kind}.String()
	ACL_KindAPIVersion   = ACL_Kind + "." + CRDGroupVersion.String()
	ACL_GroupVersionKind = CRDGroupVersion.WithKind(ACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ACL{}, &ACLList{})
}
