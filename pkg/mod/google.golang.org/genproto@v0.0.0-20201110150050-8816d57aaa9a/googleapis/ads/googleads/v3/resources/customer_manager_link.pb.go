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
// source: google/ads/googleads/v3/resources/customer_manager_link.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Represents customer-manager link relationship.
type CustomerManagerLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Name of the resource.
	// CustomerManagerLink resource names have the form:
	// `customers/{customer_id}/customerManagerLinks/{manager_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The manager customer linked to the customer.
	ManagerCustomer *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=manager_customer,json=managerCustomer,proto3" json:"manager_customer,omitempty"`
	// Output only. ID of the customer-manager link. This field is read only.
	ManagerLinkId *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// Status of the link between the customer and the manager.
	Status enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
}

func (x *CustomerManagerLink) Reset() {
	*x = CustomerManagerLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_resources_customer_manager_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerManagerLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerManagerLink) ProtoMessage() {}

func (x *CustomerManagerLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_resources_customer_manager_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerManagerLink.ProtoReflect.Descriptor instead.
func (*CustomerManagerLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerManagerLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerManagerLink) GetManagerCustomer() *wrapperspb.StringValue {
	if x != nil {
		return x.ManagerCustomer
	}
	return nil
}

func (x *CustomerManagerLink) GetManagerLinkId() *wrapperspb.Int64Value {
	if x != nil {
		return x.ManagerLinkId
	}
	return nil
}

func (x *CustomerManagerLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if x != nil {
		return x.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

var File_google_ads_googleads_v3_resources_customer_manager_link_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x03, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x59, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x64,
	0x12, 0x5e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3a, 0x74, 0xea, 0x41, 0x71, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x41, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x7d, 0x42, 0x85, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x18, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02,
	0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescData = file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDesc
)

func file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDescData
}

var file_google_ads_googleads_v3_resources_customer_manager_link_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_resources_customer_manager_link_proto_goTypes = []interface{}{
	(*CustomerManagerLink)(nil),                        // 0: google.ads.googleads.v3.resources.CustomerManagerLink
	(*wrapperspb.StringValue)(nil),                     // 1: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),                      // 2: google.protobuf.Int64Value
	(enums.ManagerLinkStatusEnum_ManagerLinkStatus)(0), // 3: google.ads.googleads.v3.enums.ManagerLinkStatusEnum.ManagerLinkStatus
}
var file_google_ads_googleads_v3_resources_customer_manager_link_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v3.resources.CustomerManagerLink.manager_customer:type_name -> google.protobuf.StringValue
	2, // 1: google.ads.googleads.v3.resources.CustomerManagerLink.manager_link_id:type_name -> google.protobuf.Int64Value
	3, // 2: google.ads.googleads.v3.resources.CustomerManagerLink.status:type_name -> google.ads.googleads.v3.enums.ManagerLinkStatusEnum.ManagerLinkStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_resources_customer_manager_link_proto_init() }
func file_google_ads_googleads_v3_resources_customer_manager_link_proto_init() {
	if File_google_ads_googleads_v3_resources_customer_manager_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_resources_customer_manager_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerManagerLink); i {
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
			RawDescriptor: file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_resources_customer_manager_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_resources_customer_manager_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v3_resources_customer_manager_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_resources_customer_manager_link_proto = out.File
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_rawDesc = nil
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_goTypes = nil
	file_google_ads_googleads_v3_resources_customer_manager_link_proto_depIdxs = nil
}
