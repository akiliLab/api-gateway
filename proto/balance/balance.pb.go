// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/balance/balance.proto

package balance

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Request struct {
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type Response struct {
	Completed            bool     `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_be989d5aed6f6a6c, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Response) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Response) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "balance.Request")
	proto.RegisterType((*Response)(nil), "balance.Response")
}

func init() { proto.RegisterFile("proto/balance/balance.proto", fileDescriptor_be989d5aed6f6a6c) }

var fileDescriptor_be989d5aed6f6a6c = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4a, 0xcc, 0x49, 0xcc, 0x4b, 0x4e, 0x85, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0x76,
	0x28, 0x57, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31,
	0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18, 0xa2, 0x4c, 0x49, 0x83, 0x8b, 0x3d,
	0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x96, 0x8b, 0x2b, 0x31, 0x39, 0x39, 0xbf, 0x34,
	0xaf, 0x24, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x13, 0x2a, 0xe2, 0x99,
	0xa2, 0x14, 0xc3, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc3, 0xc5,
	0x99, 0x9c, 0x9f, 0x5b, 0x90, 0x93, 0x5a, 0x92, 0x0a, 0x51, 0xc9, 0x11, 0x84, 0x10, 0x10, 0x12,
	0xe3, 0x62, 0x4b, 0xcc, 0x05, 0xe9, 0x92, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0xf2, 0x84,
	0xa4, 0xb8, 0x38, 0x92, 0x4b, 0x8b, 0x8a, 0x52, 0xf3, 0x92, 0x2b, 0x25, 0x98, 0xc1, 0xc6, 0xc3,
	0xf9, 0x46, 0xa1, 0x5c, 0xec, 0x4e, 0x10, 0x07, 0x0b, 0x79, 0x71, 0xb1, 0x3a, 0x67, 0xa4, 0x26,
	0x67, 0x0b, 0x09, 0xe8, 0xc1, 0xbc, 0x04, 0x75, 0xa2, 0x94, 0x20, 0x92, 0x08, 0xc4, 0x29, 0x4a,
	0x32, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x53, 0x12, 0xd4, 0x2f, 0x33, 0x84, 0x07, 0x45, 0x32,
	0x48, 0xbf, 0x15, 0xa3, 0x56, 0x12, 0x1b, 0xd8, 0x97, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x94, 0x1b, 0xc0, 0xc8, 0x2b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BalanceClient is the client API for Balance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BalanceClient interface {
	Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type balanceClient struct {
	cc *grpc.ClientConn
}

func NewBalanceClient(cc *grpc.ClientConn) BalanceClient {
	return &balanceClient{cc}
}

func (c *balanceClient) Check(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/balance.Balance/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServer is the server API for Balance service.
type BalanceServer interface {
	Check(context.Context, *Request) (*Response, error)
}

// UnimplementedBalanceServer can be embedded to have forward compatible implementations.
type UnimplementedBalanceServer struct {
}

func (*UnimplementedBalanceServer) Check(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterBalanceServer(s *grpc.Server, srv BalanceServer) {
	s.RegisterService(&_Balance_serviceDesc, srv)
}

func _Balance_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/balance.Balance/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServer).Check(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Balance_serviceDesc = grpc.ServiceDesc{
	ServiceName: "balance.Balance",
	HandlerType: (*BalanceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Balance_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/balance/balance.proto",
}