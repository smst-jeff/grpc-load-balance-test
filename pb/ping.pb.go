// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: pb/ping.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_pb_ping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_ping_proto_msgTypes[0]
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
	return file_pb_ping_proto_rawDescGZIP(), []int{0}
}

var File_pb_ping_proto protoreflect.FileDescriptor

const file_pb_ping_proto_rawDesc = "" +
	"\n" +
	"\rpb/ping.proto\x12\x02pb\"\a\n" +
	"\x05Empty2(\n" +
	"\x06Pinger\x12\x1e\n" +
	"\x04Ping\x12\t.pb.Empty\x1a\t.pb.Empty\"\x00B\x1bZ\x19grpc-load-balance-test/pbb\x06proto3"

var (
	file_pb_ping_proto_rawDescOnce sync.Once
	file_pb_ping_proto_rawDescData []byte
)

func file_pb_ping_proto_rawDescGZIP() []byte {
	file_pb_ping_proto_rawDescOnce.Do(func() {
		file_pb_ping_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pb_ping_proto_rawDesc), len(file_pb_ping_proto_rawDesc)))
	})
	return file_pb_ping_proto_rawDescData
}

var file_pb_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_ping_proto_goTypes = []any{
	(*Empty)(nil), // 0: pb.Empty
}
var file_pb_ping_proto_depIdxs = []int32{
	0, // 0: pb.Pinger.Ping:input_type -> pb.Empty
	0, // 1: pb.Pinger.Ping:output_type -> pb.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_ping_proto_init() }
func file_pb_ping_proto_init() {
	if File_pb_ping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pb_ping_proto_rawDesc), len(file_pb_ping_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_ping_proto_goTypes,
		DependencyIndexes: file_pb_ping_proto_depIdxs,
		MessageInfos:      file_pb_ping_proto_msgTypes,
	}.Build()
	File_pb_ping_proto = out.File
	file_pb_ping_proto_goTypes = nil
	file_pb_ping_proto_depIdxs = nil
}
