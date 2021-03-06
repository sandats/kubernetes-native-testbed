// Code generated by protoc-gen-go. DO NOT EDIT.
// source: delivery_status.proto

package deliverystatuspb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeliveryStatus struct {
	OrderUUID            string               `protobuf:"bytes,1,opt,name=orderUUID,proto3" json:"orderUUID,omitempty"`
	UserUUID             string               `protobuf:"bytes,2,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	Status               int32                `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	InquiryNumber        string               `protobuf:"bytes,4,opt,name=inquiryNumber,proto3" json:"inquiryNumber,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeliveryStatus) Reset()         { *m = DeliveryStatus{} }
func (m *DeliveryStatus) String() string { return proto.CompactTextString(m) }
func (*DeliveryStatus) ProtoMessage()    {}
func (*DeliveryStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{0}
}

func (m *DeliveryStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliveryStatus.Unmarshal(m, b)
}
func (m *DeliveryStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliveryStatus.Marshal(b, m, deterministic)
}
func (m *DeliveryStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliveryStatus.Merge(m, src)
}
func (m *DeliveryStatus) XXX_Size() int {
	return xxx_messageInfo_DeliveryStatus.Size(m)
}
func (m *DeliveryStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliveryStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeliveryStatus proto.InternalMessageInfo

func (m *DeliveryStatus) GetOrderUUID() string {
	if m != nil {
		return m.OrderUUID
	}
	return ""
}

func (m *DeliveryStatus) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *DeliveryStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DeliveryStatus) GetInquiryNumber() string {
	if m != nil {
		return m.InquiryNumber
	}
	return ""
}

func (m *DeliveryStatus) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *DeliveryStatus) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *DeliveryStatus) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type GetRequest struct {
	OrderUUID            string   `protobuf:"bytes,1,opt,name=orderUUID,proto3" json:"orderUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetOrderUUID() string {
	if m != nil {
		return m.OrderUUID
	}
	return ""
}

type GetResponse struct {
	DeliveryStatus       *DeliveryStatus `protobuf:"bytes,1,opt,name=deliveryStatus,proto3" json:"deliveryStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetDeliveryStatus() *DeliveryStatus {
	if m != nil {
		return m.DeliveryStatus
	}
	return nil
}

type SetRequest struct {
	DeliveryStatus       *DeliveryStatus `protobuf:"bytes,1,opt,name=deliveryStatus,proto3" json:"deliveryStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{3}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetDeliveryStatus() *DeliveryStatus {
	if m != nil {
		return m.DeliveryStatus
	}
	return nil
}

type SetResponse struct {
	OrderUUID            string   `protobuf:"bytes,1,opt,name=orderUUID,proto3" json:"orderUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{4}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetOrderUUID() string {
	if m != nil {
		return m.OrderUUID
	}
	return ""
}

type UpdateRequest struct {
	DeliveryStatus       *DeliveryStatus `protobuf:"bytes,1,opt,name=deliveryStatus,proto3" json:"deliveryStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetDeliveryStatus() *DeliveryStatus {
	if m != nil {
		return m.DeliveryStatus
	}
	return nil
}

