// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: internal/context/notebookserver/interface/grpc/proto/notebookserver.proto

package proto

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

const (
	NotebookServerService_CreateNotebookServer_FullMethodName         = "/proto.NotebookServerService/CreateNotebookServer"
	NotebookServerService_GetNotebookServer_FullMethodName            = "/proto.NotebookServerService/GetNotebookServer"
	NotebookServerService_UpdateNotebookServerSettings_FullMethodName = "/proto.NotebookServerService/UpdateNotebookServerSettings"
	NotebookServerService_DeleteNotebookServer_FullMethodName         = "/proto.NotebookServerService/DeleteNotebookServer"
	NotebookServerService_SwitchNotebookServer_FullMethodName         = "/proto.NotebookServerService/SwitchNotebookServer"
	NotebookServerService_ListNotebookServers_FullMethodName          = "/proto.NotebookServerService/ListNotebookServers"
)

// NotebookServerServiceClient is the client API for NotebookServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotebookServerServiceClient interface {
	CreateNotebookServer(ctx context.Context, in *CreateNotebookServerRequest, opts ...grpc.CallOption) (*CreateNotebookServerResponse, error)
	GetNotebookServer(ctx context.Context, in *GetNotebookServerRequest, opts ...grpc.CallOption) (*GetNotebookServerResponse, error)
	UpdateNotebookServerSettings(ctx context.Context, in *UpdateNotebookServerSettingsRequest, opts ...grpc.CallOption) (*UpdateNotebookServerSettingsResponse, error)
	DeleteNotebookServer(ctx context.Context, in *DeleteNotebookServerRequest, opts ...grpc.CallOption) (*DeleteNotebookServerResponse, error)
	SwitchNotebookServer(ctx context.Context, in *SwitchNotebookServerRequest, opts ...grpc.CallOption) (*SwitchNotebookServerResponse, error)
	ListNotebookServers(ctx context.Context, in *ListNotebookServersRequest, opts ...grpc.CallOption) (*ListNotebookServersResponse, error)
}

type notebookServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotebookServerServiceClient(cc grpc.ClientConnInterface) NotebookServerServiceClient {
	return &notebookServerServiceClient{cc}
}

