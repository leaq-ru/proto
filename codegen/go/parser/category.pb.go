// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: parser/category.proto

package parser

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type GetCategoryHintsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetCategoryHintsRequest) Reset() {
	*x = GetCategoryHintsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryHintsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryHintsRequest) ProtoMessage() {}

func (x *GetCategoryHintsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryHintsRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryHintsRequest) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{0}
}

func (x *GetCategoryHintsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetCategoryHintsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCategoryBySlugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *GetCategoryBySlugRequest) Reset() {
	*x = GetCategoryBySlugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryBySlugRequest) ProtoMessage() {}

func (x *GetCategoryBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryBySlugRequest) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoryBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetCategoryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *GetCategoryByIdRequest) Reset() {
	*x = GetCategoryByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdRequest) ProtoMessage() {}

func (x *GetCategoryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdRequest) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryByIdRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type GetCategoryByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryIds []string `protobuf:"bytes,1,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
}

func (x *GetCategoryByIdsRequest) Reset() {
	*x = GetCategoryByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryByIdsRequest) ProtoMessage() {}

func (x *GetCategoryByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryByIdsRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryByIdsRequest) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{3}
}

func (x *GetCategoryByIdsRequest) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

type CategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*CategoryItem `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoriesResponse) Reset() {
	*x = CategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesResponse) ProtoMessage() {}

func (x *CategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesResponse.ProtoReflect.Descriptor instead.
func (*CategoriesResponse) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoriesResponse) GetCategories() []*CategoryItem {
	if x != nil {
		return x.Categories
	}
	return nil
}

type FindCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *FindCategoryRequest) Reset() {
	*x = FindCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCategoryRequest) ProtoMessage() {}

func (x *FindCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCategoryRequest.ProtoReflect.Descriptor instead.
func (*FindCategoryRequest) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{5}
}

func (x *FindCategoryRequest) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

type FindCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *FindCategoryResponse) Reset() {
	*x = FindCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCategoryResponse) ProtoMessage() {}

func (x *FindCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCategoryResponse.ProtoReflect.Descriptor instead.
func (*FindCategoryResponse) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{6}
}

func (x *FindCategoryResponse) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type CategoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug  string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *CategoryItem) Reset() {
	*x = CategoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryItem) ProtoMessage() {}

func (x *CategoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_parser_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryItem.ProtoReflect.Descriptor instead.
func (*CategoryItem) Descriptor() ([]byte, []int) {
	return file_parser_category_proto_rawDescGZIP(), []int{7}
}

func (x *CategoryItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CategoryItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CategoryItem) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_parser_category_proto protoreflect.FileDescriptor

var file_parser_category_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x21, 0x92, 0x41, 0x1e, 0x32, 0x0a, 0x32,
	0x30, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x59, 0x40, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x3c, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74,
	0x6d, 0x6c, 0x22, 0x37, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x32, 0xec, 0x04, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x49, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79,
	0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65,
	0x6d, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75,
	0x67, 0x12, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x12, 0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x48, 0x69, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x48,
	0x69, 0x6e, 0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x61, 0x71, 0x2d, 0x72, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parser_category_proto_rawDescOnce sync.Once
	file_parser_category_proto_rawDescData = file_parser_category_proto_rawDesc
)

func file_parser_category_proto_rawDescGZIP() []byte {
	file_parser_category_proto_rawDescOnce.Do(func() {
		file_parser_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_parser_category_proto_rawDescData)
	})
	return file_parser_category_proto_rawDescData
}

