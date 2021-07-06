// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: parser/review.proto

package parser

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	opts "github.com/nnqq/scr-proto/codegen/go/opts"
	user "github.com/nnqq/scr-proto/codegen/go/user"
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

type ChangeStatusRequestStatus int32

const (
	ChangeStatusRequest_NONE   ChangeStatusRequestStatus = 0
	ChangeStatusRequest_OK     ChangeStatusRequestStatus = 1
	ChangeStatusRequest_DELETE ChangeStatusRequestStatus = 2
)

// Enum value maps for ChangeStatusRequestStatus.
var (
	ChangeStatusRequestStatus_name = map[int32]string{
		0: "NONE",
		1: "OK",
		2: "DELETE",
	}
	ChangeStatusRequestStatus_value = map[string]int32{
		"NONE":   0,
		"OK":     1,
		"DELETE": 2,
	}
)

func (x ChangeStatusRequestStatus) Enum() *ChangeStatusRequestStatus {
	p := new(ChangeStatusRequestStatus)
	*p = x
	return p
}

func (x ChangeStatusRequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeStatusRequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_parser_review_proto_enumTypes[0].Descriptor()
}

func (ChangeStatusRequestStatus) Type() protoreflect.EnumType {
	return &file_parser_review_proto_enumTypes[0]
}

