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

type CachePolicyCookieBehavior string

const (
	CachePolicyCookieBehavior_none      CachePolicyCookieBehavior = "none"
	CachePolicyCookieBehavior_whitelist CachePolicyCookieBehavior = "whitelist"
	CachePolicyCookieBehavior_allExcept CachePolicyCookieBehavior = "allExcept"
	CachePolicyCookieBehavior_all       CachePolicyCookieBehavior = "all"
)

type CachePolicyHeaderBehavior string

const (
	CachePolicyHeaderBehavior_none      CachePolicyHeaderBehavior = "none"
	CachePolicyHeaderBehavior_whitelist CachePolicyHeaderBehavior = "whitelist"
)

type CachePolicyQueryStringBehavior string

const (
	CachePolicyQueryStringBehavior_none      CachePolicyQueryStringBehavior = "none"
	CachePolicyQueryStringBehavior_whitelist CachePolicyQueryStringBehavior = "whitelist"
	CachePolicyQueryStringBehavior_allExcept CachePolicyQueryStringBehavior = "allExcept"
	CachePolicyQueryStringBehavior_all       CachePolicyQueryStringBehavior = "all"
)

type CachePolicyType string

const (
	CachePolicyType_managed CachePolicyType = "managed"
	CachePolicyType_custom  CachePolicyType = "custom"
)

type CertificateSource string

const (
	CertificateSource_cloudfront CertificateSource = "cloudfront"
	CertificateSource_iam        CertificateSource = "iam"
	CertificateSource_acm        CertificateSource = "acm"
)

type ContinuousDeploymentPolicyType string

const (
	ContinuousDeploymentPolicyType_SingleWeight ContinuousDeploymentPolicyType = "SingleWeight"
	ContinuousDeploymentPolicyType_SingleHeader ContinuousDeploymentPolicyType = "SingleHeader"
)

type EventType string

const (
	EventType_viewer_request  EventType = "viewer-request"
	EventType_viewer_response EventType = "viewer-response"
	EventType_origin_request  EventType = "origin-request"
	EventType_origin_response EventType = "origin-response"
)

type Format string

const (
	Format_URLEncoded Format = "URLEncoded"
)

type FrameOptionsList string

const (
	FrameOptionsList_DENY       FrameOptionsList = "DENY"
	FrameOptionsList_SAMEORIGIN FrameOptionsList = "SAMEORIGIN"
)

type FunctionRuntime string

const (
	FunctionRuntime_cloudfront_js_1_0 FunctionRuntime = "cloudfront-js-1.0"
)

type FunctionStage string

const (
	FunctionStage_DEVELOPMENT FunctionStage = "DEVELOPMENT"
	FunctionStage_LIVE        FunctionStage = "LIVE"
)

type GeoRestrictionType string

const (
	GeoRestrictionType_blacklist GeoRestrictionType = "blacklist"
	GeoRestrictionType_whitelist GeoRestrictionType = "whitelist"
	GeoRestrictionType_none      GeoRestrictionType = "none"
)

type HTTPVersion string

const (
	HTTPVersion_http1_1   HTTPVersion = "http1.1"
	HTTPVersion_http2     HTTPVersion = "http2"
	HTTPVersion_http3     HTTPVersion = "http3"
	HTTPVersion_http2and3 HTTPVersion = "http2and3"
)

type ICPRecordalStatus string

const (
	ICPRecordalStatus_APPROVED  ICPRecordalStatus = "APPROVED"
	ICPRecordalStatus_SUSPENDED ICPRecordalStatus = "SUSPENDED"
	ICPRecordalStatus_PENDING   ICPRecordalStatus = "PENDING"
)

type ItemSelection string

const (
	ItemSelection_none      ItemSelection = "none"
	ItemSelection_whitelist ItemSelection = "whitelist"
	ItemSelection_all       ItemSelection = "all"
)

type Method string

const (
	Method_GET     Method = "GET"
	Method_HEAD    Method = "HEAD"
	Method_POST    Method = "POST"
	Method_PUT     Method = "PUT"
	Method_PATCH   Method = "PATCH"
	Method_OPTIONS Method = "OPTIONS"
	Method_DELETE  Method = "DELETE"
)

type MinimumProtocolVersion string

const (
	MinimumProtocolVersion_SSLv3        MinimumProtocolVersion = "SSLv3"
	MinimumProtocolVersion_TLSv1        MinimumProtocolVersion = "TLSv1"
	MinimumProtocolVersion_TLSv1_2016   MinimumProtocolVersion = "TLSv1_2016"
	MinimumProtocolVersion_TLSv1_1_2016 MinimumProtocolVersion = "TLSv1.1_2016"
	MinimumProtocolVersion_TLSv1_2_2018 MinimumProtocolVersion = "TLSv1.2_2018"
	MinimumProtocolVersion_TLSv1_2_2019 MinimumProtocolVersion = "TLSv1.2_2019"
	MinimumProtocolVersion_TLSv1_2_2021 MinimumProtocolVersion = "TLSv1.2_2021"
)

type OriginAccessControlOriginTypes string

const (
	OriginAccessControlOriginTypes_s3 OriginAccessControlOriginTypes = "s3"
)

type OriginAccessControlSigningBehaviors string

const (
	OriginAccessControlSigningBehaviors_never       OriginAccessControlSigningBehaviors = "never"
	OriginAccessControlSigningBehaviors_always      OriginAccessControlSigningBehaviors = "always"
	OriginAccessControlSigningBehaviors_no_override OriginAccessControlSigningBehaviors = "no-override"
)

type OriginAccessControlSigningProtocols string

const (
	OriginAccessControlSigningProtocols_sigv4 OriginAccessControlSigningProtocols = "sigv4"
)

