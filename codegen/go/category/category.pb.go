// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: category/category.proto

package category

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_category_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRequest.ProtoReflect.Descriptor instead.
func (*FindRequest) Descriptor() ([]byte, []int) {
	return file_category_category_proto_rawDescGZIP(), []int{0}
}

func (x *FindRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type FindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	IsFound    bool   `protobuf:"varint,2,opt,name=is_found,json=isFound,proto3" json:"is_found,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_category_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_category_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindResponse.ProtoReflect.Descriptor instead.
func (*FindResponse) Descriptor() ([]byte, []int) {
	return file_category_category_proto_rawDescGZIP(), []int{1}
}

func (x *FindResponse) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *FindResponse) GetIsFound() bool {
	if x != nil {
		return x.IsFound
	}
	return false
}

var File_category_category_proto protoreflect.FileDescriptor

var file_category_category_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x22, 0x4a, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x32, 0x41, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x35,
	0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71, 0x2f, 0x73, 0x63, 0x72, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_category_category_proto_rawDescOnce sync.Once
	file_category_category_proto_rawDescData = file_category_category_proto_rawDesc
)

func file_category_category_proto_rawDescGZIP() []byte {
	file_category_category_proto_rawDescOnce.Do(func() {
		file_category_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_category_category_proto_rawDescData)
	})
	return file_category_category_proto_rawDescData
}

var file_category_category_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_category_category_proto_goTypes = []interface{}{
	(*FindRequest)(nil),  // 0: category.FindRequest
	(*FindResponse)(nil), // 1: category.FindResponse
}
var file_category_category_proto_depIdxs = []int32{
	0, // 0: category.Category.Find:input_type -> category.FindRequest
	1, // 1: category.Category.Find:output_type -> category.FindResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_category_category_proto_init() }
func file_category_category_proto_init() {
	if File_category_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_category_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRequest); i {
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
		file_category_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindResponse); i {
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
			RawDescriptor: file_category_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_category_category_proto_goTypes,
		DependencyIndexes: file_category_category_proto_depIdxs,
		MessageInfos:      file_category_category_proto_msgTypes,
	}.Build()
	File_category_category_proto = out.File
	file_category_category_proto_rawDesc = nil
	file_category_category_proto_goTypes = nil
	file_category_category_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CategoryClient is the client API for Category service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/category.Category/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
type CategoryServer interface {
	Find(context.Context, *FindRequest) (*FindResponse, error)
}

// UnimplementedCategoryServer can be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (*UnimplementedCategoryServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterCategoryServer(s *grpc.Server, srv CategoryServer) {
	s.RegisterService(&_Category_serviceDesc, srv)
}

func _Category_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/category.Category/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Category_serviceDesc = grpc.ServiceDesc{
	ServiceName: "category.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _Category_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category/category.proto",
}