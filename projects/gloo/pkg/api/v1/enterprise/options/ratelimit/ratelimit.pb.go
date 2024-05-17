// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/ratelimit/ratelimit.proto

package ratelimit

import (
	reflect "reflect"
	sync "sync"

	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Basic rate-limiting API
type IngressRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizedLimits *v1alpha1.RateLimit `protobuf:"bytes,1,opt,name=authorized_limits,json=authorizedLimits,proto3" json:"authorized_limits,omitempty"`
	AnonymousLimits  *v1alpha1.RateLimit `protobuf:"bytes,2,opt,name=anonymous_limits,json=anonymousLimits,proto3" json:"anonymous_limits,omitempty"`
}

func (x *IngressRateLimit) Reset() {
	*x = IngressRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressRateLimit) ProtoMessage() {}

func (x *IngressRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressRateLimit.ProtoReflect.Descriptor instead.
func (*IngressRateLimit) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{0}
}

func (x *IngressRateLimit) GetAuthorizedLimits() *v1alpha1.RateLimit {
	if x != nil {
		return x.AuthorizedLimits
	}
	return nil
}

func (x *IngressRateLimit) GetAnonymousLimits() *v1alpha1.RateLimit {
	if x != nil {
		return x.AnonymousLimits
	}
	return nil
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatelimitServerRef *core.ResourceRef  `protobuf:"bytes,1,opt,name=ratelimit_server_ref,json=ratelimitServerRef,proto3" json:"ratelimit_server_ref,omitempty"`
	RequestTimeout     *duration.Duration `protobuf:"bytes,2,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	DenyOnFail         bool               `protobuf:"varint,3,opt,name=deny_on_fail,json=denyOnFail,proto3" json:"deny_on_fail,omitempty"`
	// Set this to true to return Envoy's X-RateLimit headers to the downstream.
	// reference docs here: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/ratelimit/v3/rate_limit.proto.html#envoy-v3-api-field-extensions-filters-http-ratelimit-v3-ratelimit-enable-x-ratelimit-headers
	EnableXRatelimitHeaders bool `protobuf:"varint,4,opt,name=enable_x_ratelimit_headers,json=enableXRatelimitHeaders,proto3" json:"enable_x_ratelimit_headers,omitempty"`
	// Set this is set to true if you would like to rate limit traffic before applying external auth to it.
	// *Note*: When this is true, you will lose some features like being able to rate limit a request based on its auth state
	RateLimitBeforeAuth bool `protobuf:"varint,9,opt,name=rate_limit_before_auth,json=rateLimitBeforeAuth,proto3" json:"rate_limit_before_auth,omitempty"`
	// Types that are assignable to ServiceType:
	//
	//	*Settings_GrpcService
	ServiceType isSettings_ServiceType `protobuf_oneof:"service_type"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{1}
}

func (x *Settings) GetRatelimitServerRef() *core.ResourceRef {
	if x != nil {
		return x.RatelimitServerRef
	}
	return nil
}

func (x *Settings) GetRequestTimeout() *duration.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *Settings) GetDenyOnFail() bool {
	if x != nil {
		return x.DenyOnFail
	}
	return false
}

func (x *Settings) GetEnableXRatelimitHeaders() bool {
	if x != nil {
		return x.EnableXRatelimitHeaders
	}
	return false
}

func (x *Settings) GetRateLimitBeforeAuth() bool {
	if x != nil {
		return x.RateLimitBeforeAuth
	}
	return false
}

func (m *Settings) GetServiceType() isSettings_ServiceType {
	if m != nil {
		return m.ServiceType
	}
	return nil
}