var file_parser_category_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_parser_category_proto_goTypes = []interface{}{
	(*GetCategoryHintsRequest)(nil),  // 0: parser.GetCategoryHintsRequest
	(*GetCategoryBySlugRequest)(nil), // 1: parser.GetCategoryBySlugRequest
	(*GetCategoryByIdRequest)(nil),   // 2: parser.GetCategoryByIdRequest
	(*GetCategoryByIdsRequest)(nil),  // 3: parser.GetCategoryByIdsRequest
	(*CategoriesResponse)(nil),       // 4: parser.CategoriesResponse
	(*FindCategoryRequest)(nil),      // 5: parser.FindCategoryRequest
	(*FindCategoryResponse)(nil),     // 6: parser.FindCategoryResponse
	(*CategoryItem)(nil),             // 7: parser.CategoryItem
	(*emptypb.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_parser_category_proto_depIdxs = []int32{
	7, // 0: parser.CategoriesResponse.categories:type_name -> parser.CategoryItem
	5, // 1: parser.Category.FindCategory:input_type -> parser.FindCategoryRequest
	2, // 2: parser.Category.GetCategoryById:input_type -> parser.GetCategoryByIdRequest
	1, // 3: parser.Category.GetCategoryBySlug:input_type -> parser.GetCategoryBySlugRequest
	3, // 4: parser.Category.GetCategoryByIds:input_type -> parser.GetCategoryByIdsRequest
	8, // 5: parser.Category.GetAllCategory:input_type -> google.protobuf.Empty
	0, // 6: parser.Category.GetCategoryHints:input_type -> parser.GetCategoryHintsRequest
	6, // 7: parser.Category.FindCategory:output_type -> parser.FindCategoryResponse
	7, // 8: parser.Category.GetCategoryById:output_type -> parser.CategoryItem
	7, // 9: parser.Category.GetCategoryBySlug:output_type -> parser.CategoryItem
	4, // 10: parser.Category.GetCategoryByIds:output_type -> parser.CategoriesResponse
	4, // 11: parser.Category.GetAllCategory:output_type -> parser.CategoriesResponse
	4, // 12: parser.Category.GetCategoryHints:output_type -> parser.CategoriesResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parser_category_proto_init() }
func file_parser_category_proto_init() {
	if File_parser_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parser_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryHintsRequest); i {
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
		file_parser_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryBySlugRequest); i {
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
		file_parser_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdRequest); i {
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
		file_parser_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryByIdsRequest); i {
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
		file_parser_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesResponse); i {
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
		file_parser_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCategoryRequest); i {
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
		file_parser_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCategoryResponse); i {
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
		file_parser_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryItem); i {
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
			RawDescriptor: file_parser_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parser_category_proto_goTypes,
		DependencyIndexes: file_parser_category_proto_depIdxs,
		MessageInfos:      file_parser_category_proto_msgTypes,
	}.Build()
	File_parser_category_proto = out.File
	file_parser_category_proto_rawDesc = nil
	file_parser_category_proto_goTypes = nil
	file_parser_category_proto_depIdxs = nil
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
	FindCategory(ctx context.Context, in *FindCategoryRequest, opts ...grpc.CallOption) (*FindCategoryResponse, error)
	GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*CategoryItem, error)
	GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*CategoryItem, error)
	GetCategoryByIds(ctx context.Context, in *GetCategoryByIdsRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
	GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResponse, error)
	GetCategoryHints(ctx context.Context, in *GetCategoryHintsRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
}

type categoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryClient(cc grpc.ClientConnInterface) CategoryClient {
	return &categoryClient{cc}
}

func (c *categoryClient) FindCategory(ctx context.Context, in *FindCategoryRequest, opts ...grpc.CallOption) (*FindCategoryResponse, error) {
	out := new(FindCategoryResponse)
	err := c.cc.Invoke(ctx, "/parser.Category/FindCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetCategoryById(ctx context.Context, in *GetCategoryByIdRequest, opts ...grpc.CallOption) (*CategoryItem, error) {
	out := new(CategoryItem)
	err := c.cc.Invoke(ctx, "/parser.Category/GetCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetCategoryBySlug(ctx context.Context, in *GetCategoryBySlugRequest, opts ...grpc.CallOption) (*CategoryItem, error) {
	out := new(CategoryItem)
	err := c.cc.Invoke(ctx, "/parser.Category/GetCategoryBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetCategoryByIds(ctx context.Context, in *GetCategoryByIdsRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/parser.Category/GetCategoryByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetAllCategory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/parser.Category/GetAllCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryClient) GetCategoryHints(ctx context.Context, in *GetCategoryHintsRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/parser.Category/GetCategoryHints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServer is the server API for Category service.
type CategoryServer interface {
	FindCategory(context.Context, *FindCategoryRequest) (*FindCategoryResponse, error)
	GetCategoryById(context.Context, *GetCategoryByIdRequest) (*CategoryItem, error)
	GetCategoryBySlug(context.Context, *GetCategoryBySlugRequest) (*CategoryItem, error)
	GetCategoryByIds(context.Context, *GetCategoryByIdsRequest) (*CategoriesResponse, error)
	GetAllCategory(context.Context, *emptypb.Empty) (*CategoriesResponse, error)
	GetCategoryHints(context.Context, *GetCategoryHintsRequest) (*CategoriesResponse, error)
}

// UnimplementedCategoryServer can be embedded to have forward compatible implementations.
type UnimplementedCategoryServer struct {
}

func (*UnimplementedCategoryServer) FindCategory(context.Context, *FindCategoryRequest) (*FindCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCategory not implemented")
}
func (*UnimplementedCategoryServer) GetCategoryById(context.Context, *GetCategoryByIdRequest) (*CategoryItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
func (*UnimplementedCategoryServer) GetCategoryBySlug(context.Context, *GetCategoryBySlugRequest) (*CategoryItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryBySlug not implemented")
}
func (*UnimplementedCategoryServer) GetCategoryByIds(context.Context, *GetCategoryByIdsRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryByIds not implemented")
}
func (*UnimplementedCategoryServer) GetAllCategory(context.Context, *emptypb.Empty) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategory not implemented")
}
func (*UnimplementedCategoryServer) GetCategoryHints(context.Context, *GetCategoryHintsRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryHints not implemented")
}

func RegisterCategoryServer(s *grpc.Server, srv CategoryServer) {
	s.RegisterService(&_Category_serviceDesc, srv)
}

func _Category_FindCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).FindCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/FindCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).FindCategory(ctx, req.(*FindCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/GetCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryById(ctx, req.(*GetCategoryByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetCategoryBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryBySlugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/GetCategoryBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryBySlug(ctx, req.(*GetCategoryBySlugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetCategoryByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/GetCategoryByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryByIds(ctx, req.(*GetCategoryByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetAllCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetAllCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/GetAllCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetAllCategory(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Category_GetCategoryHints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryHintsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServer).GetCategoryHints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Category/GetCategoryHints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServer).GetCategoryHints(ctx, req.(*GetCategoryHintsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Category_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parser.Category",
	HandlerType: (*CategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCategory",
			Handler:    _Category_FindCategory_Handler,
		},
		{
			MethodName: "GetCategoryById",
			Handler:    _Category_GetCategoryById_Handler,
		},
		{
			MethodName: "GetCategoryBySlug",
			Handler:    _Category_GetCategoryBySlug_Handler,
		},
		{
			MethodName: "GetCategoryByIds",
			Handler:    _Category_GetCategoryByIds_Handler,
		},
		{
			MethodName: "GetAllCategory",
			Handler:    _Category_GetAllCategory_Handler,
		},
		{
			MethodName: "GetCategoryHints",
			Handler:    _Category_GetCategoryHints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser/category.proto",
}
