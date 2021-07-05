// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcimpl

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

// CAClient is the client API for CA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CAClient interface {
	RequestPolicy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyReply, error)
	GenerateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error)
}

type cAClient struct {
	cc grpc.ClientConnInterface
}

func NewCAClient(cc grpc.ClientConnInterface) CAClient {
	return &cAClient{cc}
}

func (c *cAClient) RequestPolicy(ctx context.Context, in *PolicyRequest, opts ...grpc.CallOption) (*PolicyReply, error) {
	out := new(PolicyReply)
	err := c.cc.Invoke(ctx, "/edgeca.CA/RequestPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAClient) GenerateCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateReply, error) {
	out := new(CertificateReply)
	err := c.cc.Invoke(ctx, "/edgeca.CA/GenerateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAServer is the server API for CA service.
// All implementations must embed UnimplementedCAServer
// for forward compatibility
type CAServer interface {
	RequestPolicy(context.Context, *PolicyRequest) (*PolicyReply, error)
	GenerateCertificate(context.Context, *CertificateRequest) (*CertificateReply, error)
	mustEmbedUnimplementedCAServer()
}

// UnimplementedCAServer must be embedded to have forward compatible implementations.
type UnimplementedCAServer struct {
}

func (UnimplementedCAServer) RequestPolicy(context.Context, *PolicyRequest) (*PolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPolicy not implemented")
}
func (UnimplementedCAServer) GenerateCertificate(context.Context, *CertificateRequest) (*CertificateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCertificate not implemented")
}
func (UnimplementedCAServer) mustEmbedUnimplementedCAServer() {}

// UnsafeCAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CAServer will
// result in compilation errors.
type UnsafeCAServer interface {
	mustEmbedUnimplementedCAServer()
}

func RegisterCAServer(s grpc.ServiceRegistrar, srv CAServer) {
	s.RegisterService(&CA_ServiceDesc, srv)
}

func _CA_RequestPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).RequestPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeca.CA/RequestPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).RequestPolicy(ctx, req.(*PolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CA_GenerateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServer).GenerateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgeca.CA/GenerateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServer).GenerateCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CA_ServiceDesc is the grpc.ServiceDesc for CA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edgeca.CA",
	HandlerType: (*CAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestPolicy",
			Handler:    _CA_RequestPolicy_Handler,
		},
		{
			MethodName: "GenerateCertificate",
			Handler:    _CA_GenerateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edgeca.proto",
}