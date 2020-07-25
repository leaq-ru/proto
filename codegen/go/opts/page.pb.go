// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: opts/page.proto

package opts

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Opts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip  uint32 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Opts) Reset() {
	*x = Opts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_opts_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opts) ProtoMessage() {}

func (x *Opts) ProtoReflect() protoreflect.Message {
	mi := &file_opts_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opts.ProtoReflect.Descriptor instead.
func (*Opts) Descriptor() ([]byte, []int) {
	return file_opts_page_proto_rawDescGZIP(), []int{0}
}

func (x *Opts) GetSkip() uint32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *Opts) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_opts_page_proto protoreflect.FileDescriptor

var file_opts_page_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x04, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0x92, 0x41, 0x03,
	0x3a, 0x01, 0x30, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x12, 0x26, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x10, 0x92, 0x41, 0x0d, 0x3a, 0x02, 0x32,
	0x30, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x59, 0x40, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x6e, 0x71, 0x71, 0x2f, 0x73, 0x63, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_opts_page_proto_rawDescOnce sync.Once
	file_opts_page_proto_rawDescData = file_opts_page_proto_rawDesc
)

func file_opts_page_proto_rawDescGZIP() []byte {
	file_opts_page_proto_rawDescOnce.Do(func() {
		file_opts_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_opts_page_proto_rawDescData)
	})
	return file_opts_page_proto_rawDescData
}

var file_opts_page_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_opts_page_proto_goTypes = []interface{}{
	(*Opts)(nil), // 0: opts.Opts
}
var file_opts_page_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_opts_page_proto_init() }
func file_opts_page_proto_init() {
	if File_opts_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_opts_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opts); i {
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
			RawDescriptor: file_opts_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_opts_page_proto_goTypes,
		DependencyIndexes: file_opts_page_proto_depIdxs,
		MessageInfos:      file_opts_page_proto_msgTypes,
	}.Build()
	File_opts_page_proto = out.File
	file_opts_page_proto_rawDesc = nil
	file_opts_page_proto_goTypes = nil
	file_opts_page_proto_depIdxs = nil
}
