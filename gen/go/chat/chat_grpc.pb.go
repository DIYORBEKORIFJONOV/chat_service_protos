// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: chat/chat.proto

package chatgen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_CreateChat_FullMethodName          = "/ChatService/CreateChat"
	ChatService_SendMessage_FullMethodName         = "/ChatService/SendMessage"
	ChatService_GetChatHistory_FullMethodName      = "/ChatService/GetChatHistory"
	ChatService_GetChatParticipants_FullMethodName = "/ChatService/GetChatParticipants"
	ChatService_AddParticipant_FullMethodName      = "/ChatService/AddParticipant"
	ChatService_RemoveParticipant_FullMethodName   = "/ChatService/RemoveParticipant"
	ChatService_GetMessageStatus_FullMethodName    = "/ChatService/GetMessageStatus"
	ChatService_DeleteMessage_FullMethodName       = "/ChatService/DeleteMessage"
	ChatService_AddAdmin_FullMethodName            = "/ChatService/AddAdmin"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SendMessageRequest, MessageResponse], error)
	GetChatHistory(ctx context.Context, in *GetChatHistoryRequest, opts ...grpc.CallOption) (*GetChatHistoryResponse, error)
	GetChatParticipants(ctx context.Context, in *GetChatParticipantsRequest, opts ...grpc.CallOption) (*GetChatParticipantsResponse, error)
	AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*AddParticipantResponse, error)
	RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*RemoveParticipantResponse, error)
	GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SendMessageRequest, MessageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_SendMessage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SendMessageRequest, MessageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_SendMessageClient = grpc.BidiStreamingClient[SendMessageRequest, MessageResponse]

func (c *chatServiceClient) GetChatHistory(ctx context.Context, in *GetChatHistoryRequest, opts ...grpc.CallOption) (*GetChatHistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatHistoryResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChatHistory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetChatParticipants(ctx context.Context, in *GetChatParticipantsRequest, opts ...grpc.CallOption) (*GetChatParticipantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChatParticipantsResponse)
	err := c.cc.Invoke(ctx, ChatService_GetChatParticipants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddParticipant(ctx context.Context, in *AddParticipantRequest, opts ...grpc.CallOption) (*AddParticipantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddParticipantResponse)
	err := c.cc.Invoke(ctx, ChatService_AddParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RemoveParticipant(ctx context.Context, in *RemoveParticipantRequest, opts ...grpc.CallOption) (*RemoveParticipantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveParticipantResponse)
	err := c.cc.Invoke(ctx, ChatService_RemoveParticipant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetMessageStatus(ctx context.Context, in *GetMessageStatusRequest, opts ...grpc.CallOption) (*MessageStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessageStatusResponse)
	err := c.cc.Invoke(ctx, ChatService_GetMessageStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_DeleteMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddAdminResponse)
	err := c.cc.Invoke(ctx, ChatService_AddAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
type ChatServiceServer interface {
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	SendMessage(grpc.BidiStreamingServer[SendMessageRequest, MessageResponse]) error
	GetChatHistory(context.Context, *GetChatHistoryRequest) (*GetChatHistoryResponse, error)
	GetChatParticipants(context.Context, *GetChatParticipantsRequest) (*GetChatParticipantsResponse, error)
	AddParticipant(context.Context, *AddParticipantRequest) (*AddParticipantResponse, error)
	RemoveParticipant(context.Context, *RemoveParticipantRequest) (*RemoveParticipantResponse, error)
	GetMessageStatus(context.Context, *GetMessageStatusRequest) (*MessageStatusResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(grpc.BidiStreamingServer[SendMessageRequest, MessageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetChatHistory(context.Context, *GetChatHistoryRequest) (*GetChatHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatHistory not implemented")
}
func (UnimplementedChatServiceServer) GetChatParticipants(context.Context, *GetChatParticipantsRequest) (*GetChatParticipantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChatParticipants not implemented")
}
func (UnimplementedChatServiceServer) AddParticipant(context.Context, *AddParticipantRequest) (*AddParticipantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParticipant not implemented")
}
func (UnimplementedChatServiceServer) RemoveParticipant(context.Context, *RemoveParticipantRequest) (*RemoveParticipantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveParticipant not implemented")
}
func (UnimplementedChatServiceServer) GetMessageStatus(context.Context, *GetMessageStatusRequest) (*MessageStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageStatus not implemented")
}
func (UnimplementedChatServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedChatServiceServer) AddAdmin(context.Context, *AddAdminRequest) (*AddAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&grpc.GenericServerStream[SendMessageRequest, MessageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatService_SendMessageServer = grpc.BidiStreamingServer[SendMessageRequest, MessageResponse]

func _ChatService_GetChatHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatHistory(ctx, req.(*GetChatHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetChatParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChatParticipantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetChatParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetChatParticipants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetChatParticipants(ctx, req.(*GetChatParticipantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddParticipant(ctx, req.(*AddParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RemoveParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveParticipantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RemoveParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_RemoveParticipant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RemoveParticipant(ctx, req.(*RemoveParticipantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetMessageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetMessageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetMessageStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetMessageStatus(ctx, req.(*GetMessageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddAdmin(ctx, req.(*AddAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChat",
			Handler:    _ChatService_CreateChat_Handler,
		},
		{
			MethodName: "GetChatHistory",
			Handler:    _ChatService_GetChatHistory_Handler,
		},
		{
			MethodName: "GetChatParticipants",
			Handler:    _ChatService_GetChatParticipants_Handler,
		},
		{
			MethodName: "AddParticipant",
			Handler:    _ChatService_AddParticipant_Handler,
		},
		{
			MethodName: "RemoveParticipant",
			Handler:    _ChatService_RemoveParticipant_Handler,
		},
		{
			MethodName: "GetMessageStatus",
			Handler:    _ChatService_GetMessageStatus_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _ChatService_DeleteMessage_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _ChatService_AddAdmin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