type DeleteRequest struct {
	OrderUUID            string   `protobuf:"bytes,1,opt,name=orderUUID,proto3" json:"orderUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bce9cfbc17985f1, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetOrderUUID() string {
	if m != nil {
		return m.OrderUUID
	}
	return ""
}

func init() {
	proto.RegisterType((*DeliveryStatus)(nil), "deliverystatuspb.DeliveryStatus")
	proto.RegisterType((*GetRequest)(nil), "deliverystatuspb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "deliverystatuspb.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "deliverystatuspb.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "deliverystatuspb.SetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "deliverystatuspb.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "deliverystatuspb.DeleteRequest")
}

func init() { proto.RegisterFile("delivery_status.proto", fileDescriptor_0bce9cfbc17985f1) }

var fileDescriptor_0bce9cfbc17985f1 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4f, 0xab, 0x40,
	0x14, 0xc5, 0x1f, 0xed, 0x2b, 0xef, 0xf5, 0x92, 0x36, 0xef, 0x4d, 0x62, 0x43, 0xb0, 0xa6, 0x84,
	0xb8, 0x68, 0x34, 0xd2, 0xa4, 0x6e, 0xdc, 0x36, 0x62, 0x6a, 0x37, 0xc6, 0x80, 0xd5, 0xb8, 0x32,
	0x45, 0xae, 0x0d, 0x49, 0x29, 0x14, 0x06, 0x93, 0x7e, 0x24, 0xbf, 0x80, 0x9f, 0xcf, 0x30, 0x43,
	0x8b, 0xd3, 0xbf, 0x2e, 0xba, 0x9c, 0x7b, 0xcf, 0xf9, 0x71, 0x38, 0x03, 0x70, 0xe4, 0xe1, 0xc4,
	0x7f, 0xc7, 0x78, 0xfe, 0x92, 0xd0, 0x11, 0x4d, 0x13, 0x33, 0x8a, 0x43, 0x1a, 0x92, 0x7f, 0x8b,
	0x31, 0x9f, 0x46, 0xae, 0x76, 0x3c, 0x0e, 0xc3, 0xf1, 0x04, 0x3b, 0x6c, 0xef, 0xa6, 0x6f, 0x1d,
	0x0c, 0x22, 0x3a, 0xe7, 0x72, 0xad, 0xb5, 0xba, 0xa4, 0x7e, 0x80, 0x09, 0x1d, 0x05, 0x11, 0x17,
	0x18, 0x9f, 0x25, 0xa8, 0x5b, 0x39, 0xd2, 0x61, 0x48, 0xd2, 0x84, 0x6a, 0x18, 0x7b, 0x18, 0x0f,
	0x87, 0x03, 0x4b, 0x95, 0x74, 0xa9, 0x5d, 0xb5, 0x8b, 0x01, 0xd1, 0xe0, 0x6f, 0x9a, 0xe4, 0xcb,
	0x12, 0x5b, 0x2e, 0xcf, 0xa4, 0x01, 0x32, 0x8f, 0xa5, 0x96, 0x75, 0xa9, 0x5d, 0xb1, 0xf3, 0x13,
	0x39, 0x85, 0x9a, 0x3f, 0x9d, 0xa5, 0x7e, 0x3c, 0xbf, 0x4b, 0x03, 0x17, 0x63, 0xf5, 0x37, 0x33,
	0x8a, 0x43, 0x72, 0x05, 0xd5, 0xd7, 0x18, 0x47, 0x14, 0xbd, 0x1e, 0x55, 0x2b, 0xba, 0xd4, 0x56,
	0xba, 0x9a, 0xc9, 0xf3, 0x9b, 0x8b, 0xfc, 0xe6, 0xc3, 0x22, 0xbf, 0x5d, 0x88, 0x33, 0x67, 0x1a,
	0x79, 0xb9, 0x53, 0xde, 0xef, 0x5c, 0x8a, 0x33, 0xa7, 0x87, 0x13, 0xe4, 0xce, 0x3f, 0xfb, 0x9d,
	0x4b, 0xb1, 0x71, 0x06, 0xd0, 0x47, 0x6a, 0xe3, 0x2c, 0xc5, 0x84, 0xee, 0xee, 0xcc, 0x78, 0x02,
	0x85, 0x69, 0x93, 0x28, 0x9c, 0x26, 0x48, 0x6e, 0xa1, 0xee, 0x09, 0x95, 0x33, 0x87, 0xd2, 0xd5,
	0xcd, 0xd5, 0xcb, 0x35, 0xc5, 0xab, 0xb1, 0x57, 0x7c, 0xc6, 0x23, 0x80, 0x53, 0x84, 0x38, 0x1c,
	0xf7, 0x1c, 0x14, 0xe7, 0x5b, 0xe0, 0xdd, 0x6f, 0xf7, 0x0c, 0xb5, 0x21, 0x2b, 0xf4, 0xf0, 0x39,
	0x2e, 0xa0, 0x66, 0xb1, 0xc6, 0x7f, 0xd4, 0x73, 0xf7, 0xa3, 0x04, 0xff, 0x45, 0x62, 0xef, 0x7e,
	0x40, 0x2c, 0x28, 0xf7, 0x91, 0x92, 0xe6, 0xfa, 0xd3, 0x8b, 0x0b, 0xd4, 0x4e, 0xb6, 0x6c, 0x79,
	0x03, 0xc6, 0xaf, 0x8c, 0xe2, 0x6c, 0xa6, 0x38, 0x3b, 0x29, 0x8e, 0x40, 0xb9, 0x06, 0x99, 0x77,
	0x45, 0x5a, 0xeb, 0x52, 0xa1, 0x45, 0xad, 0xb1, 0xf6, 0x1d, 0xde, 0x64, 0x3f, 0x36, 0x87, 0xf0,
	0x56, 0x36, 0x41, 0x84, 0xbe, 0xb6, 0x43, 0x5c, 0x99, 0x4d, 0x2e, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0x9d, 0xa7, 0xd3, 0x68, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeliveryStatusAPIClient is the client API for DeliveryStatusAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeliveryStatusAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type deliveryStatusAPIClient struct {
	cc *grpc.ClientConn
}

func NewDeliveryStatusAPIClient(cc *grpc.ClientConn) DeliveryStatusAPIClient {
	return &deliveryStatusAPIClient{cc}
}

func (c *deliveryStatusAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/deliverystatuspb.DeliveryStatusAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryStatusAPIClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/deliverystatuspb.DeliveryStatusAPI/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryStatusAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/deliverystatuspb.DeliveryStatusAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryStatusAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/deliverystatuspb.DeliveryStatusAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryStatusAPIServer is the server API for DeliveryStatusAPI service.
type DeliveryStatusAPIServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedDeliveryStatusAPIServer can be embedded to have forward compatible implementations.
type UnimplementedDeliveryStatusAPIServer struct {
}

func (*UnimplementedDeliveryStatusAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDeliveryStatusAPIServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedDeliveryStatusAPIServer) Update(ctx context.Context, req *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDeliveryStatusAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterDeliveryStatusAPIServer(s *grpc.Server, srv DeliveryStatusAPIServer) {
	s.RegisterService(&_DeliveryStatusAPI_serviceDesc, srv)
}

func _DeliveryStatusAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryStatusAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliverystatuspb.DeliveryStatusAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryStatusAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryStatusAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryStatusAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliverystatuspb.DeliveryStatusAPI/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryStatusAPIServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryStatusAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryStatusAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliverystatuspb.DeliveryStatusAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryStatusAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryStatusAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryStatusAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deliverystatuspb.DeliveryStatusAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryStatusAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeliveryStatusAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deliverystatuspb.DeliveryStatusAPI",
	HandlerType: (*DeliveryStatusAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DeliveryStatusAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _DeliveryStatusAPI_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeliveryStatusAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeliveryStatusAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery_status.proto",
}
