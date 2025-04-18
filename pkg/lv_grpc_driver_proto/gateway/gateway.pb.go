// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: gateway/gateway.proto

package gateway

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GateWayInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Env           string                 `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	GwId          int64                  `protobuf:"varint,3,opt,name=gwId,proto3" json:"gwId,omitempty"`
	LocalKey      string                 `protobuf:"bytes,4,opt,name=localKey,proto3" json:"localKey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GateWayInfoResponse) Reset() {
	*x = GateWayInfoResponse{}
	mi := &file_gateway_gateway_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GateWayInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateWayInfoResponse) ProtoMessage() {}

func (x *GateWayInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateWayInfoResponse.ProtoReflect.Descriptor instead.
func (*GateWayInfoResponse) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GateWayInfoResponse) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *GateWayInfoResponse) GetGwId() int64 {
	if x != nil {
		return x.GwId
	}
	return 0
}

func (x *GateWayInfoResponse) GetLocalKey() string {
	if x != nil {
		return x.LocalKey
	}
	return ""
}

var File_gateway_gateway_proto protoreflect.FileDescriptor

var file_gateway_gateway_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a,
	0x13, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x77, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x77, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4b, 0x65, 0x79, 0x32, 0x56, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x73,
	0x74, 0x76, 0x69, 0x70, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_gateway_gateway_proto_rawDescOnce sync.Once
	file_gateway_gateway_proto_rawDescData []byte
)

func file_gateway_gateway_proto_rawDescGZIP() []byte {
	file_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gateway_gateway_proto_rawDesc), len(file_gateway_gateway_proto_rawDesc)))
	})
	return file_gateway_gateway_proto_rawDescData
}

var file_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_gateway_proto_goTypes = []any{
	(*GateWayInfoResponse)(nil), // 0: gateway.GateWayInfoResponse
	(*emptypb.Empty)(nil),       // 1: google.protobuf.Empty
}
var file_gateway_gateway_proto_depIdxs = []int32{
	1, // 0: gateway.RpcGateway.GetGatewayInfo:input_type -> google.protobuf.Empty
	0, // 1: gateway.RpcGateway.GetGatewayInfo:output_type -> gateway.GateWayInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_gateway_proto_init() }
func file_gateway_gateway_proto_init() {
	if File_gateway_gateway_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gateway_gateway_proto_rawDesc), len(file_gateway_gateway_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_proto_msgTypes,
	}.Build()
	File_gateway_gateway_proto = out.File
	file_gateway_gateway_proto_goTypes = nil
	file_gateway_gateway_proto_depIdxs = nil
}
