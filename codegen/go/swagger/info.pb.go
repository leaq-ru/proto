// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: swagger/info.proto

package swagger

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_swagger_info_proto protoreflect.FileDescriptor

var file_swagger_info_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xb1, 0x04, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71, 0x2f,
	0x73, 0x63, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x92, 0x41, 0xff, 0x03,
	0x12, 0x83, 0x03, 0x0a, 0x08, 0x4c, 0x45, 0x41, 0x51, 0x20, 0x41, 0x50, 0x49, 0x12, 0xc2, 0x02,
	0xd0, 0x94, 0xd0, 0xbe, 0xd1, 0x81, 0xd1, 0x82, 0xd1, 0x83, 0xd0, 0xbf, 0x20, 0xd0, 0xba, 0x20,
	0x41, 0x50, 0x49, 0x20, 0xd0, 0xb1, 0xd0, 0xb5, 0xd0, 0xb7, 0x20, 0xd0, 0xbe, 0xd0, 0xb3, 0xd1,
	0x80, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0,
	0xb9, 0x20, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xba, 0xd0, 0xbe, 0xd0, 0xbb, 0x2d, 0xd0, 0xb2,
	0xd0, 0xbe, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x81, 0xd0,
	0xbe, 0xd0, 0xb2, 0x20, 0xd0, 0xb2, 0x20, 0xd0, 0xb4, 0xd0, 0xb5, 0xd0, 0xbd, 0xd1, 0x8c, 0x2f,
	0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x81, 0xd1, 0x8f, 0xd1, 0x86, 0x2e, 0x20, 0xd0, 0x9e, 0xd0, 0xb3,
	0xd1, 0x80, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8,
	0xd0, 0xb5, 0x20, 0x31, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1,
	0x81, 0x20, 0xd0, 0xb2, 0x20, 0xd1, 0x81, 0xd0, 0xb5, 0xd0, 0xba, 0xd1, 0x83, 0xd0, 0xbd, 0xd0,
	0xb4, 0xd1, 0x83, 0x20, 0xd1, 0x81, 0x20, 0xd0, 0xbe, 0xd0, 0xb4, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0,
	0xb3, 0xd0, 0xbe, 0x20, 0x49, 0x50, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xb1,
	0xd0, 0xb5, 0xd1, 0x81, 0xd0, 0xbf, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbd, 0xd0, 0xbe,
	0xd0, 0xb3, 0xd0, 0xbe, 0x20, 0xd1, 0x82, 0xd0, 0xb0, 0xd1, 0x80, 0xd0, 0xb8, 0xd1, 0x84, 0xd0,
	0xb0, 0x2c, 0x20, 0xd0, 0xb8, 0x20, 0x33, 0x30, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20,
	0xd0, 0xbf, 0xd0, 0xbb, 0xd0, 0xb0, 0xd1, 0x82, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe,
	0x2e, 0x20, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xb4, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb1, 0xd0, 0xbd,
	0xd0, 0xb5, 0xd0, 0xb5, 0x20, 0xd0, 0xbe, 0x20, 0xd1, 0x82, 0xd0, 0xb0, 0xd1, 0x80, 0xd0, 0xb8,
	0xd1, 0x84, 0xd0, 0xb0, 0xd1, 0x85, 0x3a, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x6c, 0x65, 0x61, 0x71, 0x2e, 0x72, 0x75, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x23, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x2b, 0x0a, 0x12, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xb4, 0xd0, 0xb4, 0xd0, 0xb5,
	0xd1, 0x80, 0xd0, 0xb6, 0xd0, 0xba, 0xd0, 0xb0, 0x12, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x76, 0x6b, 0x2e, 0x6d, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x71, 0x5f, 0x72, 0x75, 0x32,
	0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x42, 0x0a,
	0x40, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x08, 0x02, 0x12, 0x21, 0x50,
	0x61, 0x73, 0x74, 0x65, 0x20, 0x27, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x58, 0x58, 0x58,
	0x58, 0x58, 0x27, 0x2c, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x27, 0x58, 0x58, 0x58, 0x58, 0x58, 0x27,
	0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_swagger_info_proto_goTypes = []interface{}{}
var file_swagger_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_swagger_info_proto_init() }
func file_swagger_info_proto_init() {
	if File_swagger_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_swagger_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_swagger_info_proto_goTypes,
		DependencyIndexes: file_swagger_info_proto_depIdxs,
	}.Build()
	File_swagger_info_proto = out.File
	file_swagger_info_proto_rawDesc = nil
	file_swagger_info_proto_goTypes = nil
	file_swagger_info_proto_depIdxs = nil
}
