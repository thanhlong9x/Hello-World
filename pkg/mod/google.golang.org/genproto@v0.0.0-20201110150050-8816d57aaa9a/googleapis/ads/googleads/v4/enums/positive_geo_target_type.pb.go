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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/ads/googleads/v4/enums/positive_geo_target_type.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The possible positive geo target types.
type PositiveGeoTargetTypeEnum_PositiveGeoTargetType int32

const (
	// Not specified.
	PositiveGeoTargetTypeEnum_UNSPECIFIED PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 0
	// The value is unknown in this version.
	PositiveGeoTargetTypeEnum_UNKNOWN PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 1
	// Specifies that an ad is triggered if the user is in,
	// or shows interest in, advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE_OR_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 5
	// Specifies that an ad is triggered if the user
	// searches for advertiser's targeted locations.
	// This can only be used with Search and standard
	// Shopping campaigns.
	PositiveGeoTargetTypeEnum_SEARCH_INTEREST PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 6
	// Specifies that an ad is triggered if the user is in
	// or regularly in advertiser's targeted locations.
	PositiveGeoTargetTypeEnum_PRESENCE PositiveGeoTargetTypeEnum_PositiveGeoTargetType = 7
)

// Enum value maps for PositiveGeoTargetTypeEnum_PositiveGeoTargetType.
var (
	PositiveGeoTargetTypeEnum_PositiveGeoTargetType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		5: "PRESENCE_OR_INTEREST",
		6: "SEARCH_INTEREST",
		7: "PRESENCE",
	}
	PositiveGeoTargetTypeEnum_PositiveGeoTargetType_value = map[string]int32{
		"UNSPECIFIED":          0,
		"UNKNOWN":              1,
		"PRESENCE_OR_INTEREST": 5,
		"SEARCH_INTEREST":      6,
		"PRESENCE":             7,
	}
)

func (x PositiveGeoTargetTypeEnum_PositiveGeoTargetType) Enum() *PositiveGeoTargetTypeEnum_PositiveGeoTargetType {
	p := new(PositiveGeoTargetTypeEnum_PositiveGeoTargetType)
	*p = x
	return p
}

func (x PositiveGeoTargetTypeEnum_PositiveGeoTargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PositiveGeoTargetTypeEnum_PositiveGeoTargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_enumTypes[0].Descriptor()
}

func (PositiveGeoTargetTypeEnum_PositiveGeoTargetType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_enumTypes[0]
}

func (x PositiveGeoTargetTypeEnum_PositiveGeoTargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PositiveGeoTargetTypeEnum_PositiveGeoTargetType.Descriptor instead.
func (PositiveGeoTargetTypeEnum_PositiveGeoTargetType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible positive geo target types.
type PositiveGeoTargetTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PositiveGeoTargetTypeEnum) Reset() {
	*x = PositiveGeoTargetTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositiveGeoTargetTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositiveGeoTargetTypeEnum) ProtoMessage() {}

func (x *PositiveGeoTargetTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositiveGeoTargetTypeEnum.ProtoReflect.Descriptor instead.
func (*PositiveGeoTargetTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v4_enums_positive_geo_target_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x65, 0x6f, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x19,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x72, 0x0a, 0x15, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x47, 0x65, 0x6f, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4f, 0x52, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x52, 0x45, 0x53, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x07, 0x42, 0xef, 0x01,
	0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x42, 0x1a, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x47, 0x65, 0x6f,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x34, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x34, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x34, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescData = file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDesc
)

func file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDescData
}

var file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_goTypes = []interface{}{
	(PositiveGeoTargetTypeEnum_PositiveGeoTargetType)(0), // 0: google.ads.googleads.v4.enums.PositiveGeoTargetTypeEnum.PositiveGeoTargetType
	(*PositiveGeoTargetTypeEnum)(nil),                    // 1: google.ads.googleads.v4.enums.PositiveGeoTargetTypeEnum
}
var file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_init() }
func file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_init() {
	if File_google_ads_googleads_v4_enums_positive_geo_target_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositiveGeoTargetTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v4_enums_positive_geo_target_type_proto = out.File
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_rawDesc = nil
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_goTypes = nil
	file_google_ads_googleads_v4_enums_positive_geo_target_type_proto_depIdxs = nil
}
