// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: common.proto

package clusterpb

import (
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

type PageDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total    int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PageDTO) Reset() {
	*x = PageDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDTO) ProtoMessage() {}

func (x *PageDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDTO.ProtoReflect.Descriptor instead.
func (*PageDTO) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *PageDTO) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageDTO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageDTO) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ResponseStatusDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseStatusDTO) Reset() {
	*x = ResponseStatusDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseStatusDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseStatusDTO) ProtoMessage() {}

func (x *ResponseStatusDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseStatusDTO.ProtoReflect.Descriptor instead.
func (*ResponseStatusDTO) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseStatusDTO) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseStatusDTO) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UsageDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     float32 `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Used      float32 `protobuf:"fixed32,2,opt,name=used,proto3" json:"used,omitempty"`
	UsageRate float32 `protobuf:"fixed32,3,opt,name=usageRate,proto3" json:"usageRate,omitempty"`
}

func (x *UsageDTO) Reset() {
	*x = UsageDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageDTO) ProtoMessage() {}

func (x *UsageDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageDTO.ProtoReflect.Descriptor instead.
func (*UsageDTO) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *UsageDTO) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UsageDTO) GetUsed() float32 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *UsageDTO) GetUsageRate() float32 {
	if x != nil {
		return x.UsageRate
	}
	return 0
}

type OperatorDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TenantId       string `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	ManualOperator bool   `protobuf:"varint,4,opt,name=manualOperator,proto3" json:"manualOperator,omitempty"`
}

func (x *OperatorDTO) Reset() {
	*x = OperatorDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorDTO) ProtoMessage() {}

func (x *OperatorDTO) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorDTO.ProtoReflect.Descriptor instead.
func (*OperatorDTO) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *OperatorDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OperatorDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OperatorDTO) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *OperatorDTO) GetManualOperator() bool {
	if x != nil {
		return x.ManualOperator
	}
	return false
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x41, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x52, 0x0a, 0x08, 0x55, 0x73, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0x75, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x1b, 0x5a,
	0x19, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f,
	0x3b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(*PageDTO)(nil),           // 0: PageDTO
	(*ResponseStatusDTO)(nil), // 1: ResponseStatusDTO
	(*UsageDTO)(nil),          // 2: UsageDTO
	(*OperatorDTO)(nil),       // 3: OperatorDTO
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDTO); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseStatusDTO); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageDTO); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorDTO); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
