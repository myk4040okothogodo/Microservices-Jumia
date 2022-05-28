// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: protos/JumiaServiceA.proto

package protos

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

// ServiceAClient is the client API for ServiceA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceAClient interface {
	GetCsvData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
}

type serviceAClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAClient(cc grpc.ClientConnInterface) ServiceAClient {
	return &serviceAClient{cc}
}

func (c *serviceAClient) GetCsvData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, "/serviceA/GetCsvData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAServer is the server API for ServiceA service.
// All implementations must embed UnimplementedServiceAServer
// for forward compatibility
type ServiceAServer interface {
	GetCsvData(context.Context, *DataRequest) (*DataResponse, error)
	//mustEmbedUnimplementedServiceAServer()
}

// UnimplementedServiceAServer must be embedded to have forward compatible implementations.
type UnimplementedServiceAServer struct {
}

func (UnimplementedServiceAServer) GetCsvData(context.Context, *DataRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCsvData not implemented")
}
//func (UnimplementedServiceAServer) mustEmbedUnimplementedServiceAServer() {}

// UnsafeServiceAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAServer will
// result in compilation errors.
//type UnsafeServiceAServer interface {
//	mustEmbedUnimplementedServiceAServer()
//}

func RegisterServiceAServer(s grpc.ServiceRegistrar, srv ServiceAServer) {
	s.RegisterService(&ServiceA_ServiceDesc, srv)
}

func _ServiceA_GetCsvData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAServer).GetCsvData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serviceA/GetCsvData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAServer).GetCsvData(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceA_ServiceDesc is the grpc.ServiceDesc for ServiceA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "serviceA",
	HandlerType: (*ServiceAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCsvData",
			Handler:    _ServiceA_GetCsvData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/JumiaServiceA.proto",
}
