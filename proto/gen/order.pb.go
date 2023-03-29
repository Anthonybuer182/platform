// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: order.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GetListDeleteOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListDeleteOrdersRequest) Reset() {
	*x = GetListDeleteOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListDeleteOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListDeleteOrdersRequest) ProtoMessage() {}

func (x *GetListDeleteOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListDeleteOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetListDeleteOrdersRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

type GetListOrderDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*OrderDto `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetListOrderDeleteResponse) Reset() {
	*x = GetListOrderDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOrderDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOrderDeleteResponse) ProtoMessage() {}

func (x *GetListOrderDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOrderDeleteResponse.ProtoReflect.Descriptor instead.
func (*GetListOrderDeleteResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *GetListOrderDeleteResponse) GetOrders() []*OrderDto {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderNum    string        `protobuf:"bytes,2,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	OrderStatus string        `protobuf:"bytes,3,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	Products    []*ProductDto `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	Users       []*UserDto    `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *OrderDto) Reset() {
	*x = OrderDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDto) ProtoMessage() {}

func (x *OrderDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDto.ProtoReflect.Descriptor instead.
func (*OrderDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderDto) GetOrderNum() string {
	if x != nil {
		return x.OrderNum
	}
	return ""
}

func (x *OrderDto) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *OrderDto) GetProducts() []*ProductDto {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *OrderDto) GetUsers() []*UserDto {
	if x != nil {
		return x.Users
	}
	return nil
}

type ProductDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sku   string  `protobuf:"bytes,2,opt,name=sku,proto3" json:"sku,omitempty"`
	Name  string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ProductDto) Reset() {
	*x = ProductDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDto) ProtoMessage() {}

func (x *ProductDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDto.ProtoReflect.Descriptor instead.
func (*ProductDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *ProductDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductDto) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *ProductDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDto) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Telephone string `protobuf:"bytes,3,opt,name=telephone,proto3" json:"telephone,omitempty"`
}

func (x *UserDto) Reset() {
	*x = UserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDto) ProtoMessage() {}

func (x *UserDto) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDto.ProtoReflect.Descriptor instead.
func (*UserDto) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *UserDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDto) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67,
	0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x22, 0xdd, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x58, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44,
	0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4b,
	0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32, 0xf9, 0x01, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe8, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x38, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38,
	0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x3d, 0x0a, 0x06, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6e, 0x67, 0x63, 0x68, 0x75, 0x6e, 0x67,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_proto_goTypes = []interface{}{
	(*GetListDeleteOrdersRequest)(nil), // 0: go.coffeeshop.proto.orderapi.GetListDeleteOrdersRequest
	(*GetListOrderDeleteResponse)(nil), // 1: go.coffeeshop.proto.orderapi.GetListOrderDeleteResponse
	(*OrderDto)(nil),                   // 2: go.coffeeshop.proto.orderapi.OrderDto
	(*ProductDto)(nil),                 // 3: go.coffeeshop.proto.orderapi.ProductDto
	(*UserDto)(nil),                    // 4: go.coffeeshop.proto.orderapi.UserDto
}
var file_order_proto_depIdxs = []int32{
	2, // 0: go.coffeeshop.proto.orderapi.GetListOrderDeleteResponse.orders:type_name -> go.coffeeshop.proto.orderapi.OrderDto
	3, // 1: go.coffeeshop.proto.orderapi.OrderDto.products:type_name -> go.coffeeshop.proto.orderapi.ProductDto
	4, // 2: go.coffeeshop.proto.orderapi.OrderDto.users:type_name -> go.coffeeshop.proto.orderapi.UserDto
	0, // 3: go.coffeeshop.proto.orderapi.OrderService.GetListDeleteOrders:input_type -> go.coffeeshop.proto.orderapi.GetListDeleteOrdersRequest
	1, // 4: go.coffeeshop.proto.orderapi.OrderService.GetListDeleteOrders:output_type -> go.coffeeshop.proto.orderapi.GetListOrderDeleteResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListDeleteOrdersRequest); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOrderDeleteResponse); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDto); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDto); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDto); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
