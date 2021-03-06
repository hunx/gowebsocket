// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/payments_account_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Request message for fetching all accessible Payments accounts.
type ListPaymentsAccountsRequest struct {
	// The ID of the customer to apply the PaymentsAccount list operation to.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPaymentsAccountsRequest) Reset()         { *m = ListPaymentsAccountsRequest{} }
func (m *ListPaymentsAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsRequest) ProtoMessage()    {}
func (*ListPaymentsAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{0}
}

func (m *ListPaymentsAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsRequest.Merge(m, src)
}
func (m *ListPaymentsAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Size(m)
}
func (m *ListPaymentsAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsRequest proto.InternalMessageInfo

func (m *ListPaymentsAccountsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for
// [PaymentsAccountService.ListPaymentsAccounts][google.ads.googleads.v1.services.PaymentsAccountService.ListPaymentsAccounts].
type ListPaymentsAccountsResponse struct {
	// The list of accessible Payments accounts.
	PaymentsAccounts     []*resources.PaymentsAccount `protobuf:"bytes,1,rep,name=payments_accounts,json=paymentsAccounts,proto3" json:"payments_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListPaymentsAccountsResponse) Reset()         { *m = ListPaymentsAccountsResponse{} }
func (m *ListPaymentsAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsResponse) ProtoMessage()    {}
func (*ListPaymentsAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{1}
}

func (m *ListPaymentsAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsResponse.Merge(m, src)
}
func (m *ListPaymentsAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Size(m)
}
func (m *ListPaymentsAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsResponse proto.InternalMessageInfo

func (m *ListPaymentsAccountsResponse) GetPaymentsAccounts() []*resources.PaymentsAccount {
	if m != nil {
		return m.PaymentsAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPaymentsAccountsRequest)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsRequest")
	proto.RegisterType((*ListPaymentsAccountsResponse)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/payments_account_service.proto", fileDescriptor_68acb38ba4f095b7)
}

var fileDescriptor_68acb38ba4f095b7 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0x0e, 0x0c, 0xa6, 0x5c, 0x36, 0x33, 0x46, 0x48, 0x02, 0x33, 0x21, 0x87, 0xb0, 0x83,
	0x34, 0x67, 0x30, 0x86, 0x46, 0x32, 0x9c, 0x4b, 0x36, 0xd8, 0x21, 0x64, 0x90, 0xc3, 0x30, 0x18,
	0xcf, 0x16, 0xc6, 0x90, 0x48, 0xae, 0x3f, 0x39, 0x50, 0x4a, 0x29, 0xe4, 0x15, 0xfa, 0x06, 0x3d,
	0xf6, 0x51, 0x0a, 0x3d, 0xf5, 0x15, 0x7a, 0x28, 0x7d, 0x8a, 0xe2, 0xc8, 0x32, 0xc5, 0x38, 0x0d,
	0xf4, 0xf6, 0x43, 0xfa, 0xfd, 0xd1, 0xf7, 0xd3, 0x87, 0x7e, 0xc6, 0x42, 0xc4, 0x6b, 0x46, 0x82,
	0x08, 0x88, 0x82, 0x05, 0xda, 0x3a, 0x04, 0x58, 0xb6, 0x4d, 0x42, 0x06, 0x24, 0x0d, 0x4e, 0x37,
	0x8c, 0x4b, 0xf0, 0x83, 0x30, 0x14, 0x39, 0x97, 0x7e, 0x79, 0x83, 0xd3, 0x4c, 0x48, 0x61, 0xd9,
	0x4a, 0x85, 0x83, 0x08, 0x70, 0x65, 0x80, 0xb7, 0x0e, 0xd6, 0x06, 0xdd, 0xef, 0x87, 0x22, 0x32,
	0x06, 0x22, 0xcf, 0x9a, 0x32, 0x94, 0x77, 0xb7, 0xaf, 0x95, 0x69, 0x42, 0x02, 0xce, 0x85, 0x0c,
	0x64, 0x22, 0x38, 0xa8, 0xdb, 0xc1, 0x14, 0xf5, 0xfe, 0x24, 0x20, 0x17, 0xa5, 0xd6, 0x55, 0x52,
	0x58, 0xb2, 0x93, 0x9c, 0x81, 0xb4, 0x3e, 0xa1, 0x76, 0x98, 0x83, 0x14, 0x1b, 0x96, 0xf9, 0x49,
	0xd4, 0x31, 0x6c, 0x63, 0xf4, 0x76, 0x89, 0xf4, 0xd1, 0xef, 0x68, 0x70, 0x81, 0xfa, 0xcd, 0x7a,
	0x48, 0x05, 0x07, 0x66, 0xf9, 0xe8, 0x7d, 0xfd, 0x5d, 0xd0, 0x31, 0xec, 0xd6, 0xa8, 0x3d, 0x1e,
	0xe3, 0x43, 0x53, 0x57, 0x33, 0xe1, 0x9a, 0xef, 0xf2, 0x5d, 0x5a, 0x0b, 0x1a, 0x3f, 0x18, 0xe8,
	0x63, 0x8d, 0xf5, 0x57, 0x95, 0x66, 0xdd, 0x1a, 0xe8, 0x43, 0xd3, 0xe3, 0xac, 0x09, 0x3e, 0xd6,
	0x37, 0x7e, 0xa1, 0x94, 0xee, 0xf4, 0xb5, 0x72, 0xd5, 0xc9, 0xe0, 0xdb, 0xee, 0xee, 0xfe, 0xd2,
	0xfc, 0x62, 0xe1, 0xe2, 0xff, 0x74, 0x97, 0x40, 0xce, 0x9e, 0x35, 0x3d, 0xf9, 0x7c, 0x4e, 0xea,
	0xa3, 0xce, 0x76, 0x26, 0x1a, 0x86, 0x62, 0x73, 0x34, 0x7d, 0xd6, 0x6b, 0x2e, 0x64, 0x51, 0xfc,
	0xf8, 0xc2, 0xf8, 0xf7, 0xab, 0x34, 0x88, 0xc5, 0x3a, 0xe0, 0x31, 0x16, 0x59, 0x4c, 0x62, 0xc6,
	0xf7, 0xfb, 0xa0, 0x77, 0x2b, 0x4d, 0xe0, 0xf0, 0x36, 0xff, 0xd0, 0xe0, 0xca, 0x6c, 0xcd, 0x5d,
	0xf7, 0xda, 0xb4, 0xe7, 0xca, 0xd0, 0x8d, 0x00, 0x2b, 0x58, 0xa0, 0x95, 0x83, 0xcb, 0x60, 0xb8,
	0xd1, 0x14, 0xcf, 0x8d, 0xc0, 0xab, 0x28, 0xde, 0xca, 0xf1, 0x34, 0xe5, 0xd1, 0x1c, 0xaa, 0x73,
	0x4a, 0xdd, 0x08, 0x28, 0xad, 0x48, 0x94, 0xae, 0x1c, 0x4a, 0x35, 0xed, 0xff, 0x9b, 0xfd, 0x3b,
	0xbf, 0x3e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xe9, 0x84, 0xbb, 0x74, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all Payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc *grpc.ClientConn
}

func NewPaymentsAccountServiceClient(cc *grpc.ClientConn) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
type PaymentsAccountServiceServer interface {
	// Returns all Payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
}

func RegisterPaymentsAccountServiceServer(s *grpc.Server, srv PaymentsAccountServiceServer) {
	s.RegisterService(&_PaymentsAccountService_serviceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentsAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/payments_account_service.proto",
}
