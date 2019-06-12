// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
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

func (m *Request) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type Response struct {
	Output               string   `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
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

func (m *Response) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "chat.Request")
	proto.RegisterType((*Response)(nil), "chat.Response")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x0a, 0x4a, 0x4b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x25, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xfc, 0xd2, 0x12, 0x84, 0x12, 0x28, 0xcf, 0xc8, 0x8c,
	0x8b, 0xc5, 0x39, 0x23, 0xb1, 0x44, 0x48, 0x8f, 0x8b, 0xd3, 0x29, 0x33, 0x25, 0xb8, 0xa4, 0x28,
	0x35, 0x31, 0x57, 0x88, 0x57, 0x0f, 0x6c, 0x19, 0xd4, 0x74, 0x29, 0x3e, 0x18, 0x17, 0x62, 0x96,
	0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0x25, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0xb0, 0xb2, 0xa9, 0x97, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	BidStream(ctx context.Context, opts ...grpc.CallOption) (Chat_BidStreamClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) BidStream(ctx context.Context, opts ...grpc.CallOption) (Chat_BidStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/BidStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBidStreamClient{stream}
	return x, nil
}

type Chat_BidStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type chatBidStreamClient struct {
	grpc.ClientStream
}

func (x *chatBidStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatBidStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	BidStream(Chat_BidStreamServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) BidStream(srv Chat_BidStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidStream not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_BidStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).BidStream(&chatBidStreamServer{stream})
}

type Chat_BidStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type chatBidStreamServer struct {
	grpc.ServerStream
}

func (x *chatBidStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatBidStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidStream",
			Handler:       _Chat_BidStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}