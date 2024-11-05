// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: consensus.proto

package consensus

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

type SuccessStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuccessStart) Reset() {
	*x = SuccessStart{}
	mi := &file_consensus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessStart) ProtoMessage() {}

func (x *SuccessStart) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessStart.ProtoReflect.Descriptor instead.
func (*SuccessStart) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *SuccessStart) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_consensus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_consensus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{2}
}

var File_consensus_proto protoreflect.FileDescriptor

var file_consensus_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x50, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x06, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x15, 0x5a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_proto_rawDescOnce sync.Once
	file_consensus_proto_rawDescData = file_consensus_proto_rawDesc
)

func file_consensus_proto_rawDescGZIP() []byte {
	file_consensus_proto_rawDescOnce.Do(func() {
		file_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_proto_rawDescData)
	})
	return file_consensus_proto_rawDescData
}

var file_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_consensus_proto_goTypes = []any{
	(*SuccessStart)(nil), // 0: SuccessStart
	(*Token)(nil),        // 1: Token
	(*Empty)(nil),        // 2: Empty
}
var file_consensus_proto_depIdxs = []int32{
	2, // 0: consensus.StartFunction:input_type -> Empty
	1, // 1: consensus.PassToken:input_type -> Token
	0, // 2: consensus.StartFunction:output_type -> SuccessStart
	2, // 3: consensus.PassToken:output_type -> Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_consensus_proto_init() }
func file_consensus_proto_init() {
	if File_consensus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consensus_proto_goTypes,
		DependencyIndexes: file_consensus_proto_depIdxs,
		MessageInfos:      file_consensus_proto_msgTypes,
	}.Build()
	File_consensus_proto = out.File
	file_consensus_proto_rawDesc = nil
	file_consensus_proto_goTypes = nil
	file_consensus_proto_depIdxs = nil
}
