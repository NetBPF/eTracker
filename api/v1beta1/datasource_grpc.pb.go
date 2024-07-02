// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: api/v1beta1/datasource.proto

package v1beta1

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

// DataSourceServiceClient is the client API for DataSourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataSourceServiceClient interface {
	Write(ctx context.Context, in *WriteDataSourceRequest, opts ...grpc.CallOption) (*WriteDataSourceResponse, error)
	WriteStream(ctx context.Context, opts ...grpc.CallOption) (DataSourceService_WriteStreamClient, error)
}

type dataSourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataSourceServiceClient(cc grpc.ClientConnInterface) DataSourceServiceClient {
	return &dataSourceServiceClient{cc}
}

func (c *dataSourceServiceClient) Write(ctx context.Context, in *WriteDataSourceRequest, opts ...grpc.CallOption) (*WriteDataSourceResponse, error) {
	out := new(WriteDataSourceResponse)
	err := c.cc.Invoke(ctx, "/tracker.v1beta1.DataSourceService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSourceServiceClient) WriteStream(ctx context.Context, opts ...grpc.CallOption) (DataSourceService_WriteStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataSourceService_ServiceDesc.Streams[0], "/tracker.v1beta1.DataSourceService/WriteStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataSourceServiceWriteStreamClient{stream}
	return x, nil
}

type DataSourceService_WriteStreamClient interface {
	Send(*WriteDataSourceRequest) error
	CloseAndRecv() (*WriteDataSourceResponse, error)
	grpc.ClientStream
}

type dataSourceServiceWriteStreamClient struct {
	grpc.ClientStream
}

func (x *dataSourceServiceWriteStreamClient) Send(m *WriteDataSourceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataSourceServiceWriteStreamClient) CloseAndRecv() (*WriteDataSourceResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteDataSourceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataSourceServiceServer is the server API for DataSourceService service.
// All implementations must embed UnimplementedDataSourceServiceServer
// for forward compatibility
type DataSourceServiceServer interface {
	Write(context.Context, *WriteDataSourceRequest) (*WriteDataSourceResponse, error)
	WriteStream(DataSourceService_WriteStreamServer) error
	mustEmbedUnimplementedDataSourceServiceServer()
}

// UnimplementedDataSourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataSourceServiceServer struct {
}

func (UnimplementedDataSourceServiceServer) Write(context.Context, *WriteDataSourceRequest) (*WriteDataSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedDataSourceServiceServer) WriteStream(DataSourceService_WriteStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteStream not implemented")
}
func (UnimplementedDataSourceServiceServer) mustEmbedUnimplementedDataSourceServiceServer() {}

// UnsafeDataSourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataSourceServiceServer will
// result in compilation errors.
type UnsafeDataSourceServiceServer interface {
	mustEmbedUnimplementedDataSourceServiceServer()
}

func RegisterDataSourceServiceServer(s grpc.ServiceRegistrar, srv DataSourceServiceServer) {
	s.RegisterService(&DataSourceService_ServiceDesc, srv)
}

func _DataSourceService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteDataSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSourceServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracker.v1beta1.DataSourceService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSourceServiceServer).Write(ctx, req.(*WriteDataSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSourceService_WriteStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataSourceServiceServer).WriteStream(&dataSourceServiceWriteStreamServer{stream})
}

type DataSourceService_WriteStreamServer interface {
	SendAndClose(*WriteDataSourceResponse) error
	Recv() (*WriteDataSourceRequest, error)
	grpc.ServerStream
}

type dataSourceServiceWriteStreamServer struct {
	grpc.ServerStream
}

func (x *dataSourceServiceWriteStreamServer) SendAndClose(m *WriteDataSourceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataSourceServiceWriteStreamServer) Recv() (*WriteDataSourceRequest, error) {
	m := new(WriteDataSourceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataSourceService_ServiceDesc is the grpc.ServiceDesc for DataSourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataSourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracker.v1beta1.DataSourceService",
	HandlerType: (*DataSourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _DataSourceService_Write_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WriteStream",
			Handler:       _DataSourceService_WriteStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1beta1/datasource.proto",
}