func (c *notebookServerServiceClient) CreateNotebookServer(ctx context.Context, in *CreateNotebookServerRequest, opts ...grpc.CallOption) (*CreateNotebookServerResponse, error) {
	out := new(CreateNotebookServerResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_CreateNotebookServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookServerServiceClient) GetNotebookServer(ctx context.Context, in *GetNotebookServerRequest, opts ...grpc.CallOption) (*GetNotebookServerResponse, error) {
	out := new(GetNotebookServerResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_GetNotebookServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookServerServiceClient) UpdateNotebookServerSettings(ctx context.Context, in *UpdateNotebookServerSettingsRequest, opts ...grpc.CallOption) (*UpdateNotebookServerSettingsResponse, error) {
	out := new(UpdateNotebookServerSettingsResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_UpdateNotebookServerSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookServerServiceClient) DeleteNotebookServer(ctx context.Context, in *DeleteNotebookServerRequest, opts ...grpc.CallOption) (*DeleteNotebookServerResponse, error) {
	out := new(DeleteNotebookServerResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_DeleteNotebookServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookServerServiceClient) SwitchNotebookServer(ctx context.Context, in *SwitchNotebookServerRequest, opts ...grpc.CallOption) (*SwitchNotebookServerResponse, error) {
	out := new(SwitchNotebookServerResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_SwitchNotebookServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notebookServerServiceClient) ListNotebookServers(ctx context.Context, in *ListNotebookServersRequest, opts ...grpc.CallOption) (*ListNotebookServersResponse, error) {
	out := new(ListNotebookServersResponse)
	err := c.cc.Invoke(ctx, NotebookServerService_ListNotebookServers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotebookServerServiceServer is the server API for NotebookServerService service.
// All implementations must embed UnimplementedNotebookServerServiceServer
// for forward compatibility
type NotebookServerServiceServer interface {
	CreateNotebookServer(context.Context, *CreateNotebookServerRequest) (*CreateNotebookServerResponse, error)
	GetNotebookServer(context.Context, *GetNotebookServerRequest) (*GetNotebookServerResponse, error)
	UpdateNotebookServerSettings(context.Context, *UpdateNotebookServerSettingsRequest) (*UpdateNotebookServerSettingsResponse, error)
	DeleteNotebookServer(context.Context, *DeleteNotebookServerRequest) (*DeleteNotebookServerResponse, error)
	SwitchNotebookServer(context.Context, *SwitchNotebookServerRequest) (*SwitchNotebookServerResponse, error)
	ListNotebookServers(context.Context, *ListNotebookServersRequest) (*ListNotebookServersResponse, error)
	mustEmbedUnimplementedNotebookServerServiceServer()
}

// UnimplementedNotebookServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotebookServerServiceServer struct {
}

func (UnimplementedNotebookServerServiceServer) CreateNotebookServer(context.Context, *CreateNotebookServerRequest) (*CreateNotebookServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotebookServer not implemented")
}
func (UnimplementedNotebookServerServiceServer) GetNotebookServer(context.Context, *GetNotebookServerRequest) (*GetNotebookServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotebookServer not implemented")
}
func (UnimplementedNotebookServerServiceServer) UpdateNotebookServerSettings(context.Context, *UpdateNotebookServerSettingsRequest) (*UpdateNotebookServerSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotebookServerSettings not implemented")
}
func (UnimplementedNotebookServerServiceServer) DeleteNotebookServer(context.Context, *DeleteNotebookServerRequest) (*DeleteNotebookServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotebookServer not implemented")
}
func (UnimplementedNotebookServerServiceServer) SwitchNotebookServer(context.Context, *SwitchNotebookServerRequest) (*SwitchNotebookServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwitchNotebookServer not implemented")
}
func (UnimplementedNotebookServerServiceServer) ListNotebookServers(context.Context, *ListNotebookServersRequest) (*ListNotebookServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotebookServers not implemented")
}
func (UnimplementedNotebookServerServiceServer) mustEmbedUnimplementedNotebookServerServiceServer() {}

// UnsafeNotebookServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotebookServerServiceServer will
// result in compilation errors.
type UnsafeNotebookServerServiceServer interface {
	mustEmbedUnimplementedNotebookServerServiceServer()
}

func RegisterNotebookServerServiceServer(s grpc.ServiceRegistrar, srv NotebookServerServiceServer) {
	s.RegisterService(&NotebookServerService_ServiceDesc, srv)
}

func _NotebookServerService_CreateNotebookServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotebookServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).CreateNotebookServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_CreateNotebookServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).CreateNotebookServer(ctx, req.(*CreateNotebookServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotebookServerService_GetNotebookServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotebookServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).GetNotebookServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_GetNotebookServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).GetNotebookServer(ctx, req.(*GetNotebookServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotebookServerService_UpdateNotebookServerSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotebookServerSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).UpdateNotebookServerSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_UpdateNotebookServerSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).UpdateNotebookServerSettings(ctx, req.(*UpdateNotebookServerSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotebookServerService_DeleteNotebookServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotebookServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).DeleteNotebookServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_DeleteNotebookServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).DeleteNotebookServer(ctx, req.(*DeleteNotebookServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotebookServerService_SwitchNotebookServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwitchNotebookServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).SwitchNotebookServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_SwitchNotebookServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).SwitchNotebookServer(ctx, req.(*SwitchNotebookServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotebookServerService_ListNotebookServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotebookServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotebookServerServiceServer).ListNotebookServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotebookServerService_ListNotebookServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotebookServerServiceServer).ListNotebookServers(ctx, req.(*ListNotebookServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotebookServerService_ServiceDesc is the grpc.ServiceDesc for NotebookServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotebookServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NotebookServerService",
	HandlerType: (*NotebookServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotebookServer",
			Handler:    _NotebookServerService_CreateNotebookServer_Handler,
		},
		{
			MethodName: "GetNotebookServer",
			Handler:    _NotebookServerService_GetNotebookServer_Handler,
		},
		{
			MethodName: "UpdateNotebookServerSettings",
			Handler:    _NotebookServerService_UpdateNotebookServerSettings_Handler,
		},
		{
			MethodName: "DeleteNotebookServer",
			Handler:    _NotebookServerService_DeleteNotebookServer_Handler,
		},
		{
			MethodName: "SwitchNotebookServer",
			Handler:    _NotebookServerService_SwitchNotebookServer_Handler,
		},
		{
			MethodName: "ListNotebookServers",
			Handler:    _NotebookServerService_ListNotebookServers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/context/notebookserver/interface/grpc/proto/notebookserver.proto",
}
