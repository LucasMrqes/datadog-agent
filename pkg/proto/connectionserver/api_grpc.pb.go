// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: connectionserver/api.proto

package connectionserver

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

// SystemProbeClient is the client API for SystemProbe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemProbeClient interface {
	GetConnections(ctx context.Context, in *GetConnectionsRequest, opts ...grpc.CallOption) (SystemProbe_GetConnectionsClient, error)
}

type systemProbeClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemProbeClient(cc grpc.ClientConnInterface) SystemProbeClient {
	return &systemProbeClient{cc}
}

func (c *systemProbeClient) GetConnections(ctx context.Context, in *GetConnectionsRequest, opts ...grpc.CallOption) (SystemProbe_GetConnectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SystemProbe_ServiceDesc.Streams[0], "/datadog.test.v1.SystemProbe/GetConnections", opts...)
	if err != nil {
		return nil, err
	}
	x := &systemProbeGetConnectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SystemProbe_GetConnectionsClient interface {
	Recv() (*Connection, error)
	grpc.ClientStream
}

type systemProbeGetConnectionsClient struct {
	grpc.ClientStream
}

func (x *systemProbeGetConnectionsClient) Recv() (*Connection, error) {
	m := new(Connection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SystemProbeServer is the server API for SystemProbe service.
// All implementations must embed UnimplementedSystemProbeServer
// for forward compatibility
type SystemProbeServer interface {
	GetConnections(*GetConnectionsRequest, SystemProbe_GetConnectionsServer) error
	mustEmbedUnimplementedSystemProbeServer()
}

// UnimplementedSystemProbeServer must be embedded to have forward compatible implementations.
type UnimplementedSystemProbeServer struct {
}

func (UnimplementedSystemProbeServer) GetConnections(*GetConnectionsRequest, SystemProbe_GetConnectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConnections not implemented")
}
func (UnimplementedSystemProbeServer) mustEmbedUnimplementedSystemProbeServer() {}

// UnsafeSystemProbeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemProbeServer will
// result in compilation errors.
type UnsafeSystemProbeServer interface {
	mustEmbedUnimplementedSystemProbeServer()
}

func RegisterSystemProbeServer(s grpc.ServiceRegistrar, srv SystemProbeServer) {
	s.RegisterService(&SystemProbe_ServiceDesc, srv)
}

func _SystemProbe_GetConnections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetConnectionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SystemProbeServer).GetConnections(m, &systemProbeGetConnectionsServer{stream})
}

type SystemProbe_GetConnectionsServer interface {
	Send(*Connection) error
	grpc.ServerStream
}

type systemProbeGetConnectionsServer struct {
	grpc.ServerStream
}

func (x *systemProbeGetConnectionsServer) Send(m *Connection) error {
	return x.ServerStream.SendMsg(m)
}

// SystemProbe_ServiceDesc is the grpc.ServiceDesc for SystemProbe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemProbe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datadog.test.v1.SystemProbe",
	HandlerType: (*SystemProbeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConnections",
			Handler:       _SystemProbe_GetConnections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "connectionserver/api.proto",
}
