// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: event/event.proto

package event

import (
	proto "github.com/golang/protobuf/proto"
	parser "github.com/nnqq/scr-proto/codegen/go/parser"
	wappalyzer "github.com/nnqq/scr-proto/codegen/go/wappalyzer"
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

// subject review-moderation
type ReviewModeration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Review *parser.ReviewItem `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *ReviewModeration) Reset() {
	*x = ReviewModeration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewModeration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewModeration) ProtoMessage() {}

func (x *ReviewModeration) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewModeration.ProtoReflect.Descriptor instead.
func (*ReviewModeration) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewModeration) GetReview() *parser.ReviewItem {
	if x != nil {
		return x.Review
	}
	return nil
}

// subject company-new
type CompanyNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId      string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Url            string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	AvatarToUpload string `protobuf:"bytes,3,opt,name=avatar_to_upload,json=avatarToUpload,proto3" json:"avatar_to_upload,omitempty"`
}

func (x *CompanyNew) Reset() {
	*x = CompanyNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyNew) ProtoMessage() {}

func (x *CompanyNew) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyNew.ProtoReflect.Descriptor instead.
func (*CompanyNew) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyNew) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CompanyNew) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CompanyNew) GetAvatarToUpload() string {
	if x != nil {
		return x.AvatarToUpload
	}
	return ""
}

// subject analyze-result
type AnalyzeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string                      `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Result    *wappalyzer.AnalyzeResponse `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AnalyzeResult) Reset() {
	*x = AnalyzeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalyzeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeResult) ProtoMessage() {}

func (x *AnalyzeResult) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeResult.ProtoReflect.Descriptor instead.
func (*AnalyzeResult) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *AnalyzeResult) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *AnalyzeResult) GetResult() *wappalyzer.AnalyzeResponse {
	if x != nil {
		return x.Result
	}
	return nil
}

// subject image-upload-result
type ImageUploadResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId string `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	AvatarUrl string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *ImageUploadResult) Reset() {
	*x = ImageUploadResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUploadResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUploadResult) ProtoMessage() {}

func (x *ImageUploadResult) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUploadResult.ProtoReflect.Descriptor instead.
func (*ImageUploadResult) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *ImageUploadResult) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ImageUploadResult) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

var File_event_event_proto protoreflect.FileDescriptor

var file_event_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x13, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x77, 0x61, 0x70, 0x70, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x77, 0x61, 0x70, 0x70,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x67, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x6f, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x63, 0x0a, 0x0d, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x61, 0x70, 0x70, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x51, 0x0a, 0x11, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6e, 0x71, 0x71,
	0x2f, 0x73, 0x63, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_event_event_proto_rawDescOnce sync.Once
	file_event_event_proto_rawDescData = file_event_event_proto_rawDesc
)

func file_event_event_proto_rawDescGZIP() []byte {
	file_event_event_proto_rawDescOnce.Do(func() {
		file_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_event_proto_rawDescData)
	})
	return file_event_event_proto_rawDescData
}

var file_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_event_event_proto_goTypes = []interface{}{
	(*ReviewModeration)(nil),           // 0: event.ReviewModeration
	(*CompanyNew)(nil),                 // 1: event.CompanyNew
	(*AnalyzeResult)(nil),              // 2: event.AnalyzeResult
	(*ImageUploadResult)(nil),          // 3: event.ImageUploadResult
	(*parser.ReviewItem)(nil),          // 4: parser.ReviewItem
	(*wappalyzer.AnalyzeResponse)(nil), // 5: wappalyzer.AnalyzeResponse
}
var file_event_event_proto_depIdxs = []int32{
	4, // 0: event.ReviewModeration.review:type_name -> parser.ReviewItem
	5, // 1: event.AnalyzeResult.result:type_name -> wappalyzer.AnalyzeResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_event_event_proto_init() }
func file_event_event_proto_init() {
	if File_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewModeration); i {
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
		file_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyNew); i {
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
		file_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalyzeResult); i {
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
		file_event_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUploadResult); i {
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
			RawDescriptor: file_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_event_proto_goTypes,
		DependencyIndexes: file_event_event_proto_depIdxs,
		MessageInfos:      file_event_event_proto_msgTypes,
	}.Build()
	File_event_event_proto = out.File
	file_event_event_proto_rawDesc = nil
	file_event_event_proto_goTypes = nil
	file_event_event_proto_depIdxs = nil
}
