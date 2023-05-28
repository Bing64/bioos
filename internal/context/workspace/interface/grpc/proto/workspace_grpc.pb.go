// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: internal/context/workspace/interface/grpc/proto/workspace.proto

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
	WorkspaceService_GetWorkspace_FullMethodName    = "/proto.WorkspaceService/GetWorkspace"
	WorkspaceService_CreateWorkspace_FullMethodName = "/proto.WorkspaceService/CreateWorkspace"
	WorkspaceService_DeleteWorkspace_FullMethodName = "/proto.WorkspaceService/DeleteWorkspace"
	WorkspaceService_UpdateWorkspace_FullMethodName = "/proto.WorkspaceService/UpdateWorkspace"
	WorkspaceService_ListWorkspace_FullMethodName   = "/proto.WorkspaceService/ListWorkspace"
	WorkspaceService_ImportWorkspace_FullMethodName = "/proto.WorkspaceService/ImportWorkspace"
)

// WorkspaceServiceClient is the client API for WorkspaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspaceServiceClient interface {
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error)
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error)
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error)
	ListWorkspace(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error)
	ImportWorkspace(ctx context.Context, opts ...grpc.CallOption) (WorkspaceService_ImportWorkspaceClient, error)
}

type workspaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceServiceClient(cc grpc.ClientConnInterface) WorkspaceServiceClient {
	return &workspaceServiceClient{cc}
}

