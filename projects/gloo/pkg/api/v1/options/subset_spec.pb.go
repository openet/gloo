// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/subset_spec.proto

package options

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FallbackPolicy int32

const (
	// Any cluster endpoint may be returned (default)
	FallbackPolicy_ANY_ENDPOINT FallbackPolicy = 0
	// Load balancing over the endpoints matching the values from the default_subset field
	FallbackPolicy_DEFAULT_SUBSET FallbackPolicy = 1
	// A result equivalent to no healthy hosts is reported
	FallbackPolicy_NO_FALLBACK FallbackPolicy = 2
)

// Enum value maps for FallbackPolicy.
var (
	FallbackPolicy_name = map[int32]string{
		0: "ANY_ENDPOINT",
		1: "DEFAULT_SUBSET",
		2: "NO_FALLBACK",
	}
	FallbackPolicy_value = map[string]int32{
		"ANY_ENDPOINT":   0,
		"DEFAULT_SUBSET": 1,
		"NO_FALLBACK":    2,
	}
)

func (x FallbackPolicy) Enum() *FallbackPolicy {
	p := new(FallbackPolicy)
	*p = x
	return p
}

func (x FallbackPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FallbackPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_enumTypes[0].Descriptor()
}

func (FallbackPolicy) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_enumTypes[0]
}

func (x FallbackPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FallbackPolicy.Descriptor instead.
func (FallbackPolicy) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescGZIP(), []int{0}
}

// See envoy docs for details:
// https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#config-cluster-v3-cluster-lbsubsetconfig
type SubsetSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the set of subsets of the upstream
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// The behavior used when no endpoint subset matches the selected route’s metadata
	// The default value is ANY_ENDPOINT
	FallbackPolicy FallbackPolicy `protobuf:"varint,2,opt,name=fallbackPolicy,proto3,enum=options.gloo.solo.io.FallbackPolicy" json:"fallbackPolicy,omitempty"`
	// Specifies the default subset of endpoints used during fallback if fallback_policy is DEFAULT_SUBSET
	DefaultSubset *Subset `protobuf:"bytes,3,opt,name=default_subset,json=defaultSubset,proto3" json:"default_subset,omitempty"`
}

func (x *SubsetSpec) Reset() {
	*x = SubsetSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsetSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsetSpec) ProtoMessage() {}

func (x *SubsetSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsetSpec.ProtoReflect.Descriptor instead.
func (*SubsetSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescGZIP(), []int{0}
}

func (x *SubsetSpec) GetSelectors() []*Selector {
	if x != nil {
		return x.Selectors
	}
	return nil
}

func (x *SubsetSpec) GetFallbackPolicy() FallbackPolicy {
	if x != nil {
		return x.FallbackPolicy
	}
	return FallbackPolicy_ANY_ENDPOINT
}

func (x *SubsetSpec) GetDefaultSubset() *Subset {
	if x != nil {
		return x.DefaultSubset
	}
	return nil
}

type Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A subset is created for each unique combination of key and value
	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	// Selects a mode of operation in which each subset has only one host. Default is false.
	SingleHostPerSubset bool `protobuf:"varint,2,opt,name=single_host_per_subset,json=singleHostPerSubset,proto3" json:"single_host_per_subset,omitempty"`
}

func (x *Selector) Reset() {
	*x = Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selector) ProtoMessage() {}

func (x *Selector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selector.ProtoReflect.Descriptor instead.
func (*Selector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescGZIP(), []int{1}
}

func (x *Selector) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *Selector) GetSingleHostPerSubset() bool {
	if x != nil {
		return x.SingleHostPerSubset
	}
	return false
}

type Subset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Any host that matches all key/value pairs is part of this subset
	Values map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Subset) Reset() {
	*x = Subset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subset) ProtoMessage() {}

func (x *Subset) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subset.ProtoReflect.Descriptor instead.
func (*Subset) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescGZIP(), []int{2}
}

func (x *Subset) GetValues() map[string]string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDesc = []byte{
	0x0a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x3c, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x4c, 0x0a, 0x0e, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x43, 0x0a,
	0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x65, 0x74, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x62, 0x73,
	0x65, 0x74, 0x22, 0x53, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x13, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x73,
	0x65, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x65, 0x74,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x47, 0x0a, 0x0e, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e,
	0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x55, 0x42, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x46, 0x41,
	0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x42, 0x46, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_goTypes = []interface{}{
	(FallbackPolicy)(0), // 0: options.gloo.solo.io.FallbackPolicy
	(*SubsetSpec)(nil),  // 1: options.gloo.solo.io.SubsetSpec
	(*Selector)(nil),    // 2: options.gloo.solo.io.Selector
	(*Subset)(nil),      // 3: options.gloo.solo.io.Subset
	nil,                 // 4: options.gloo.solo.io.Subset.ValuesEntry
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_depIdxs = []int32{
	2, // 0: options.gloo.solo.io.SubsetSpec.selectors:type_name -> options.gloo.solo.io.Selector
	0, // 1: options.gloo.solo.io.SubsetSpec.fallbackPolicy:type_name -> options.gloo.solo.io.FallbackPolicy
	3, // 2: options.gloo.solo.io.SubsetSpec.default_subset:type_name -> options.gloo.solo.io.Subset
	4, // 3: options.gloo.solo.io.Subset.values:type_name -> options.gloo.solo.io.Subset.ValuesEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsetSpec); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selector); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subset); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_subset_spec_proto_depIdxs = nil
}