type OriginProtocolPolicy string

const (
	OriginProtocolPolicy_http_only    OriginProtocolPolicy = "http-only"
	OriginProtocolPolicy_match_viewer OriginProtocolPolicy = "match-viewer"
	OriginProtocolPolicy_https_only   OriginProtocolPolicy = "https-only"
)

type OriginRequestPolicyCookieBehavior string

const (
	OriginRequestPolicyCookieBehavior_none      OriginRequestPolicyCookieBehavior = "none"
	OriginRequestPolicyCookieBehavior_whitelist OriginRequestPolicyCookieBehavior = "whitelist"
	OriginRequestPolicyCookieBehavior_all       OriginRequestPolicyCookieBehavior = "all"
)

type OriginRequestPolicyHeaderBehavior string

const (
	OriginRequestPolicyHeaderBehavior_none                            OriginRequestPolicyHeaderBehavior = "none"
	OriginRequestPolicyHeaderBehavior_whitelist                       OriginRequestPolicyHeaderBehavior = "whitelist"
	OriginRequestPolicyHeaderBehavior_allViewer                       OriginRequestPolicyHeaderBehavior = "allViewer"
	OriginRequestPolicyHeaderBehavior_allViewerAndWhitelistCloudFront OriginRequestPolicyHeaderBehavior = "allViewerAndWhitelistCloudFront"
)

type OriginRequestPolicyQueryStringBehavior string

const (
	OriginRequestPolicyQueryStringBehavior_none      OriginRequestPolicyQueryStringBehavior = "none"
	OriginRequestPolicyQueryStringBehavior_whitelist OriginRequestPolicyQueryStringBehavior = "whitelist"
	OriginRequestPolicyQueryStringBehavior_all       OriginRequestPolicyQueryStringBehavior = "all"
)

type OriginRequestPolicyType string

const (
	OriginRequestPolicyType_managed OriginRequestPolicyType = "managed"
	OriginRequestPolicyType_custom  OriginRequestPolicyType = "custom"
)

type PriceClass string

const (
	PriceClass_PriceClass_100 PriceClass = "PriceClass_100"
	PriceClass_PriceClass_200 PriceClass = "PriceClass_200"
	PriceClass_PriceClass_All PriceClass = "PriceClass_All"
)

type RealtimeMetricsSubscriptionStatus string

const (
	RealtimeMetricsSubscriptionStatus_Enabled  RealtimeMetricsSubscriptionStatus = "Enabled"
	RealtimeMetricsSubscriptionStatus_Disabled RealtimeMetricsSubscriptionStatus = "Disabled"
)

type ReferrerPolicyList string

const (
	ReferrerPolicyList_no_referrer                     ReferrerPolicyList = "no-referrer"
	ReferrerPolicyList_no_referrer_when_downgrade      ReferrerPolicyList = "no-referrer-when-downgrade"
	ReferrerPolicyList_origin                          ReferrerPolicyList = "origin"
	ReferrerPolicyList_origin_when_cross_origin        ReferrerPolicyList = "origin-when-cross-origin"
	ReferrerPolicyList_same_origin                     ReferrerPolicyList = "same-origin"
	ReferrerPolicyList_strict_origin                   ReferrerPolicyList = "strict-origin"
	ReferrerPolicyList_strict_origin_when_cross_origin ReferrerPolicyList = "strict-origin-when-cross-origin"
	ReferrerPolicyList_unsafe_url                      ReferrerPolicyList = "unsafe-url"
)

type ResponseHeadersPolicyAccessControlAllowMethodsValues string

const (
	ResponseHeadersPolicyAccessControlAllowMethodsValues_GET     ResponseHeadersPolicyAccessControlAllowMethodsValues = "GET"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_POST    ResponseHeadersPolicyAccessControlAllowMethodsValues = "POST"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_OPTIONS ResponseHeadersPolicyAccessControlAllowMethodsValues = "OPTIONS"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_PUT     ResponseHeadersPolicyAccessControlAllowMethodsValues = "PUT"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_DELETE  ResponseHeadersPolicyAccessControlAllowMethodsValues = "DELETE"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_PATCH   ResponseHeadersPolicyAccessControlAllowMethodsValues = "PATCH"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_HEAD    ResponseHeadersPolicyAccessControlAllowMethodsValues = "HEAD"
	ResponseHeadersPolicyAccessControlAllowMethodsValues_ALL     ResponseHeadersPolicyAccessControlAllowMethodsValues = "ALL"
)

type ResponseHeadersPolicyType string

const (
	ResponseHeadersPolicyType_managed ResponseHeadersPolicyType = "managed"
	ResponseHeadersPolicyType_custom  ResponseHeadersPolicyType = "custom"
)

type SSLProtocol string

const (
	SSLProtocol_SSLv3   SSLProtocol = "SSLv3"
	SSLProtocol_TLSv1   SSLProtocol = "TLSv1"
	SSLProtocol_TLSv1_1 SSLProtocol = "TLSv1.1"
	SSLProtocol_TLSv1_2 SSLProtocol = "TLSv1.2"
)

type SSLSupportMethod string

const (
	SSLSupportMethod_sni_only  SSLSupportMethod = "sni-only"
	SSLSupportMethod_vip       SSLSupportMethod = "vip"
	SSLSupportMethod_static_ip SSLSupportMethod = "static-ip"
)

type ViewerProtocolPolicy string

const (
	ViewerProtocolPolicy_allow_all         ViewerProtocolPolicy = "allow-all"
	ViewerProtocolPolicy_https_only        ViewerProtocolPolicy = "https-only"
	ViewerProtocolPolicy_redirect_to_https ViewerProtocolPolicy = "redirect-to-https"
)
