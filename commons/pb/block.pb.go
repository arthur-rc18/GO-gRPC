// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: commons/proto/block.proto

package block

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestID) Reset() {
	*x = RequestID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_proto_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestID) ProtoMessage() {}

func (x *RequestID) ProtoReflect() protoreflect.Message {
	mi := &file_commons_proto_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestID.ProtoReflect.Descriptor instead.
func (*RequestID) Descriptor() ([]byte, []int) {
	return file_commons_proto_block_proto_rawDescGZIP(), []int{0}
}

func (x *RequestID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockJson string `protobuf:"bytes,1,opt,name=blockJson,proto3" json:"blockJson,omitempty"`
}

func (x *ResponseBlock) Reset() {
	*x = ResponseBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_proto_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseBlock) ProtoMessage() {}

func (x *ResponseBlock) ProtoReflect() protoreflect.Message {
	mi := &file_commons_proto_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseBlock.ProtoReflect.Descriptor instead.
func (*ResponseBlock) Descriptor() ([]byte, []int) {
	return file_commons_proto_block_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseBlock) GetBlockJson() string {
	if x != nil {
		return x.BlockJson
	}
	return ""
}

var File_commons_proto_block_proto protoreflect.FileDescriptor

var file_commons_proto_block_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4a, 0x73, 0x6f, 0x6e, 0x32, 0x36, 0x0a, 0x06, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x12, 0x2c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0a, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x1a, 0x0e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x62, 0x3b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_proto_block_proto_rawDescOnce sync.Once
	file_commons_proto_block_proto_rawDescData = file_commons_proto_block_proto_rawDesc
)

func file_commons_proto_block_proto_rawDescGZIP() []byte {
	file_commons_proto_block_proto_rawDescOnce.Do(func() {
		file_commons_proto_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_proto_block_proto_rawDescData)
	})
	return file_commons_proto_block_proto_rawDescData
}

var file_commons_proto_block_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commons_proto_block_proto_goTypes = []interface{}{
	(*RequestID)(nil),     // 0: RequestID
	(*ResponseBlock)(nil), // 1: ResponseBlock
}
var file_commons_proto_block_proto_depIdxs = []int32{
	0, // 0: Blocks.GetBlockById:input_type -> RequestID
	1, // 1: Blocks.GetBlockById:output_type -> ResponseBlock
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commons_proto_block_proto_init() }
func file_commons_proto_block_proto_init() {
	if File_commons_proto_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commons_proto_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestID); i {
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
		file_commons_proto_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseBlock); i {
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
			RawDescriptor: file_commons_proto_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commons_proto_block_proto_goTypes,
		DependencyIndexes: file_commons_proto_block_proto_depIdxs,
		MessageInfos:      file_commons_proto_block_proto_msgTypes,
	}.Build()
	File_commons_proto_block_proto = out.File
	file_commons_proto_block_proto_rawDesc = nil
	file_commons_proto_block_proto_goTypes = nil
	file_commons_proto_block_proto_depIdxs = nil
}