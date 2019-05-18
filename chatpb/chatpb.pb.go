// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatpb.proto

package chatpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{0}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type ChatDetails struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatDetails) Reset()         { *m = ChatDetails{} }
func (m *ChatDetails) String() string { return proto.CompactTextString(m) }
func (*ChatDetails) ProtoMessage()    {}
func (*ChatDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{1}
}
func (m *ChatDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatDetails.Unmarshal(m, b)
}
func (m *ChatDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatDetails.Marshal(b, m, deterministic)
}
func (dst *ChatDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatDetails.Merge(dst, src)
}
func (m *ChatDetails) XXX_Size() int {
	return xxx_messageInfo_ChatDetails.Size(m)
}
func (m *ChatDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ChatDetails proto.InternalMessageInfo

func (m *ChatDetails) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ChatDetails) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChatDetails) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMessagesRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessagesRequest) Reset()         { *m = SendMessagesRequest{} }
func (m *SendMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessagesRequest) ProtoMessage()    {}
func (*SendMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{2}
}
func (m *SendMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessagesRequest.Unmarshal(m, b)
}
func (m *SendMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessagesRequest.Marshal(b, m, deterministic)
}
func (dst *SendMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessagesRequest.Merge(dst, src)
}
func (m *SendMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessagesRequest.Size(m)
}
func (m *SendMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessagesRequest proto.InternalMessageInfo

func (m *SendMessagesRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *SendMessagesRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SendMessagesRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RcvMessagesRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcvMessagesRequest) Reset()         { *m = RcvMessagesRequest{} }
func (m *RcvMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*RcvMessagesRequest) ProtoMessage()    {}
func (*RcvMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{3}
}
func (m *RcvMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcvMessagesRequest.Unmarshal(m, b)
}
func (m *RcvMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcvMessagesRequest.Marshal(b, m, deterministic)
}
func (dst *RcvMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcvMessagesRequest.Merge(dst, src)
}
func (m *RcvMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_RcvMessagesRequest.Size(m)
}
func (m *RcvMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RcvMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RcvMessagesRequest proto.InternalMessageInfo

func (m *RcvMessagesRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

type RcvMessagesResponse struct {
	Usermane             string   `protobuf:"bytes,1,opt,name=usermane,proto3" json:"usermane,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcvMessagesResponse) Reset()         { *m = RcvMessagesResponse{} }
func (m *RcvMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*RcvMessagesResponse) ProtoMessage()    {}
func (*RcvMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{4}
}
func (m *RcvMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcvMessagesResponse.Unmarshal(m, b)
}
func (m *RcvMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcvMessagesResponse.Marshal(b, m, deterministic)
}
func (dst *RcvMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcvMessagesResponse.Merge(dst, src)
}
func (m *RcvMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_RcvMessagesResponse.Size(m)
}
func (m *RcvMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RcvMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RcvMessagesResponse proto.InternalMessageInfo

func (m *RcvMessagesResponse) GetUsermane() string {
	if m != nil {
		return m.Usermane
	}
	return ""
}

func (m *RcvMessagesResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type NewChatRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewChatRequest) Reset()         { *m = NewChatRequest{} }
func (m *NewChatRequest) String() string { return proto.CompactTextString(m) }
func (*NewChatRequest) ProtoMessage()    {}
func (*NewChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{5}
}
func (m *NewChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewChatRequest.Unmarshal(m, b)
}
func (m *NewChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewChatRequest.Marshal(b, m, deterministic)
}
func (dst *NewChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewChatRequest.Merge(dst, src)
}
func (m *NewChatRequest) XXX_Size() int {
	return xxx_messageInfo_NewChatRequest.Size(m)
}
func (m *NewChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewChatRequest proto.InternalMessageInfo

func (m *NewChatRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

type CloseChatRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseChatRequest) Reset()         { *m = CloseChatRequest{} }
func (m *CloseChatRequest) String() string { return proto.CompactTextString(m) }
func (*CloseChatRequest) ProtoMessage()    {}
func (*CloseChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chatpb_4e1e2c630afa795e, []int{6}
}
func (m *CloseChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseChatRequest.Unmarshal(m, b)
}
func (m *CloseChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseChatRequest.Marshal(b, m, deterministic)
}
func (dst *CloseChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseChatRequest.Merge(dst, src)
}
func (m *CloseChatRequest) XXX_Size() int {
	return xxx_messageInfo_CloseChatRequest.Size(m)
}
func (m *CloseChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseChatRequest proto.InternalMessageInfo

func (m *CloseChatRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func init() {
	proto.RegisterType((*Void)(nil), "chatpb.Void")
	proto.RegisterType((*ChatDetails)(nil), "chatpb.ChatDetails")
	proto.RegisterType((*SendMessagesRequest)(nil), "chatpb.SendMessagesRequest")
	proto.RegisterType((*RcvMessagesRequest)(nil), "chatpb.RcvMessagesRequest")
	proto.RegisterType((*RcvMessagesResponse)(nil), "chatpb.RcvMessagesResponse")
	proto.RegisterType((*NewChatRequest)(nil), "chatpb.NewChatRequest")
	proto.RegisterType((*CloseChatRequest)(nil), "chatpb.CloseChatRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamChatClient is the client API for StreamChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamChatClient interface {
	NewChat(ctx context.Context, in *NewChatRequest, opts ...grpc.CallOption) (*Void, error)
	SendMessages(ctx context.Context, in *SendMessagesRequest, opts ...grpc.CallOption) (*Void, error)
	RcvMessages(ctx context.Context, in *RcvMessagesRequest, opts ...grpc.CallOption) (StreamChat_RcvMessagesClient, error)
	CloseChat(ctx context.Context, in *CloseChatRequest, opts ...grpc.CallOption) (*Void, error)
}

type streamChatClient struct {
	cc *grpc.ClientConn
}

func NewStreamChatClient(cc *grpc.ClientConn) StreamChatClient {
	return &streamChatClient{cc}
}

func (c *streamChatClient) NewChat(ctx context.Context, in *NewChatRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/chatpb.StreamChat/NewChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamChatClient) SendMessages(ctx context.Context, in *SendMessagesRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/chatpb.StreamChat/SendMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamChatClient) RcvMessages(ctx context.Context, in *RcvMessagesRequest, opts ...grpc.CallOption) (StreamChat_RcvMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamChat_serviceDesc.Streams[0], "/chatpb.StreamChat/RcvMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamChatRcvMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamChat_RcvMessagesClient interface {
	Recv() (*RcvMessagesResponse, error)
	grpc.ClientStream
}

type streamChatRcvMessagesClient struct {
	grpc.ClientStream
}

func (x *streamChatRcvMessagesClient) Recv() (*RcvMessagesResponse, error) {
	m := new(RcvMessagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamChatClient) CloseChat(ctx context.Context, in *CloseChatRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/chatpb.StreamChat/CloseChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamChatServer is the server API for StreamChat service.
type StreamChatServer interface {
	NewChat(context.Context, *NewChatRequest) (*Void, error)
	SendMessages(context.Context, *SendMessagesRequest) (*Void, error)
	RcvMessages(*RcvMessagesRequest, StreamChat_RcvMessagesServer) error
	CloseChat(context.Context, *CloseChatRequest) (*Void, error)
}

func RegisterStreamChatServer(s *grpc.Server, srv StreamChatServer) {
	s.RegisterService(&_StreamChat_serviceDesc, srv)
}

func _StreamChat_NewChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamChatServer).NewChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.StreamChat/NewChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamChatServer).NewChat(ctx, req.(*NewChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamChat_SendMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamChatServer).SendMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.StreamChat/SendMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamChatServer).SendMessages(ctx, req.(*SendMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamChat_RcvMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RcvMessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamChatServer).RcvMessages(m, &streamChatRcvMessagesServer{stream})
}

type StreamChat_RcvMessagesServer interface {
	Send(*RcvMessagesResponse) error
	grpc.ServerStream
}

type streamChatRcvMessagesServer struct {
	grpc.ServerStream
}

func (x *streamChatRcvMessagesServer) Send(m *RcvMessagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamChat_CloseChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamChatServer).CloseChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.StreamChat/CloseChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamChatServer).CloseChat(ctx, req.(*CloseChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamChat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatpb.StreamChat",
	HandlerType: (*StreamChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewChat",
			Handler:    _StreamChat_NewChat_Handler,
		},
		{
			MethodName: "SendMessages",
			Handler:    _StreamChat_SendMessages_Handler,
		},
		{
			MethodName: "CloseChat",
			Handler:    _StreamChat_CloseChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RcvMessages",
			Handler:       _StreamChat_RcvMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chatpb.proto",
}

func init() { proto.RegisterFile("chatpb.proto", fileDescriptor_chatpb_4e1e2c630afa795e) }

var fileDescriptor_chatpb_4e1e2c630afa795e = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x2a, 0x4d, 0x3b, 0x0d, 0x22, 0x53, 0xd0, 0x90, 0x5e, 0x64, 0x4f, 0x8a, 0x58,
	0xfc, 0x83, 0x27, 0x6f, 0xd6, 0x8b, 0x8a, 0x1e, 0x52, 0xf0, 0x20, 0x82, 0x6c, 0xbb, 0x83, 0x0d,
	0x34, 0xd9, 0x98, 0xdd, 0xea, 0xc7, 0x57, 0xb2, 0xc9, 0x86, 0xa4, 0x2d, 0xd8, 0x83, 0xc7, 0x97,
	0xf7, 0x86, 0x5f, 0xf6, 0xcd, 0x80, 0x37, 0x9b, 0x73, 0x9d, 0x4e, 0x47, 0x69, 0x26, 0xb5, 0xc4,
	0x4e, 0xa1, 0x58, 0x07, 0x76, 0x5f, 0x64, 0x24, 0xd8, 0x1b, 0xf4, 0xc7, 0x73, 0xae, 0xef, 0x48,
	0xf3, 0x68, 0xa1, 0xf0, 0x10, 0xdc, 0x3c, 0xf0, 0x1e, 0x09, 0xdf, 0x39, 0x72, 0x8e, 0x7b, 0xa1,
	0xc9, 0xdf, 0x0b, 0x0c, 0xa0, 0xbb, 0x54, 0x94, 0x25, 0x3c, 0x26, 0xbf, 0x6d, 0x9c, 0x4a, 0xa3,
	0x0f, 0x6e, 0x4c, 0x4a, 0xf1, 0x0f, 0xf2, 0x77, 0x8c, 0x65, 0x25, 0x13, 0x30, 0x98, 0x50, 0x22,
	0x9e, 0x0a, 0xa9, 0x42, 0xfa, 0x5c, 0x92, 0xd2, 0xff, 0x4d, 0x39, 0x03, 0x0c, 0x67, 0x5f, 0xdb,
	0x42, 0xd8, 0x23, 0x0c, 0x1a, 0x71, 0x95, 0xca, 0x44, 0x91, 0x65, 0xc7, 0x3c, 0xa1, 0x72, 0xa0,
	0xd2, 0x75, 0x76, 0xbb, 0xc9, 0x3e, 0x81, 0xbd, 0x67, 0xfa, 0xce, 0x2b, 0xfc, 0x93, 0x7b, 0x0a,
	0xfb, 0xe3, 0x85, 0x54, 0xb4, 0x4d, 0xf8, 0xf2, 0xc7, 0x01, 0x98, 0xe8, 0x8c, 0x78, 0x9c, 0xc7,
	0xf1, 0x02, 0xdc, 0x12, 0x83, 0x07, 0xa3, 0x72, 0xa1, 0x4d, 0x6e, 0xe0, 0xd9, 0xef, 0x66, 0xaf,
	0x2d, 0xbc, 0x01, 0xaf, 0xde, 0x3d, 0x0e, 0xad, 0xbf, 0x61, 0x23, 0x6b, 0xc3, 0x0f, 0xd0, 0xaf,
	0x75, 0x84, 0x81, 0xb5, 0xd7, 0x7b, 0x0e, 0x86, 0x1b, 0xbd, 0xa2, 0x54, 0xd6, 0x3a, 0x77, 0xf0,
	0x1a, 0x7a, 0xd5, 0xbb, 0xd1, 0xb7, 0xe9, 0xd5, 0x2a, 0x56, 0x7f, 0xe1, 0xb6, 0xfb, 0x5a, 0xde,
	0xea, 0xb4, 0x63, 0x4e, 0xf7, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x88, 0x0a, 0x09, 0xa4, 0xca,
	0x02, 0x00, 0x00,
}
