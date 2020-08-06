// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: city/city.proto

package city

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type GetHintsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetHintsRequest) Reset() {
	*x = GetHintsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHintsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHintsRequest) ProtoMessage() {}

func (x *GetHintsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHintsRequest.ProtoReflect.Descriptor instead.
func (*GetHintsRequest) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{0}
}

func (x *GetHintsRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetHintsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetBySlugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *GetBySlugRequest) Reset() {
	*x = GetBySlugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySlugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySlugRequest) ProtoMessage() {}

func (x *GetBySlugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySlugRequest.ProtoReflect.Descriptor instead.
func (*GetBySlugRequest) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{1}
}

func (x *GetBySlugRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId string `protobuf:"bytes,1,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIdRequest) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

type GetByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityIds []string `protobuf:"bytes,1,rep,name=city_ids,json=cityIds,proto3" json:"city_ids,omitempty"`
}

func (x *GetByIdsRequest) Reset() {
	*x = GetByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdsRequest) ProtoMessage() {}

func (x *GetByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdsRequest.ProtoReflect.Descriptor instead.
func (*GetByIdsRequest) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{3}
}

func (x *GetByIdsRequest) GetCityIds() []string {
	if x != nil {
		return x.CityIds
	}
	return nil
}

type CitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cities []*CityItem `protobuf:"bytes,1,rep,name=cities,proto3" json:"cities,omitempty"`
}

func (x *CitiesResponse) Reset() {
	*x = CitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitiesResponse) ProtoMessage() {}

func (x *CitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitiesResponse.ProtoReflect.Descriptor instead.
func (*CitiesResponse) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{4}
}

func (x *CitiesResponse) GetCities() []*CityItem {
	if x != nil {
		return x.Cities
	}
	return nil
}

type FindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *FindRequest) Reset() {
	*x = FindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRequest) ProtoMessage() {}

func (x *FindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[5]
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
	return file_city_city_proto_rawDescGZIP(), []int{5}
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

	CityId  string `protobuf:"bytes,1,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	IsFound bool   `protobuf:"varint,2,opt,name=is_found,json=isFound,proto3" json:"is_found,omitempty"`
}

func (x *FindResponse) Reset() {
	*x = FindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindResponse) ProtoMessage() {}

func (x *FindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[6]
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
	return file_city_city_proto_rawDescGZIP(), []int{6}
}

func (x *FindResponse) GetCityId() string {
	if x != nil {
		return x.CityId
	}
	return ""
}

func (x *FindResponse) GetIsFound() bool {
	if x != nil {
		return x.IsFound
	}
	return false
}

type CityItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug  string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *CityItem) Reset() {
	*x = CityItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_city_city_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityItem) ProtoMessage() {}

func (x *CityItem) ProtoReflect() protoreflect.Message {
	mi := &file_city_city_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityItem.ProtoReflect.Descriptor instead.
func (*CityItem) Descriptor() ([]byte, []int) {
	return file_city_city_proto_rawDescGZIP(), []int{7}
}

func (x *CityItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CityItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CityItem) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_city_city_proto protoreflect.FileDescriptor

var file_city_city_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x69, 0x74, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x21, 0x92, 0x41, 0x1e, 0x32, 0x0a, 0x32,
	0x30, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x59, 0x40, 0x69, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x21, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x74, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c,
	0x22, 0x42, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x44, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x32, 0xca, 0x03, 0x0a, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x69,
	0x74, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e,
	0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x49,
	0x74, 0x65, 0x6d, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4f, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x16, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x52,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x15, 0x2e, 0x63, 0x69, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12,
	0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x73, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x52, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x15, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x67,
	0x65, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71, 0x2f, 0x73, 0x63, 0x72, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_city_city_proto_rawDescOnce sync.Once
	file_city_city_proto_rawDescData = file_city_city_proto_rawDesc
)

func file_city_city_proto_rawDescGZIP() []byte {
	file_city_city_proto_rawDescOnce.Do(func() {
		file_city_city_proto_rawDescData = protoimpl.X.CompressGZIP(file_city_city_proto_rawDescData)
	})
	return file_city_city_proto_rawDescData
}