func (x *Settings) GetGrpcService() *GrpcService {
	if x, ok := x.GetServiceType().(*Settings_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

type isSettings_ServiceType interface {
	isSettings_ServiceType()
}

type Settings_GrpcService struct {
	// Optional gRPC settings used when calling the ratelimit server.
	GrpcService *GrpcService `protobuf:"bytes,10,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

func (*Settings_GrpcService) isSettings_ServiceType() {}

type GrpcService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set the authority header when calling the gRPC service.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *GrpcService) Reset() {
	*x = GrpcService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcService) ProtoMessage() {}

func (x *GrpcService) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcService.ProtoReflect.Descriptor instead.
func (*GrpcService) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcService) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

// API based on Envoy's rate-limit service API. (reference here: https://github.com/lyft/ratelimit#configuration)
// Sample configuration below:
//
// descriptors:
//   - key: account_id
//     descriptors:
//   - key: plan
//     value: BASIC
//     rateLimit:
//     requestsPerUnit: 1
//     unit: MINUTE
//   - key: plan
//     value: PLUS
//     rateLimit:
//     requestsPerUnit: 20
//     unit: MINUTE
type ServiceSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Descriptors    []*v1alpha1.Descriptor    `protobuf:"bytes,1,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	SetDescriptors []*v1alpha1.SetDescriptor `protobuf:"bytes,2,rep,name=set_descriptors,json=setDescriptors,proto3" json:"set_descriptors,omitempty"`
}

func (x *ServiceSettings) Reset() {
	*x = ServiceSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSettings) ProtoMessage() {}

func (x *ServiceSettings) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSettings.ProtoReflect.Descriptor instead.
func (*ServiceSettings) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{3}
}

func (x *ServiceSettings) GetDescriptors() []*v1alpha1.Descriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *ServiceSettings) GetSetDescriptors() []*v1alpha1.SetDescriptor {
	if x != nil {
		return x.SetDescriptors
	}
	return nil
}

// A list of references to `RateLimitConfig` resources.
// Each resource represents a rate limit policy that will be independently enforced.
type RateLimitConfigRefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refs []*RateLimitConfigRef `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (x *RateLimitConfigRefs) Reset() {
	*x = RateLimitConfigRefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitConfigRefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitConfigRefs) ProtoMessage() {}

func (x *RateLimitConfigRefs) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitConfigRefs.ProtoReflect.Descriptor instead.
func (*RateLimitConfigRefs) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{4}
}

func (x *RateLimitConfigRefs) GetRefs() []*RateLimitConfigRef {
	if x != nil {
		return x.Refs
	}
	return nil
}

// A reference to a `RateLimitConfig` resource.
type RateLimitConfigRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *RateLimitConfigRef) Reset() {
	*x = RateLimitConfigRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitConfigRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitConfigRef) ProtoMessage() {}

func (x *RateLimitConfigRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitConfigRef.ProtoReflect.Descriptor instead.
func (*RateLimitConfigRef) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{5}
}

func (x *RateLimitConfigRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitConfigRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Use this field if you want to inline the Envoy rate limits for this VirtualHost.
// Note that this does not configure the rate limit server. If you are running Gloo Enterprise, you need to
// specify the server configuration via the appropriate field in the Gloo `Settings` resource. If you are
// running a custom rate limit server you need to configure it yourself.
type RateLimitVhostExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define individual rate limits here. Each rate limit will be evaluated, if any rate limit
	// would be throttled, the entire request returns a 429 (gets throttled)
	RateLimits []*v1alpha1.RateLimitActions `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
}

func (x *RateLimitVhostExtension) Reset() {
	*x = RateLimitVhostExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitVhostExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitVhostExtension) ProtoMessage() {}

func (x *RateLimitVhostExtension) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitVhostExtension.ProtoReflect.Descriptor instead.
func (*RateLimitVhostExtension) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{6}
}