func (x ChangeStatusRequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeStatusRequestStatus.Descriptor instead.
func (ChangeStatusRequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{2, 0}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId string `protobuf:"bytes,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetReviewId() string {
	if x != nil {
		return x.ReviewId
	}
	return ""
}

type DisallowAuthorDeleteAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DisallowAuthorDeleteAllRequest) Reset() {
	*x = DisallowAuthorDeleteAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisallowAuthorDeleteAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisallowAuthorDeleteAllRequest) ProtoMessage() {}

func (x *DisallowAuthorDeleteAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisallowAuthorDeleteAllRequest.ProtoReflect.Descriptor instead.
func (*DisallowAuthorDeleteAllRequest) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{1}
}

func (x *DisallowAuthorDeleteAllRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId string `protobuf:"bytes,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeStatusRequest) GetReviewId() string {
	if x != nil {
		return x.ReviewId
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	CompanyId string `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

type ReviewItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Text string          `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	User *user.ShortUser `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ReviewItem) Reset() {
	*x = ReviewItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewItem) ProtoMessage() {}

func (x *ReviewItem) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewItem.ProtoReflect.Descriptor instead.
func (*ReviewItem) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReviewItem) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ReviewItem) GetUser() *user.ShortUser {
	if x != nil {
		return x.User
	}
	return nil
}

type GetReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*ReviewItem `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *GetReviewsResponse) Reset() {
	*x = GetReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsResponse) ProtoMessage() {}

func (x *GetReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetReviewsResponse) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{5}
}

func (x *GetReviewsResponse) GetReviews() []*ReviewItem {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type GetReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opts      *opts.SkipLimit `protobuf:"bytes,1,opt,name=opts,proto3" json:"opts,omitempty"`
	CompanyId string          `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *GetReviewsRequest) Reset() {
	*x = GetReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parser_review_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsRequest) ProtoMessage() {}

func (x *GetReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_review_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetReviewsRequest) Descriptor() ([]byte, []int) {
	return file_parser_review_proto_rawDescGZIP(), []int{6}
}

func (x *GetReviewsRequest) GetOpts() *opts.SkipLimit {
	if x != nil {
		return x.Opts
	}
	return nil
}

func (x *GetReviewsRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

var File_parser_review_proto protoreflect.FileDescriptor

var file_parser_review_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6f, 0x70, 0x74, 0x73, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x1e, 0x44, 0x69, 0x73, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x22,
	0x42, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x57,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x53, 0x6b, 0x69, 0x70, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x04, 0x6f, 0x70, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x32, 0xb0, 0x03, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x5b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12,
	0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x12,
	0x55, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x59, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71, 0x2f, 0x73, 0x63,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_parser_review_proto_rawDescOnce sync.Once
	file_parser_review_proto_rawDescData = file_parser_review_proto_rawDesc
)

func file_parser_review_proto_rawDescGZIP() []byte {
	file_parser_review_proto_rawDescOnce.Do(func() {
		file_parser_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_parser_review_proto_rawDescData)
	})
	return file_parser_review_proto_rawDescData
}

var file_parser_review_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_parser_review_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_parser_review_proto_goTypes = []interface{}{
	(ChangeStatusRequestStatus)(0),         // 0: parser.ChangeStatusRequest.status
	(*DeleteRequest)(nil),                  // 1: parser.DeleteRequest
	(*DisallowAuthorDeleteAllRequest)(nil), // 2: parser.DisallowAuthorDeleteAllRequest
	(*ChangeStatusRequest)(nil),            // 3: parser.ChangeStatusRequest
	(*CreateRequest)(nil),                  // 4: parser.CreateRequest
	(*ReviewItem)(nil),                     // 5: parser.ReviewItem
	(*GetReviewsResponse)(nil),             // 6: parser.GetReviewsResponse
	(*GetReviewsRequest)(nil),              // 7: parser.GetReviewsRequest
	(*user.ShortUser)(nil),                 // 8: user.ShortUser
	(*opts.SkipLimit)(nil),                 // 9: opts.SkipLimit
	(*emptypb.Empty)(nil),                  // 10: google.protobuf.Empty
}
var file_parser_review_proto_depIdxs = []int32{
	8,  // 0: parser.ReviewItem.user:type_name -> user.ShortUser
	5,  // 1: parser.GetReviewsResponse.reviews:type_name -> parser.ReviewItem
	9,  // 2: parser.GetReviewsRequest.opts:type_name -> opts.SkipLimit
	7,  // 3: parser.Review.GetReviews:input_type -> parser.GetReviewsRequest
	4,  // 4: parser.Review.Create:input_type -> parser.CreateRequest
	1,  // 5: parser.Review.Delete:input_type -> parser.DeleteRequest
	3,  // 6: parser.Review.ChangeStatus:input_type -> parser.ChangeStatusRequest
	2,  // 7: parser.Review.DisallowAuthorDeleteAll:input_type -> parser.DisallowAuthorDeleteAllRequest
	6,  // 8: parser.Review.GetReviews:output_type -> parser.GetReviewsResponse
	10, // 9: parser.Review.Create:output_type -> google.protobuf.Empty
	10, // 10: parser.Review.Delete:output_type -> google.protobuf.Empty
	10, // 11: parser.Review.ChangeStatus:output_type -> google.protobuf.Empty
	10, // 12: parser.Review.DisallowAuthorDeleteAll:output_type -> google.protobuf.Empty
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_parser_review_proto_init() }
func file_parser_review_proto_init() {
	if File_parser_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parser_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_parser_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisallowAuthorDeleteAllRequest); i {
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
		file_parser_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusRequest); i {
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
		file_parser_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_parser_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewItem); i {
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
		file_parser_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewsResponse); i {
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
		file_parser_review_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewsRequest); i {
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
			RawDescriptor: file_parser_review_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parser_review_proto_goTypes,
		DependencyIndexes: file_parser_review_proto_depIdxs,
		EnumInfos:         file_parser_review_proto_enumTypes,
		MessageInfos:      file_parser_review_proto_msgTypes,
	}.Build()
	File_parser_review_proto = out.File
	file_parser_review_proto_rawDesc = nil
	file_parser_review_proto_goTypes = nil
	file_parser_review_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReviewClient is the client API for Review service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewClient interface {
	GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DisallowAuthorDeleteAll(ctx context.Context, in *DisallowAuthorDeleteAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type reviewClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewClient(cc grpc.ClientConnInterface) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, "/parser.Review/GetReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/parser.Review/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/parser.Review/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/parser.Review/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewClient) DisallowAuthorDeleteAll(ctx context.Context, in *DisallowAuthorDeleteAllRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/parser.Review/DisallowAuthorDeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServer is the server API for Review service.
type ReviewServer interface {
	GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	Create(context.Context, *CreateRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*emptypb.Empty, error)
	DisallowAuthorDeleteAll(context.Context, *DisallowAuthorDeleteAllRequest) (*emptypb.Empty, error)
}

// UnimplementedReviewServer can be embedded to have forward compatible implementations.
type UnimplementedReviewServer struct {
}

func (*UnimplementedReviewServer) GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviews not implemented")
}
func (*UnimplementedReviewServer) Create(context.Context, *CreateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReviewServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedReviewServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (*UnimplementedReviewServer) DisallowAuthorDeleteAll(context.Context, *DisallowAuthorDeleteAllRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisallowAuthorDeleteAll not implemented")
}

func RegisterReviewServer(s *grpc.Server, srv ReviewServer) {
	s.RegisterService(&_Review_serviceDesc, srv)
}

func _Review_GetReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Review/GetReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Review/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Review/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Review/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Review_DisallowAuthorDeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisallowAuthorDeleteAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).DisallowAuthorDeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Review/DisallowAuthorDeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).DisallowAuthorDeleteAll(ctx, req.(*DisallowAuthorDeleteAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Review_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parser.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReviews",
			Handler:    _Review_GetReviews_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Review_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Review_Delete_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Review_ChangeStatus_Handler,
		},
		{
			MethodName: "DisallowAuthorDeleteAll",
			Handler:    _Review_DisallowAuthorDeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser/review.proto",
}
