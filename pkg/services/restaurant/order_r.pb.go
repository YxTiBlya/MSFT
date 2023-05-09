// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: order_r.proto

package restaurant

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUpToDateOrderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUpToDateOrderListRequest) Reset() {
	*x = GetUpToDateOrderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_r_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpToDateOrderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpToDateOrderListRequest) ProtoMessage() {}

func (x *GetUpToDateOrderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_r_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpToDateOrderListRequest.ProtoReflect.Descriptor instead.
func (*GetUpToDateOrderListRequest) Descriptor() ([]byte, []int) {
	return file_order_r_proto_rawDescGZIP(), []int{0}
}

type GetUpToDateOrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalOrders          []*Order          `protobuf:"bytes,1,rep,name=total_orders,json=totalOrders,proto3" json:"total_orders,omitempty"`
	TotalOrdersByCompany []*OrdersByOffice `protobuf:"bytes,2,rep,name=total_orders_by_company,json=totalOrdersByCompany,proto3" json:"total_orders_by_company,omitempty"`
}

func (x *GetUpToDateOrderListResponse) Reset() {
	*x = GetUpToDateOrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_r_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpToDateOrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpToDateOrderListResponse) ProtoMessage() {}

func (x *GetUpToDateOrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_r_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpToDateOrderListResponse.ProtoReflect.Descriptor instead.
func (*GetUpToDateOrderListResponse) Descriptor() ([]byte, []int) {
	return file_order_r_proto_rawDescGZIP(), []int{1}
}

func (x *GetUpToDateOrderListResponse) GetTotalOrders() []*Order {
	if x != nil {
		return x.TotalOrders
	}
	return nil
}

func (x *GetUpToDateOrderListResponse) GetTotalOrdersByCompany() []*OrdersByOffice {
	if x != nil {
		return x.TotalOrdersByCompany
	}
	return nil
}

type OrdersByOffice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId     string   `protobuf:"bytes,1,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	OfficeName    string   `protobuf:"bytes,2,opt,name=office_name,json=officeName,proto3" json:"office_name,omitempty"`
	OfficeAddress string   `protobuf:"bytes,3,opt,name=office_address,json=officeAddress,proto3" json:"office_address,omitempty"`
	Result        []*Order `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *OrdersByOffice) Reset() {
	*x = OrdersByOffice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_r_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersByOffice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersByOffice) ProtoMessage() {}

func (x *OrdersByOffice) ProtoReflect() protoreflect.Message {
	mi := &file_order_r_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersByOffice.ProtoReflect.Descriptor instead.
func (*OrdersByOffice) Descriptor() ([]byte, []int) {
	return file_order_r_proto_rawDescGZIP(), []int{2}
}

func (x *OrdersByOffice) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *OrdersByOffice) GetOfficeName() string {
	if x != nil {
		return x.OfficeName
	}
	return ""
}

func (x *OrdersByOffice) GetOfficeAddress() string {
	if x != nil {
		return x.OfficeAddress
	}
	return ""
}

func (x *OrdersByOffice) GetResult() []*Order {
	if x != nil {
		return x.Result
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Count       int64  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_r_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_r_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_r_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Order) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_order_r_proto protoreflect.FileDescriptor

var file_order_r_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x51, 0x0a, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x14,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x5f, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x96, 0x01, 0x0a, 0x0c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70,
	0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12,
	0x12, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x53, 0x46, 0x54, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_r_proto_rawDescOnce sync.Once
	file_order_r_proto_rawDescData = file_order_r_proto_rawDesc
)

func file_order_r_proto_rawDescGZIP() []byte {
	file_order_r_proto_rawDescOnce.Do(func() {
		file_order_r_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_r_proto_rawDescData)
	})
	return file_order_r_proto_rawDescData
}

var file_order_r_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_order_r_proto_goTypes = []interface{}{
	(*GetUpToDateOrderListRequest)(nil),  // 0: restaurant.GetUpToDateOrderListRequest
	(*GetUpToDateOrderListResponse)(nil), // 1: restaurant.GetUpToDateOrderListResponse
	(*OrdersByOffice)(nil),               // 2: restaurant.OrdersByOffice
	(*Order)(nil),                        // 3: restaurant.Order
}
var file_order_r_proto_depIdxs = []int32{
	3, // 0: restaurant.GetUpToDateOrderListResponse.total_orders:type_name -> restaurant.Order
	2, // 1: restaurant.GetUpToDateOrderListResponse.total_orders_by_company:type_name -> restaurant.OrdersByOffice
	3, // 2: restaurant.OrdersByOffice.result:type_name -> restaurant.Order
	0, // 3: restaurant.OrderService.GetUpToDateOrderList:input_type -> restaurant.GetUpToDateOrderListRequest
	1, // 4: restaurant.OrderService.GetUpToDateOrderList:output_type -> restaurant.GetUpToDateOrderListResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_r_proto_init() }
func file_order_r_proto_init() {
	if File_order_r_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_r_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpToDateOrderListRequest); i {
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
		file_order_r_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpToDateOrderListResponse); i {
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
		file_order_r_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersByOffice); i {
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
		file_order_r_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_order_r_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_r_proto_goTypes,
		DependencyIndexes: file_order_r_proto_depIdxs,
		MessageInfos:      file_order_r_proto_msgTypes,
	}.Build()
	File_order_r_proto = out.File
	file_order_r_proto_rawDesc = nil
	file_order_r_proto_goTypes = nil
	file_order_r_proto_depIdxs = nil
}