var file_city_city_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_city_city_proto_goTypes = []interface{}{
	(*GetHintsRequest)(nil),  // 0: city.GetHintsRequest
	(*GetBySlugRequest)(nil), // 1: city.GetBySlugRequest
	(*GetByIdRequest)(nil),   // 2: city.GetByIdRequest
	(*GetByIdsRequest)(nil),  // 3: city.GetByIdsRequest
	(*CitiesResponse)(nil),   // 4: city.CitiesResponse
	(*FindRequest)(nil),      // 5: city.FindRequest
	(*FindResponse)(nil),     // 6: city.FindResponse
	(*CityItem)(nil),         // 7: city.CityItem
	(*empty.Empty)(nil),      // 8: google.protobuf.Empty
}
var file_city_city_proto_depIdxs = []int32{
	7, // 0: city.CitiesResponse.cities:type_name -> city.CityItem
	5, // 1: city.City.Find:input_type -> city.FindRequest
	2, // 2: city.City.GetById:input_type -> city.GetByIdRequest
	1, // 3: city.City.GetBySlug:input_type -> city.GetBySlugRequest
	3, // 4: city.City.GetByIds:input_type -> city.GetByIdsRequest
	8, // 5: city.City.GetAll:input_type -> google.protobuf.Empty
	0, // 6: city.City.GetHints:input_type -> city.GetHintsRequest
	6, // 7: city.City.Find:output_type -> city.FindResponse
	7, // 8: city.City.GetById:output_type -> city.CityItem
	7, // 9: city.City.GetBySlug:output_type -> city.CityItem
	4, // 10: city.City.GetByIds:output_type -> city.CitiesResponse
	4, // 11: city.City.GetAll:output_type -> city.CitiesResponse
	4, // 12: city.City.GetHints:output_type -> city.CitiesResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_city_city_proto_init() }
func file_city_city_proto_init() {
	if File_city_city_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_city_city_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHintsRequest); i {
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
		file_city_city_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySlugRequest); i {
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
		file_city_city_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_city_city_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdsRequest); i {
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
		file_city_city_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitiesResponse); i {
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
		file_city_city_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_city_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_city_city_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityItem); i {
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
			RawDescriptor: file_city_city_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_city_proto_goTypes,
		DependencyIndexes: file_city_city_proto_depIdxs,
		MessageInfos:      file_city_city_proto_msgTypes,
	}.Build()
	File_city_city_proto = out.File
	file_city_city_proto_rawDesc = nil
	file_city_city_proto_goTypes = nil
	file_city_city_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CityClient is the client API for City service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CityClient interface {
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CityItem, error)
	GetBySlug(ctx context.Context, in *GetBySlugRequest, opts ...grpc.CallOption) (*CityItem, error)
	GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*CitiesResponse, error)
	GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CitiesResponse, error)
	GetHints(ctx context.Context, in *GetHintsRequest, opts ...grpc.CallOption) (*CitiesResponse, error)
}

type cityClient struct {
	cc grpc.ClientConnInterface
}

func NewCityClient(cc grpc.ClientConnInterface) CityClient {
	return &cityClient{cc}
}

func (c *cityClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/city.City/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CityItem, error) {
	out := new(CityItem)
	err := c.cc.Invoke(ctx, "/city.City/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityClient) GetBySlug(ctx context.Context, in *GetBySlugRequest, opts ...grpc.CallOption) (*CityItem, error) {
	out := new(CityItem)
	err := c.cc.Invoke(ctx, "/city.City/GetBySlug", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityClient) GetByIds(ctx context.Context, in *GetByIdsRequest, opts ...grpc.CallOption) (*CitiesResponse, error) {
	out := new(CitiesResponse)
	err := c.cc.Invoke(ctx, "/city.City/GetByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityClient) GetAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CitiesResponse, error) {
	out := new(CitiesResponse)
	err := c.cc.Invoke(ctx, "/city.City/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityClient) GetHints(ctx context.Context, in *GetHintsRequest, opts ...grpc.CallOption) (*CitiesResponse, error) {
	out := new(CitiesResponse)
	err := c.cc.Invoke(ctx, "/city.City/GetHints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityServer is the server API for City service.
type CityServer interface {
	Find(context.Context, *FindRequest) (*FindResponse, error)
	GetById(context.Context, *GetByIdRequest) (*CityItem, error)
	GetBySlug(context.Context, *GetBySlugRequest) (*CityItem, error)
	GetByIds(context.Context, *GetByIdsRequest) (*CitiesResponse, error)
	GetAll(context.Context, *empty.Empty) (*CitiesResponse, error)
	GetHints(context.Context, *GetHintsRequest) (*CitiesResponse, error)
}

// UnimplementedCityServer can be embedded to have forward compatible implementations.
type UnimplementedCityServer struct {
}

func (*UnimplementedCityServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedCityServer) GetById(context.Context, *GetByIdRequest) (*CityItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (*UnimplementedCityServer) GetBySlug(context.Context, *GetBySlugRequest) (*CityItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySlug not implemented")
}
func (*UnimplementedCityServer) GetByIds(context.Context, *GetByIdsRequest) (*CitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIds not implemented")
}
func (*UnimplementedCityServer) GetAll(context.Context, *empty.Empty) (*CitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCityServer) GetHints(context.Context, *GetHintsRequest) (*CitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHints not implemented")
}

func RegisterCityServer(s *grpc.Server, srv CityServer) {
	s.RegisterService(&_City_serviceDesc, srv)
}

func _City_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _City_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _City_GetBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySlugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).GetBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/GetBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).GetBySlug(ctx, req.(*GetBySlugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _City_GetByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).GetByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/GetByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).GetByIds(ctx, req.(*GetByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _City_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).GetAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _City_GetHints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHintsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServer).GetHints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.City/GetHints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServer).GetHints(ctx, req.(*GetHintsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _City_serviceDesc = grpc.ServiceDesc{
	ServiceName: "city.City",
	HandlerType: (*CityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _City_Find_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _City_GetById_Handler,
		},
		{
			MethodName: "GetBySlug",
			Handler:    _City_GetBySlug_Handler,
		},
		{
			MethodName: "GetByIds",
			Handler:    _City_GetByIds_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _City_GetAll_Handler,
		},
		{
			MethodName: "GetHints",
			Handler:    _City_GetHints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "city/city.proto",
}
