// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: parser/post.proto

package parser

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	opts "github.com/leaq-ru/proto/codegen/go/opts"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opts      *opts.Page `protobuf:"bytes,1,opt,name=opts,proto3" json:"opts,omitempty"`
	CompanyId string     `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_parser_post_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostsRequest) GetOpts() *opts.Page {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *GetPostsRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type PhotoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UrlM string `protobuf:"bytes,1,opt,name=url_m,json=urlM,proto3" json:"url_m,omitempty"`
	UrlR string `protobuf:"bytes,2,opt,name=url_r,json=urlR,proto3" json:"url_r,omitempty"`
}

func (x *PhotoItem) Reset() {
	*x = PhotoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoItem) ProtoMessage() {}

func (x *PhotoItem) ProtoReflect() protoreflect.Message {
	mi := &file_parser_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoItem.ProtoReflect.Descriptor instead.
func (*PhotoItem) Descriptor() ([]byte, []int) {
	return file_parser_post_proto_rawDescGZIP(), []int{1}
}

func (x *PhotoItem) GetUrlM() string {
	if x != nil {
		return x.UrlM
	}
	return ""
}

func (x *PhotoItem) GetUrlR() string {
	if x != nil {
		return x.UrlR
	}
	return ""
}

type PostItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId string       `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Date      string       `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Text      string       `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Photos    []*PhotoItem `protobuf:"bytes,5,rep,name=photos,proto3" json:"photos,omitempty"`
}

func (x *PostItem) Reset() {
	*x = PostItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostItem) ProtoMessage() {}

func (x *PostItem) ProtoReflect() protoreflect.Message {
	mi := &file_parser_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostItem.ProtoReflect.Descriptor instead.
func (*PostItem) Descriptor() ([]byte, []int) {
	return file_parser_post_proto_rawDescGZIP(), []int{2}
}

func (x *PostItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostItem) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *PostItem) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PostItem) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PostItem) GetPhotos() []*PhotoItem {
	if x != nil {
		return x.Photos
	}
	return nil
}

type GetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*PostItem `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parser_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_parser_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostsResponse) GetPosts() []*PostItem {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_parser_post_proto protoreflect.FileDescriptor

var file_parser_post_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6f, 0x70, 0x74, 0x73, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6f, 0x70,
	0x74, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x09,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x13, 0x0a, 0x05, 0x75, 0x72, 0x6c,
	0x5f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x4d, 0x12, 0x13,
	0x0a, 0x05, 0x75, 0x72, 0x6c, 0x5f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x72, 0x6c, 0x52, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x73, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0x5b,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x61, 0x71, 0x2d, 0x72,
	0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_parser_post_proto_rawDescOnce sync.Once
	file_parser_post_proto_rawDescData = file_parser_post_proto_rawDesc
)

func file_parser_post_proto_rawDescGZIP() []byte {
	file_parser_post_proto_rawDescOnce.Do(func() {
		file_parser_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_parser_post_proto_rawDescData)
	})
	return file_parser_post_proto_rawDescData
}

var file_parser_post_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_parser_post_proto_goTypes = []interface{}{
	(*GetPostsRequest)(nil),  // 0: parser.GetPostsRequest
	(*PhotoItem)(nil),        // 1: parser.PhotoItem
	(*PostItem)(nil),         // 2: parser.PostItem
	(*GetPostsResponse)(nil), // 3: parser.GetPostsResponse
	(*opts.Page)(nil),        // 4: opts.Page
}
var file_parser_post_proto_depIdxs = []int32{
	4, // 0: parser.GetPostsRequest.opts:type_name -> opts.Page
	1, // 1: parser.PostItem.photos:type_name -> parser.PhotoItem
	2, // 2: parser.GetPostsResponse.posts:type_name -> parser.PostItem
	0, // 3: parser.Post.GetPosts:input_type -> parser.GetPostsRequest
	3, // 4: parser.Post.GetPosts:output_type -> parser.GetPostsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_parser_post_proto_init() }
func file_parser_post_proto_init() {
	if File_parser_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parser_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
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
		file_parser_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoItem); i {
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
		file_parser_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostItem); i {
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
		file_parser_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsResponse); i {
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
			RawDescriptor: file_parser_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parser_post_proto_goTypes,
		DependencyIndexes: file_parser_post_proto_depIdxs,
		MessageInfos:      file_parser_post_proto_msgTypes,
	}.Build()
	File_parser_post_proto = out.File
	file_parser_post_proto_rawDesc = nil
	file_parser_post_proto_goTypes = nil
	file_parser_post_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PostClient is the client API for Post service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostClient interface {
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
}

type postClient struct {
	cc grpc.ClientConnInterface
}

func NewPostClient(cc grpc.ClientConnInterface) PostClient {
	return &postClient{cc}
}

func (c *postClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, "/parser.Post/GetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServer is the server API for Post service.
type PostServer interface {
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
}

// UnimplementedPostServer can be embedded to have forward compatible implementations.
type UnimplementedPostServer struct {
}

func (*UnimplementedPostServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}

func RegisterPostServer(s *grpc.Server, srv PostServer) {
	s.RegisterService(&_Post_serviceDesc, srv)
}

func _Post_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Post/GetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Post_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parser.Post",
	HandlerType: (*PostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPosts",
			Handler:    _Post_GetPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser/post.proto",
}
