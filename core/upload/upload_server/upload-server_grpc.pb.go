// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: upload-server.proto

package upload_server

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
	UploadServer_GeneratePresignedURL_FullMethodName = "/upload_server.UploadServer/GeneratePresignedURL"
)

// UploadServerClient is the client API for UploadServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServerClient interface {
	GeneratePresignedURL(ctx context.Context, in *PresighedURLRequest, opts ...grpc.CallOption) (*PresighedURLResponse, error)
}

type uploadServerClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServerClient(cc grpc.ClientConnInterface) UploadServerClient {
	return &uploadServerClient{cc}
}

func (c *uploadServerClient) GeneratePresignedURL(ctx context.Context, in *PresighedURLRequest, opts ...grpc.CallOption) (*PresighedURLResponse, error) {
	out := new(PresighedURLResponse)
	err := c.cc.Invoke(ctx, UploadServer_GeneratePresignedURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServerServer is the server API for UploadServer service.
// All implementations must embed UnimplementedUploadServerServer
// for forward compatibility
type UploadServerServer interface {
	GeneratePresignedURL(context.Context, *PresighedURLRequest) (*PresighedURLResponse, error)
	mustEmbedUnimplementedUploadServerServer()
}

// UnimplementedUploadServerServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServerServer struct {
}

func (UnimplementedUploadServerServer) GeneratePresignedURL(context.Context, *PresighedURLRequest) (*PresighedURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePresignedURL not implemented")
}
func (UnimplementedUploadServerServer) mustEmbedUnimplementedUploadServerServer() {}

// UnsafeUploadServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServerServer will
// result in compilation errors.
type UnsafeUploadServerServer interface {
	mustEmbedUnimplementedUploadServerServer()
}

func RegisterUploadServerServer(s grpc.ServiceRegistrar, srv UploadServerServer) {
	s.RegisterService(&UploadServer_ServiceDesc, srv)
}

func _UploadServer_GeneratePresignedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PresighedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServerServer).GeneratePresignedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadServer_GeneratePresignedURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServerServer).GeneratePresignedURL(ctx, req.(*PresighedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadServer_ServiceDesc is the grpc.ServiceDesc for UploadServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upload_server.UploadServer",
	HandlerType: (*UploadServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePresignedURL",
			Handler:    _UploadServer_GeneratePresignedURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upload-server.proto",
}
