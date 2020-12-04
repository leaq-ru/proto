// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: exporter/exporter.proto

package exporter

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	opts "github.com/nnqq/scr-proto/codegen/go/opts"
	parser "github.com/nnqq/scr-proto/codegen/go/parser"
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

type Status int32

const (
	Status_NONE        Status = 0
	Status_QUEUED      Status = 1
	Status_IN_PROGRESS Status = 2
	Status_SUCCESS     Status = 3
	Status_FAILED      Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "NONE",
		1: "QUEUED",
		2: "IN_PROGRESS",
		3: "SUCCESS",
		4: "FAILED",
	}
	Status_value = map[string]int32{
		"NONE":        0,
		"QUEUED":      1,
		"IN_PROGRESS": 2,
		"SUCCESS":     3,
		"FAILED":      4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_exporter_exporter_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_exporter_exporter_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_exporter_exporter_proto_rawDescGZIP(), []int{0}
}

type GetMyFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opts *opts.SkipLimit `protobuf:"bytes,1,opt,name=opts,proto3" json:"opts,omitempty"`
}

func (x *GetMyFilesRequest) Reset() {
	*x = GetMyFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exporter_exporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyFilesRequest) ProtoMessage() {}

func (x *GetMyFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exporter_exporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyFilesRequest.ProtoReflect.Descriptor instead.
func (*GetMyFilesRequest) Descriptor() ([]byte, []int) {
	return file_exporter_exporter_proto_rawDescGZIP(), []int{0}
}

func (x *GetMyFilesRequest) GetOpts() *opts.SkipLimit {
	if x != nil {
		return x.Opts
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exporter_exporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_exporter_exporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_exporter_exporter_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *File) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetMyFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *GetMyFilesResponse) Reset() {
	*x = GetMyFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exporter_exporter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyFilesResponse) ProtoMessage() {}

func (x *GetMyFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exporter_exporter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyFilesResponse.ProtoReflect.Descriptor instead.
func (*GetMyFilesResponse) Descriptor() ([]byte, []int) {
	return file_exporter_exporter_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyFilesResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type ExportCompaniesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ExportCompaniesResponse) Reset() {
	*x = ExportCompaniesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exporter_exporter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportCompaniesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportCompaniesResponse) ProtoMessage() {}

func (x *ExportCompaniesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exporter_exporter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportCompaniesResponse.ProtoReflect.Descriptor instead.
func (*ExportCompaniesResponse) Descriptor() ([]byte, []int) {
	return file_exporter_exporter_proto_rawDescGZIP(), []int{3}
}

func (x *ExportCompaniesResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_exporter_exporter_proto protoreflect.FileDescriptor

var file_exporter_exporter_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6f, 0x70,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e,
	0x53, 0x6b, 0x69, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x22,
	0x5b, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x2a, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45,
	0x55, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32,
	0xdb, 0x02, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x72, 0x0a, 0x0f,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73,
	0x12, 0x71, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x69, 0x65, 0x73, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23,
	0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x41, 0x73,
	0x79, 0x6e, 0x63, 0x12, 0x68, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x4d, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71,
	0x2f, 0x73, 0x63, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exporter_exporter_proto_rawDescOnce sync.Once
	file_exporter_exporter_proto_rawDescData = file_exporter_exporter_proto_rawDesc
)

func file_exporter_exporter_proto_rawDescGZIP() []byte {
	file_exporter_exporter_proto_rawDescOnce.Do(func() {
		file_exporter_exporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_exporter_exporter_proto_rawDescData)
	})
	return file_exporter_exporter_proto_rawDescData
}

var file_exporter_exporter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_exporter_exporter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_exporter_exporter_proto_goTypes = []interface{}{
	(Status)(0),                     // 0: exporter.Status
	(*GetMyFilesRequest)(nil),       // 1: exporter.GetMyFilesRequest
	(*File)(nil),                    // 2: exporter.File
	(*GetMyFilesResponse)(nil),      // 3: exporter.GetMyFilesResponse
	(*ExportCompaniesResponse)(nil), // 4: exporter.ExportCompaniesResponse
	(*opts.SkipLimit)(nil),          // 5: opts.SkipLimit
	(*parser.GetListRequest)(nil),   // 6: parser.GetListRequest
	(*empty.Empty)(nil),             // 7: google.protobuf.Empty
}
var file_exporter_exporter_proto_depIdxs = []int32{
	5, // 0: exporter.GetMyFilesRequest.opts:type_name -> opts.SkipLimit
	2, // 1: exporter.GetMyFilesResponse.files:type_name -> exporter.File
	6, // 2: exporter.Exporter.ExportCompanies:input_type -> parser.GetListRequest
	6, // 3: exporter.Exporter.ExportCompaniesAsync:input_type -> parser.GetListRequest
	1, // 4: exporter.Exporter.GetMyFiles:input_type -> exporter.GetMyFilesRequest
	4, // 5: exporter.Exporter.ExportCompanies:output_type -> exporter.ExportCompaniesResponse
	7, // 6: exporter.Exporter.ExportCompaniesAsync:output_type -> google.protobuf.Empty
	3, // 7: exporter.Exporter.GetMyFiles:output_type -> exporter.GetMyFilesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exporter_exporter_proto_init() }
func file_exporter_exporter_proto_init() {
	if File_exporter_exporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exporter_exporter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyFilesRequest); i {
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
		file_exporter_exporter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_exporter_exporter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyFilesResponse); i {
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
		file_exporter_exporter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportCompaniesResponse); i {
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
			RawDescriptor: file_exporter_exporter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exporter_exporter_proto_goTypes,
		DependencyIndexes: file_exporter_exporter_proto_depIdxs,
		EnumInfos:         file_exporter_exporter_proto_enumTypes,
		MessageInfos:      file_exporter_exporter_proto_msgTypes,
	}.Build()
	File_exporter_exporter_proto = out.File
	file_exporter_exporter_proto_rawDesc = nil
	file_exporter_exporter_proto_goTypes = nil
	file_exporter_exporter_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExporterClient is the client API for Exporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExporterClient interface {
	ExportCompanies(ctx context.Context, in *parser.GetListRequest, opts ...grpc.CallOption) (*ExportCompaniesResponse, error)
	ExportCompaniesAsync(ctx context.Context, in *parser.GetListRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetMyFiles(ctx context.Context, in *GetMyFilesRequest, opts ...grpc.CallOption) (*GetMyFilesResponse, error)
}

type exporterClient struct {
	cc grpc.ClientConnInterface
}

func NewExporterClient(cc grpc.ClientConnInterface) ExporterClient {
	return &exporterClient{cc}
}

func (c *exporterClient) ExportCompanies(ctx context.Context, in *parser.GetListRequest, opts ...grpc.CallOption) (*ExportCompaniesResponse, error) {
	out := new(ExportCompaniesResponse)
	err := c.cc.Invoke(ctx, "/exporter.Exporter/ExportCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exporterClient) ExportCompaniesAsync(ctx context.Context, in *parser.GetListRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/exporter.Exporter/ExportCompaniesAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exporterClient) GetMyFiles(ctx context.Context, in *GetMyFilesRequest, opts ...grpc.CallOption) (*GetMyFilesResponse, error) {
	out := new(GetMyFilesResponse)
	err := c.cc.Invoke(ctx, "/exporter.Exporter/GetMyFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExporterServer is the server API for Exporter service.
type ExporterServer interface {
	ExportCompanies(context.Context, *parser.GetListRequest) (*ExportCompaniesResponse, error)
	ExportCompaniesAsync(context.Context, *parser.GetListRequest) (*empty.Empty, error)
	GetMyFiles(context.Context, *GetMyFilesRequest) (*GetMyFilesResponse, error)
}

// UnimplementedExporterServer can be embedded to have forward compatible implementations.
type UnimplementedExporterServer struct {
}

func (*UnimplementedExporterServer) ExportCompanies(context.Context, *parser.GetListRequest) (*ExportCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportCompanies not implemented")
}
func (*UnimplementedExporterServer) ExportCompaniesAsync(context.Context, *parser.GetListRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportCompaniesAsync not implemented")
}
func (*UnimplementedExporterServer) GetMyFiles(context.Context, *GetMyFilesRequest) (*GetMyFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyFiles not implemented")
}

func RegisterExporterServer(s *grpc.Server, srv ExporterServer) {
	s.RegisterService(&_Exporter_serviceDesc, srv)
}

func _Exporter_ExportCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(parser.GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServer).ExportCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exporter.Exporter/ExportCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServer).ExportCompanies(ctx, req.(*parser.GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exporter_ExportCompaniesAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(parser.GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServer).ExportCompaniesAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exporter.Exporter/ExportCompaniesAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServer).ExportCompaniesAsync(ctx, req.(*parser.GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exporter_GetMyFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExporterServer).GetMyFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exporter.Exporter/GetMyFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExporterServer).GetMyFiles(ctx, req.(*GetMyFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exporter.Exporter",
	HandlerType: (*ExporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportCompanies",
			Handler:    _Exporter_ExportCompanies_Handler,
		},
		{
			MethodName: "ExportCompaniesAsync",
			Handler:    _Exporter_ExportCompaniesAsync_Handler,
		},
		{
			MethodName: "GetMyFiles",
			Handler:    _Exporter_GetMyFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exporter/exporter.proto",
}
