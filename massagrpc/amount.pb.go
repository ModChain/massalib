// Copyright (c) 2023 MASSA LABS <info@massa.net>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.27.2
// source: massa/model/v1/amount.proto

package massagrpc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NativeAmount is represented as a fraction so precision can be adjusted in
// the future. value = mantissa / (10^scale)
type NativeAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mantissa
	Mantissa uint64 `protobuf:"varint,1,opt,name=mantissa,proto3" json:"mantissa,omitempty"`
	// Scale
	Scale uint32 `protobuf:"varint,2,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *NativeAmount) Reset() {
	*x = NativeAmount{}
	mi := &file_massa_model_v1_amount_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NativeAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NativeAmount) ProtoMessage() {}

func (x *NativeAmount) ProtoReflect() protoreflect.Message {
	mi := &file_massa_model_v1_amount_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NativeAmount.ProtoReflect.Descriptor instead.
func (*NativeAmount) Descriptor() ([]byte, []int) {
	return file_massa_model_v1_amount_proto_rawDescGZIP(), []int{0}
}

func (x *NativeAmount) GetMantissa() uint64 {
	if x != nil {
		return x.Mantissa
	}
	return 0
}

func (x *NativeAmount) GetScale() uint32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

var File_massa_model_v1_amount_proto protoreflect.FileDescriptor

var file_massa_model_v1_amount_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d,
	0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a,
	0x0c, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x73, 0x73, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x42,
	0x89, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x61,
	0x73, 0x73, 0x61, 0x67, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x06, 0x4d, 0x4d, 0x4f, 0x44, 0x45, 0x4c,
	0xaa, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x2e, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0xba, 0x02, 0x06, 0x4d, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0xca, 0x02,
	0x12, 0x43, 0x6f, 0x6d, 0x5c, 0x4d, 0x61, 0x73, 0x73, 0x61, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x3a, 0x3a, 0x4d, 0x61, 0x73, 0x73, 0x61,
	0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_massa_model_v1_amount_proto_rawDescOnce sync.Once
	file_massa_model_v1_amount_proto_rawDescData = file_massa_model_v1_amount_proto_rawDesc
)

func file_massa_model_v1_amount_proto_rawDescGZIP() []byte {
	file_massa_model_v1_amount_proto_rawDescOnce.Do(func() {
		file_massa_model_v1_amount_proto_rawDescData = protoimpl.X.CompressGZIP(file_massa_model_v1_amount_proto_rawDescData)
	})
	return file_massa_model_v1_amount_proto_rawDescData
}

var file_massa_model_v1_amount_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_massa_model_v1_amount_proto_goTypes = []any{
	(*NativeAmount)(nil), // 0: massa.model.v1.NativeAmount
}
var file_massa_model_v1_amount_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_massa_model_v1_amount_proto_init() }
func file_massa_model_v1_amount_proto_init() {
	if File_massa_model_v1_amount_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_massa_model_v1_amount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_massa_model_v1_amount_proto_goTypes,
		DependencyIndexes: file_massa_model_v1_amount_proto_depIdxs,
		MessageInfos:      file_massa_model_v1_amount_proto_msgTypes,
	}.Build()
	File_massa_model_v1_amount_proto = out.File
	file_massa_model_v1_amount_proto_rawDesc = nil
	file_massa_model_v1_amount_proto_goTypes = nil
	file_massa_model_v1_amount_proto_depIdxs = nil
}