func (c *workspaceServiceClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error) {
	out := new(GetWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_GetWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error) {
	out := new(CreateWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_CreateWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_DeleteWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) UpdateWorkspace(ctx context.Context, in *UpdateWorkspaceRequest, opts ...grpc.CallOption) (*UpdateWorkspaceResponse, error) {
	out := new(UpdateWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_UpdateWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ListWorkspace(ctx context.Context, in *ListWorkspaceRequest, opts ...grpc.CallOption) (*ListWorkspaceResponse, error) {
	out := new(ListWorkspaceResponse)
	err := c.cc.Invoke(ctx, WorkspaceService_ListWorkspace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspaceServiceClient) ImportWorkspace(ctx context.Context, opts ...grpc.CallOption) (WorkspaceService_ImportWorkspaceClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkspaceService_ServiceDesc.Streams[0], WorkspaceService_ImportWorkspace_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &workspaceServiceImportWorkspaceClient{stream}
	return x, nil
}

type WorkspaceService_ImportWorkspaceClient interface {
	Send(*ImportWorkspaceRequest) error
	CloseAndRecv() (*ImportWorkspaceResponse, error)
	grpc.ClientStream
}

type workspaceServiceImportWorkspaceClient struct {
	grpc.ClientStream
}

func (x *workspaceServiceImportWorkspaceClient) Send(m *ImportWorkspaceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workspaceServiceImportWorkspaceClient) CloseAndRecv() (*ImportWorkspaceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImportWorkspaceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkspaceServiceServer is the server API for WorkspaceService service.
// All implementations must embed UnimplementedWorkspaceServiceServer
// for forward compatibility
type WorkspaceServiceServer interface {
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error)
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error)
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error)
	ListWorkspace(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error)
	ImportWorkspace(WorkspaceService_ImportWorkspaceServer) error
	mustEmbedUnimplementedWorkspaceServiceServer()
}

// UnimplementedWorkspaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspaceServiceServer struct {
}

func (UnimplementedWorkspaceServiceServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) UpdateWorkspace(context.Context, *UpdateWorkspaceRequest) (*UpdateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ListWorkspace(context.Context, *ListWorkspaceRequest) (*ListWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) ImportWorkspace(WorkspaceService_ImportWorkspaceServer) error {
	return status.Errorf(codes.Unimplemented, "method ImportWorkspace not implemented")
}
func (UnimplementedWorkspaceServiceServer) mustEmbedUnimplementedWorkspaceServiceServer() {}

// UnsafeWorkspaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspaceServiceServer will
// result in compilation errors.
type UnsafeWorkspaceServiceServer interface {
	mustEmbedUnimplementedWorkspaceServiceServer()
}

func RegisterWorkspaceServiceServer(s grpc.ServiceRegistrar, srv WorkspaceServiceServer) {
	s.RegisterService(&WorkspaceService_ServiceDesc, srv)
}

func _WorkspaceService_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_GetWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_CreateWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_DeleteWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_UpdateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_UpdateWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).UpdateWorkspace(ctx, req.(*UpdateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ListWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceServiceServer).ListWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkspaceService_ListWorkspace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceServiceServer).ListWorkspace(ctx, req.(*ListWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspaceService_ImportWorkspace_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkspaceServiceServer).ImportWorkspace(&workspaceServiceImportWorkspaceServer{stream})
}

type WorkspaceService_ImportWorkspaceServer interface {
	SendAndClose(*ImportWorkspaceResponse) error
	Recv() (*ImportWorkspaceRequest, error)
	grpc.ServerStream
}

type workspaceServiceImportWorkspaceServer struct {
	grpc.ServerStream
}

func (x *workspaceServiceImportWorkspaceServer) SendAndClose(m *ImportWorkspaceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workspaceServiceImportWorkspaceServer) Recv() (*ImportWorkspaceRequest, error) {
	m := new(ImportWorkspaceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkspaceService_ServiceDesc is the grpc.ServiceDesc for WorkspaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WorkspaceService",
	HandlerType: (*WorkspaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkspace",
			Handler:    _WorkspaceService_GetWorkspace_Handler,
		},
		{
			MethodName: "CreateWorkspace",
			Handler:    _WorkspaceService_CreateWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspaceService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "UpdateWorkspace",
			Handler:    _WorkspaceService_UpdateWorkspace_Handler,
		},
		{
			MethodName: "ListWorkspace",
			Handler:    _WorkspaceService_ListWorkspace_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImportWorkspace",
			Handler:       _WorkspaceService_ImportWorkspace_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/context/workspace/interface/grpc/proto/workspace.proto",
}

const (
	DataModelService_ListDataModels_FullMethodName         = "/proto.DataModelService/ListDataModels"
	DataModelService_GetDataModel_FullMethodName           = "/proto.DataModelService/GetDataModel"
	DataModelService_ListDataModelRows_FullMethodName      = "/proto.DataModelService/ListDataModelRows"
	DataModelService_PatchDataModel_FullMethodName         = "/proto.DataModelService/PatchDataModel"
	DataModelService_DeleteDataModel_FullMethodName        = "/proto.DataModelService/DeleteDataModel"
	DataModelService_ListAllDataModelRowIDs_FullMethodName = "/proto.DataModelService/ListAllDataModelRowIDs"
)

// DataModelServiceClient is the client API for DataModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataModelServiceClient interface {
	ListDataModels(ctx context.Context, in *ListDataModelsRequest, opts ...grpc.CallOption) (*ListDataModelsResponse, error)
	GetDataModel(ctx context.Context, in *GetDataModelRequest, opts ...grpc.CallOption) (*GetDataModelResponse, error)
	ListDataModelRows(ctx context.Context, in *ListDataModelRowsRequest, opts ...grpc.CallOption) (*ListDataModelRowsResponse, error)
	PatchDataModel(ctx context.Context, in *PatchDataModelRequest, opts ...grpc.CallOption) (*PatchDataModelResponse, error)
	DeleteDataModel(ctx context.Context, in *DeleteDataModelRequest, opts ...grpc.CallOption) (*DeleteDataModelResponse, error)
	ListAllDataModelRowIDs(ctx context.Context, in *ListAllDataModelRowIDsRequest, opts ...grpc.CallOption) (*ListAllDataModelRowIDsResponse, error)
}

type dataModelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataModelServiceClient(cc grpc.ClientConnInterface) DataModelServiceClient {
	return &dataModelServiceClient{cc}
}

func (c *dataModelServiceClient) ListDataModels(ctx context.Context, in *ListDataModelsRequest, opts ...grpc.CallOption) (*ListDataModelsResponse, error) {
	out := new(ListDataModelsResponse)
	err := c.cc.Invoke(ctx, DataModelService_ListDataModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataModelServiceClient) GetDataModel(ctx context.Context, in *GetDataModelRequest, opts ...grpc.CallOption) (*GetDataModelResponse, error) {
	out := new(GetDataModelResponse)
	err := c.cc.Invoke(ctx, DataModelService_GetDataModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataModelServiceClient) ListDataModelRows(ctx context.Context, in *ListDataModelRowsRequest, opts ...grpc.CallOption) (*ListDataModelRowsResponse, error) {
	out := new(ListDataModelRowsResponse)
	err := c.cc.Invoke(ctx, DataModelService_ListDataModelRows_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataModelServiceClient) PatchDataModel(ctx context.Context, in *PatchDataModelRequest, opts ...grpc.CallOption) (*PatchDataModelResponse, error) {
	out := new(PatchDataModelResponse)
	err := c.cc.Invoke(ctx, DataModelService_PatchDataModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataModelServiceClient) DeleteDataModel(ctx context.Context, in *DeleteDataModelRequest, opts ...grpc.CallOption) (*DeleteDataModelResponse, error) {
	out := new(DeleteDataModelResponse)
	err := c.cc.Invoke(ctx, DataModelService_DeleteDataModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataModelServiceClient) ListAllDataModelRowIDs(ctx context.Context, in *ListAllDataModelRowIDsRequest, opts ...grpc.CallOption) (*ListAllDataModelRowIDsResponse, error) {
	out := new(ListAllDataModelRowIDsResponse)
	err := c.cc.Invoke(ctx, DataModelService_ListAllDataModelRowIDs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataModelServiceServer is the server API for DataModelService service.
// All implementations must embed UnimplementedDataModelServiceServer
// for forward compatibility
type DataModelServiceServer interface {
	ListDataModels(context.Context, *ListDataModelsRequest) (*ListDataModelsResponse, error)
	GetDataModel(context.Context, *GetDataModelRequest) (*GetDataModelResponse, error)
	ListDataModelRows(context.Context, *ListDataModelRowsRequest) (*ListDataModelRowsResponse, error)
	PatchDataModel(context.Context, *PatchDataModelRequest) (*PatchDataModelResponse, error)
	DeleteDataModel(context.Context, *DeleteDataModelRequest) (*DeleteDataModelResponse, error)
	ListAllDataModelRowIDs(context.Context, *ListAllDataModelRowIDsRequest) (*ListAllDataModelRowIDsResponse, error)
	mustEmbedUnimplementedDataModelServiceServer()
}

// UnimplementedDataModelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataModelServiceServer struct {
}

func (UnimplementedDataModelServiceServer) ListDataModels(context.Context, *ListDataModelsRequest) (*ListDataModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataModels not implemented")
}
func (UnimplementedDataModelServiceServer) GetDataModel(context.Context, *GetDataModelRequest) (*GetDataModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataModel not implemented")
}
func (UnimplementedDataModelServiceServer) ListDataModelRows(context.Context, *ListDataModelRowsRequest) (*ListDataModelRowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataModelRows not implemented")
}
func (UnimplementedDataModelServiceServer) PatchDataModel(context.Context, *PatchDataModelRequest) (*PatchDataModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchDataModel not implemented")
}
func (UnimplementedDataModelServiceServer) DeleteDataModel(context.Context, *DeleteDataModelRequest) (*DeleteDataModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataModel not implemented")
}
func (UnimplementedDataModelServiceServer) ListAllDataModelRowIDs(context.Context, *ListAllDataModelRowIDsRequest) (*ListAllDataModelRowIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllDataModelRowIDs not implemented")
}
func (UnimplementedDataModelServiceServer) mustEmbedUnimplementedDataModelServiceServer() {}

// UnsafeDataModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataModelServiceServer will
// result in compilation errors.
type UnsafeDataModelServiceServer interface {
	mustEmbedUnimplementedDataModelServiceServer()
}

func RegisterDataModelServiceServer(s grpc.ServiceRegistrar, srv DataModelServiceServer) {
	s.RegisterService(&DataModelService_ServiceDesc, srv)
}

func _DataModelService_ListDataModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).ListDataModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_ListDataModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).ListDataModels(ctx, req.(*ListDataModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataModelService_GetDataModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).GetDataModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_GetDataModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).GetDataModel(ctx, req.(*GetDataModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataModelService_ListDataModelRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataModelRowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).ListDataModelRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_ListDataModelRows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).ListDataModelRows(ctx, req.(*ListDataModelRowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataModelService_PatchDataModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchDataModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).PatchDataModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_PatchDataModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).PatchDataModel(ctx, req.(*PatchDataModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataModelService_DeleteDataModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).DeleteDataModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_DeleteDataModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).DeleteDataModel(ctx, req.(*DeleteDataModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataModelService_ListAllDataModelRowIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllDataModelRowIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataModelServiceServer).ListAllDataModelRowIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataModelService_ListAllDataModelRowIDs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataModelServiceServer).ListAllDataModelRowIDs(ctx, req.(*ListAllDataModelRowIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataModelService_ServiceDesc is the grpc.ServiceDesc for DataModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DataModelService",
	HandlerType: (*DataModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataModels",
			Handler:    _DataModelService_ListDataModels_Handler,
		},
		{
			MethodName: "GetDataModel",
			Handler:    _DataModelService_GetDataModel_Handler,
		},
		{
			MethodName: "ListDataModelRows",
			Handler:    _DataModelService_ListDataModelRows_Handler,
		},
		{
			MethodName: "PatchDataModel",
			Handler:    _DataModelService_PatchDataModel_Handler,
		},
		{
			MethodName: "DeleteDataModel",
			Handler:    _DataModelService_DeleteDataModel_Handler,
		},
		{
			MethodName: "ListAllDataModelRowIDs",
			Handler:    _DataModelService_ListAllDataModelRowIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/context/workspace/interface/grpc/proto/workspace.proto",
}
