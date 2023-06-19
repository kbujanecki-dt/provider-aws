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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AccessPointParameters defines the desired state of AccessPoint
type AccessPointParameters struct {
	// Region is which region the AccessPoint will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The Amazon Web Services account ID for the account that owns the specified
	// access point.
	// +kubebuilder:validation:Required
	AccountID *string `json:"accountID"`
	// The name of the bucket that you want to associate this access point with.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you
	// must specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services
	// SDK and CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>.
	// For example, to access the bucket reports through outpost my-outpost owned
	// by account 123456789012 in Region us-west-2, use the URL encoding of arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket"`
	// The Amazon Web Services account ID associated with the S3 bucket associated
	// with this access point.
	BucketAccountID *string `json:"bucketAccountID,omitempty"`
	// The PublicAccessBlock configuration that you want to apply to the access
	// point.
	PublicAccessBlockConfiguration *PublicAccessBlockConfiguration `json:"publicAccessBlockConfiguration,omitempty"`
	// If you include this field, Amazon S3 restricts access to this access point
	// to requests from the specified virtual private cloud (VPC).
	//
	// This is required for creating an access point for Amazon S3 on Outposts buckets.
	VPCConfiguration            *VPCConfiguration `json:"vpcConfiguration,omitempty"`
	CustomAccessPointParameters `json:",inline"`
}

// AccessPointSpec defines the desired state of AccessPoint
type AccessPointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AccessPointParameters `json:"forProvider"`
}

// AccessPointObservation defines the observed state of AccessPoint
type AccessPointObservation struct {
	// The ARN of the access point.
	//
	// This is only supported by Amazon S3 on Outposts.
	AccessPointARN *string `json:"accessPointARN,omitempty"`
	// The name or alias of the access point.
	Alias *string `json:"alias,omitempty"`
}

// AccessPointStatus defines the observed state of AccessPoint.
type AccessPointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPoint is the Schema for the AccessPoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessPointSpec   `json:"spec"`
	Status            AccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointList contains a list of AccessPoints
type AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPoint `json:"items"`
}

// Repository type metadata.
var (
	AccessPointKind             = "AccessPoint"
	AccessPointGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPointKind}.String()
	AccessPointKindAPIVersion   = AccessPointKind + "." + GroupVersion.String()
	AccessPointGroupVersionKind = GroupVersion.WithKind(AccessPointKind)
)

func init() {
	SchemeBuilder.Register(&AccessPoint{}, &AccessPointList{})
}