func (x *RateLimitVhostExtension) GetRateLimits() []*v1alpha1.RateLimitActions {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

// Use this field if you want to inline the Envoy rate limits for this Route.
// Note that this does not configure the rate limit server. If you are running Gloo Enterprise, you need to
// specify the server configuration via the appropriate field in the Gloo `Settings` resource. If you are
// running a custom rate limit server you need to configure it yourself.
type RateLimitRouteExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to include rate limits as defined on the VirtualHost in addition to rate limits on the Route.
	IncludeVhRateLimits bool `protobuf:"varint,1,opt,name=include_vh_rate_limits,json=includeVhRateLimits,proto3" json:"include_vh_rate_limits,omitempty"`
	// Define individual rate limits here. Each rate limit will be evaluated, if any rate limit
	// would be throttled, the entire request returns a 429 (gets throttled)
	RateLimits []*v1alpha1.RateLimitActions `protobuf:"bytes,2,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
}

func (x *RateLimitRouteExtension) Reset() {
	*x = RateLimitRouteExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitRouteExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitRouteExtension) ProtoMessage() {}

func (x *RateLimitRouteExtension) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitRouteExtension.ProtoReflect.Descriptor instead.
func (*RateLimitRouteExtension) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP(), []int{7}
}

func (x *RateLimitRouteExtension) GetIncludeVhRateLimits() bool {
	if x != nil {
		return x.IncludeVhRateLimits
	}
	return false
}

func (x *RateLimitRouteExtension) GetRateLimits() []*v1alpha1.RateLimitActions {
	if x != nil {
		return x.RateLimits
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDesc = []byte{
	0x0a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x61, 0x74,
	0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x2d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x4d, 0x0a, 0x11, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x10, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x4b, 0x0a, 0x10, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x91, 0x03, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x14, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x12, 0x72, 0x61,
	0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x6f, 0x6e, 0x5f,
	0x66, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x6e, 0x79,
	0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x58, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x50, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x0b, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x4d, 0x0a, 0x0f, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52,
	0x0e, 0x73, 0x65, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x22,
	0x5d, 0x0a, 0x13, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x66, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x72, 0x65, 0x66, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x52, 0x04, 0x72, 0x65, 0x66, 0x73, 0x22, 0x46,
	0x0a, 0x12, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x63, 0x0a, 0x17, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x56, 0x68, 0x6f, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x48, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0a, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x17,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x5f, 0x76, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x56, 0x68, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x0b,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x42, 0x5b, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01,
	0xd0, 0xf5, 0x04, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_goTypes = []interface{}{
	(*IngressRateLimit)(nil),          // 0: ratelimit.options.gloo.solo.io.IngressRateLimit
	(*Settings)(nil),                  // 1: ratelimit.options.gloo.solo.io.Settings
	(*GrpcService)(nil),               // 2: ratelimit.options.gloo.solo.io.GrpcService
	(*ServiceSettings)(nil),           // 3: ratelimit.options.gloo.solo.io.ServiceSettings
	(*RateLimitConfigRefs)(nil),       // 4: ratelimit.options.gloo.solo.io.RateLimitConfigRefs
	(*RateLimitConfigRef)(nil),        // 5: ratelimit.options.gloo.solo.io.RateLimitConfigRef
	(*RateLimitVhostExtension)(nil),   // 6: ratelimit.options.gloo.solo.io.RateLimitVhostExtension
	(*RateLimitRouteExtension)(nil),   // 7: ratelimit.options.gloo.solo.io.RateLimitRouteExtension
	(*v1alpha1.RateLimit)(nil),        // 8: ratelimit.api.solo.io.RateLimit
	(*core.ResourceRef)(nil),          // 9: core.solo.io.ResourceRef
	(*duration.Duration)(nil),         // 10: google.protobuf.Duration
	(*v1alpha1.Descriptor)(nil),       // 11: ratelimit.api.solo.io.Descriptor
	(*v1alpha1.SetDescriptor)(nil),    // 12: ratelimit.api.solo.io.SetDescriptor
	(*v1alpha1.RateLimitActions)(nil), // 13: ratelimit.api.solo.io.RateLimitActions
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_depIdxs = []int32{
	8,  // 0: ratelimit.options.gloo.solo.io.IngressRateLimit.authorized_limits:type_name -> ratelimit.api.solo.io.RateLimit
	8,  // 1: ratelimit.options.gloo.solo.io.IngressRateLimit.anonymous_limits:type_name -> ratelimit.api.solo.io.RateLimit
	9,  // 2: ratelimit.options.gloo.solo.io.Settings.ratelimit_server_ref:type_name -> core.solo.io.ResourceRef
	10, // 3: ratelimit.options.gloo.solo.io.Settings.request_timeout:type_name -> google.protobuf.Duration
	2,  // 4: ratelimit.options.gloo.solo.io.Settings.grpc_service:type_name -> ratelimit.options.gloo.solo.io.GrpcService
	11, // 5: ratelimit.options.gloo.solo.io.ServiceSettings.descriptors:type_name -> ratelimit.api.solo.io.Descriptor
	12, // 6: ratelimit.options.gloo.solo.io.ServiceSettings.set_descriptors:type_name -> ratelimit.api.solo.io.SetDescriptor
	5,  // 7: ratelimit.options.gloo.solo.io.RateLimitConfigRefs.refs:type_name -> ratelimit.options.gloo.solo.io.RateLimitConfigRef
	13, // 8: ratelimit.options.gloo.solo.io.RateLimitVhostExtension.rate_limits:type_name -> ratelimit.api.solo.io.RateLimitActions
	13, // 9: ratelimit.options.gloo.solo.io.RateLimitRouteExtension.rate_limits:type_name -> ratelimit.api.solo.io.RateLimitActions
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressRateLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSettings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitConfigRefs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitConfigRef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitVhostExtension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitRouteExtension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Settings_GrpcService)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_ratelimit_ratelimit_proto_depIdxs = nil
}
