// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: office.proto

package customer

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOfficeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreateOfficeRequest) Reset() {
	*x = CreateOfficeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfficeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfficeRequest) ProtoMessage() {}

func (x *CreateOfficeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfficeRequest.ProtoReflect.Descriptor instead.
func (*CreateOfficeRequest) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOfficeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateOfficeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreateOfficeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOfficeResponse) Reset() {
	*x = CreateOfficeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfficeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfficeResponse) ProtoMessage() {}

func (x *CreateOfficeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfficeResponse.ProtoReflect.Descriptor instead.
func (*CreateOfficeResponse) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{1}
}

type GetOfficeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOfficeListRequest) Reset() {
	*x = GetOfficeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficeListRequest) ProtoMessage() {}

func (x *GetOfficeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficeListRequest.ProtoReflect.Descriptor instead.
func (*GetOfficeListRequest) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{2}
}

type GetOfficeListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*Office `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetOfficeListResponse) Reset() {
	*x = GetOfficeListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficeListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficeListResponse) ProtoMessage() {}

func (x *GetOfficeListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficeListResponse.ProtoReflect.Descriptor instead.
func (*GetOfficeListResponse) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{3}
}

func (x *GetOfficeListResponse) GetResult() []*Office {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetOfficeByUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficeUuid string `protobuf:"bytes,1,opt,name=office_uuid,json=officeUuid,proto3" json:"office_uuid,omitempty"`
}

func (x *GetOfficeByUUIDRequest) Reset() {
	*x = GetOfficeByUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficeByUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficeByUUIDRequest) ProtoMessage() {}

func (x *GetOfficeByUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficeByUUIDRequest.ProtoReflect.Descriptor instead.
func (*GetOfficeByUUIDRequest) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{4}
}

func (x *GetOfficeByUUIDRequest) GetOfficeUuid() string {
	if x != nil {
		return x.OfficeUuid
	}
	return ""
}

type GetOfficeByUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Office `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetOfficeByUUIDResponse) Reset() {
	*x = GetOfficeByUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOfficeByUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOfficeByUUIDResponse) ProtoMessage() {}

func (x *GetOfficeByUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOfficeByUUIDResponse.ProtoReflect.Descriptor instead.
func (*GetOfficeByUUIDResponse) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{5}
}

func (x *GetOfficeByUUIDResponse) GetResult() *Office {
	if x != nil {
		return x.Result
	}
	return nil
}

type Office struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Office) Reset() {
	*x = Office{}
	if protoimpl.UnsafeEnabled {
		mi := &file_office_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Office) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Office) ProtoMessage() {}

func (x *Office) ProtoReflect() protoreflect.Message {
	mi := &file_office_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Office.ProtoReflect.Descriptor instead.
func (*Office) Descriptor() ([]byte, []int) {
	return file_office_proto_rawDescGZIP(), []int{6}
}

func (x *Office) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Office) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Office) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Office) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_office_proto protoreflect.FileDescriptor

var file_office_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x55, 0x75, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x85, 0x01, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xe1, 0x02, 0x0a, 0x0d, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x12, 0x6b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x6f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x53, 0x46, 0x54, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_office_proto_rawDescOnce sync.Once
	file_office_proto_rawDescData = file_office_proto_rawDesc
)

func file_office_proto_rawDescGZIP() []byte {
	file_office_proto_rawDescOnce.Do(func() {
		file_office_proto_rawDescData = protoimpl.X.CompressGZIP(file_office_proto_rawDescData)
	})
	return file_office_proto_rawDescData
}

var file_office_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_office_proto_goTypes = []interface{}{
	(*CreateOfficeRequest)(nil),     // 0: customer.CreateOfficeRequest
	(*CreateOfficeResponse)(nil),    // 1: customer.CreateOfficeResponse
	(*GetOfficeListRequest)(nil),    // 2: customer.GetOfficeListRequest
	(*GetOfficeListResponse)(nil),   // 3: customer.GetOfficeListResponse
	(*GetOfficeByUUIDRequest)(nil),  // 4: customer.GetOfficeByUUIDRequest
	(*GetOfficeByUUIDResponse)(nil), // 5: customer.GetOfficeByUUIDResponse
	(*Office)(nil),                  // 6: customer.Office
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_office_proto_depIdxs = []int32{
	6, // 0: customer.GetOfficeListResponse.result:type_name -> customer.Office
	6, // 1: customer.GetOfficeByUUIDResponse.result:type_name -> customer.Office
	7, // 2: customer.Office.created_at:type_name -> google.protobuf.Timestamp
	0, // 3: customer.OfficeService.CreateOffice:input_type -> customer.CreateOfficeRequest
	2, // 4: customer.OfficeService.GetOfficeList:input_type -> customer.GetOfficeListRequest
	4, // 5: customer.OfficeService.GetOfficeByUUID:input_type -> customer.GetOfficeByUUIDRequest
	1, // 6: customer.OfficeService.CreateOffice:output_type -> customer.CreateOfficeResponse
	3, // 7: customer.OfficeService.GetOfficeList:output_type -> customer.GetOfficeListResponse
	5, // 8: customer.OfficeService.GetOfficeByUUID:output_type -> customer.GetOfficeByUUIDResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_office_proto_init() }
func file_office_proto_init() {
	if File_office_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_office_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOfficeRequest); i {
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
		file_office_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOfficeResponse); i {
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
		file_office_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOfficeListRequest); i {
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
		file_office_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOfficeListResponse); i {
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
		file_office_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOfficeByUUIDRequest); i {
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
		file_office_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOfficeByUUIDResponse); i {
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
		file_office_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Office); i {
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
			RawDescriptor: file_office_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_office_proto_goTypes,
		DependencyIndexes: file_office_proto_depIdxs,
		MessageInfos:      file_office_proto_msgTypes,
	}.Build()
	File_office_proto = out.File
	file_office_proto_rawDesc = nil
	file_office_proto_goTypes = nil
	file_office_proto_depIdxs = nil
}
