// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthValidationServiceClient is the client API for AuthValidationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthValidationServiceClient interface {
	Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Responce, error)
}

type authValidationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthValidationServiceClient(cc grpc.ClientConnInterface) AuthValidationServiceClient {
	return &authValidationServiceClient{cc}
}

func (c *authValidationServiceClient) Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Responce, error) {
	out := new(Responce)
	err := c.cc.Invoke(ctx, "/auth.authValidationService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthValidationServiceServer is the server API for AuthValidationService service.
// All implementations must embed UnimplementedAuthValidationServiceServer
// for forward compatibility
type AuthValidationServiceServer interface {
	Validate(context.Context, *Request) (*Responce, error)
	mustEmbedUnimplementedAuthValidationServiceServer()
}

// UnimplementedAuthValidationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthValidationServiceServer struct {
}

func (UnimplementedAuthValidationServiceServer) Validate(context.Context, *Request) (*Responce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthValidationServiceServer) mustEmbedUnimplementedAuthValidationServiceServer() {}

// UnsafeAuthValidationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthValidationServiceServer will
// result in compilation errors.
type UnsafeAuthValidationServiceServer interface {
	mustEmbedUnimplementedAuthValidationServiceServer()
}

func RegisterAuthValidationServiceServer(s grpc.ServiceRegistrar, srv AuthValidationServiceServer) {
	s.RegisterService(&_AuthValidationService_serviceDesc, srv)
}

func _AuthValidationService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthValidationServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.authValidationService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthValidationServiceServer).Validate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthValidationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.authValidationService",
	HandlerType: (*AuthValidationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _AuthValidationService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}