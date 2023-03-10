// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: contracts/awesomechat.proto

package awesomechat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatCoreClient is the client API for ChatCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatCoreClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatCore_ConnectClient, error)
	SendMessage(ctx context.Context, in *SentMessage, opts ...grpc.CallOption) (*Empty, error)
	CreateGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error)
	JoinGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error)
	LeftGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error)
	ListChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListChannelsResponse, error)
}

type chatCoreClient struct {
	cc grpc.ClientConnInterface
}

func NewChatCoreClient(cc grpc.ClientConnInterface) ChatCoreClient {
	return &chatCoreClient{cc}
}

func (c *chatCoreClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/ChatCore/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCoreClient) Connect(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatCore_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatCore_ServiceDesc.Streams[0], "/ChatCore/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatCoreConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatCore_ConnectClient interface {
	Recv() (*ReceivedMessage, error)
	grpc.ClientStream
}

type chatCoreConnectClient struct {
	grpc.ClientStream
}

func (x *chatCoreConnectClient) Recv() (*ReceivedMessage, error) {
	m := new(ReceivedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatCoreClient) SendMessage(ctx context.Context, in *SentMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatCore/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCoreClient) CreateGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatCore/CreateGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCoreClient) JoinGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatCore/JoinGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCoreClient) LeftGroupChat(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/ChatCore/LeftGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatCoreClient) ListChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, "/ChatCore/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatCoreServer is the server API for ChatCore service.
// All implementations must embed UnimplementedChatCoreServer
// for forward compatibility
type ChatCoreServer interface {
	CreateUser(context.Context, *User) (*CreateUserResponse, error)
	Connect(*User, ChatCore_ConnectServer) error
	SendMessage(context.Context, *SentMessage) (*Empty, error)
	CreateGroupChat(context.Context, *Group) (*Empty, error)
	JoinGroupChat(context.Context, *Group) (*Empty, error)
	LeftGroupChat(context.Context, *Group) (*Empty, error)
	ListChannels(context.Context, *Empty) (*ListChannelsResponse, error)
	mustEmbedUnimplementedChatCoreServer()
}

// UnimplementedChatCoreServer must be embedded to have forward compatible implementations.
type UnimplementedChatCoreServer struct {
}

func (UnimplementedChatCoreServer) CreateUser(context.Context, *User) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatCoreServer) Connect(*User, ChatCore_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatCoreServer) SendMessage(context.Context, *SentMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatCoreServer) CreateGroupChat(context.Context, *Group) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChat not implemented")
}
func (UnimplementedChatCoreServer) JoinGroupChat(context.Context, *Group) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroupChat not implemented")
}
func (UnimplementedChatCoreServer) LeftGroupChat(context.Context, *Group) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeftGroupChat not implemented")
}
func (UnimplementedChatCoreServer) ListChannels(context.Context, *Empty) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (UnimplementedChatCoreServer) mustEmbedUnimplementedChatCoreServer() {}

// UnsafeChatCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatCoreServer will
// result in compilation errors.
type UnsafeChatCoreServer interface {
	mustEmbedUnimplementedChatCoreServer()
}

func RegisterChatCoreServer(s grpc.ServiceRegistrar, srv ChatCoreServer) {
	s.RegisterService(&ChatCore_ServiceDesc, srv)
}

func _ChatCore_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCore_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatCoreServer).Connect(m, &chatCoreConnectServer{stream})
}

type ChatCore_ConnectServer interface {
	Send(*ReceivedMessage) error
	grpc.ServerStream
}

type chatCoreConnectServer struct {
	grpc.ServerStream
}

func (x *chatCoreConnectServer) Send(m *ReceivedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatCore_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).SendMessage(ctx, req.(*SentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCore_CreateGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).CreateGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/CreateGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).CreateGroupChat(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCore_JoinGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).JoinGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/JoinGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).JoinGroupChat(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCore_LeftGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).LeftGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/LeftGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).LeftGroupChat(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatCore_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatCoreServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatCore/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatCoreServer).ListChannels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatCore_ServiceDesc is the grpc.ServiceDesc for ChatCore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatCore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatCore",
	HandlerType: (*ChatCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ChatCore_CreateUser_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatCore_SendMessage_Handler,
		},
		{
			MethodName: "CreateGroupChat",
			Handler:    _ChatCore_CreateGroupChat_Handler,
		},
		{
			MethodName: "JoinGroupChat",
			Handler:    _ChatCore_JoinGroupChat_Handler,
		},
		{
			MethodName: "LeftGroupChat",
			Handler:    _ChatCore_LeftGroupChat_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _ChatCore_ListChannels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatCore_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "contracts/awesomechat.proto",
}
