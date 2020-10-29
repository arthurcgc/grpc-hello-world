// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: protos/hello/hello.proto

package hello

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=Hello,proto3" json:"Hello,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_protos_hello_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloResponse) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type EmptyParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyParam) Reset() {
	*x = EmptyParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyParam) ProtoMessage() {}

func (x *EmptyParam) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyParam.ProtoReflect.Descriptor instead.
func (*EmptyParam) Descriptor() ([]byte, []int) {
	return file_protos_hello_hello_proto_rawDescGZIP(), []int{1}
}

var File_protos_hello_hello_proto protoreflect.FileDescriptor

var file_protos_hello_hello_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x32,
	0x30, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x27, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x74, 0x68, 0x75, 0x72, 0x63, 0x67, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_hello_hello_proto_rawDescOnce sync.Once
	file_protos_hello_hello_proto_rawDescData = file_protos_hello_hello_proto_rawDesc
)

func file_protos_hello_hello_proto_rawDescGZIP() []byte {
	file_protos_hello_hello_proto_rawDescOnce.Do(func() {
		file_protos_hello_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_hello_hello_proto_rawDescData)
	})
	return file_protos_hello_hello_proto_rawDescData
}

var file_protos_hello_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_hello_hello_proto_goTypes = []interface{}{
	(*HelloResponse)(nil), // 0: HelloResponse
	(*EmptyParam)(nil),    // 1: EmptyParam
}
var file_protos_hello_hello_proto_depIdxs = []int32{
	1, // 0: Hello.SayHello:input_type -> EmptyParam
	0, // 1: Hello.SayHello:output_type -> HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_hello_hello_proto_init() }
func file_protos_hello_hello_proto_init() {
	if File_protos_hello_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_hello_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_protos_hello_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyParam); i {
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
			RawDescriptor: file_protos_hello_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_hello_hello_proto_goTypes,
		DependencyIndexes: file_protos_hello_hello_proto_depIdxs,
		MessageInfos:      file_protos_hello_hello_proto_msgTypes,
	}.Build()
	File_protos_hello_hello_proto = out.File
	file_protos_hello_hello_proto_rawDesc = nil
	file_protos_hello_hello_proto_goTypes = nil
	file_protos_hello_hello_proto_depIdxs = nil
}
