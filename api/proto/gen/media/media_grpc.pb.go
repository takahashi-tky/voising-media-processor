// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: media.proto

package media

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Media_PostUserProfileImage_FullMethodName       = "/media.Media/PostUserProfileImage"
	Media_PostUserReportCoverImage_FullMethodName   = "/media.Media/PostUserReportCoverImage"
	Media_PostUserReportContentImage_FullMethodName = "/media.Media/PostUserReportContentImage"
	Media_PatchUserImageStatus_FullMethodName       = "/media.Media/PatchUserImageStatus"
	Media_PathUserImageName_FullMethodName          = "/media.Media/PathUserImageName"
	Media_CreateUserProfileImage_FullMethodName     = "/media.Media/CreateUserProfileImage"
)

// MediaClient is the client API for Media service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaClient interface {
	PostUserProfileImage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PostUserProfileImageResponse, error)
	PostUserReportCoverImage(ctx context.Context, in *PostUserReportCoverImageRequest, opts ...grpc.CallOption) (*PostUserReportCoverImageResponse, error)
	PostUserReportContentImage(ctx context.Context, in *PostUserReportContentImageRequest, opts ...grpc.CallOption) (*PostUserReportContentImageResponse, error)
	PatchUserImageStatus(ctx context.Context, in *PatchUserImageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PathUserImageName(ctx context.Context, in *PatchUserImageNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUserProfileImage(ctx context.Context, in *CreateUserProfileImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type mediaClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaClient(cc grpc.ClientConnInterface) MediaClient {
	return &mediaClient{cc}
}

func (c *mediaClient) PostUserProfileImage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PostUserProfileImageResponse, error) {
	out := new(PostUserProfileImageResponse)
	err := c.cc.Invoke(ctx, Media_PostUserProfileImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) PostUserReportCoverImage(ctx context.Context, in *PostUserReportCoverImageRequest, opts ...grpc.CallOption) (*PostUserReportCoverImageResponse, error) {
	out := new(PostUserReportCoverImageResponse)
	err := c.cc.Invoke(ctx, Media_PostUserReportCoverImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) PostUserReportContentImage(ctx context.Context, in *PostUserReportContentImageRequest, opts ...grpc.CallOption) (*PostUserReportContentImageResponse, error) {
	out := new(PostUserReportContentImageResponse)
	err := c.cc.Invoke(ctx, Media_PostUserReportContentImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) PatchUserImageStatus(ctx context.Context, in *PatchUserImageStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Media_PatchUserImageStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) PathUserImageName(ctx context.Context, in *PatchUserImageNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Media_PathUserImageName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaClient) CreateUserProfileImage(ctx context.Context, in *CreateUserProfileImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Media_CreateUserProfileImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServer is the server API for Media service.
// All implementations must embed UnimplementedMediaServer
// for forward compatibility
type MediaServer interface {
	PostUserProfileImage(context.Context, *emptypb.Empty) (*PostUserProfileImageResponse, error)
	PostUserReportCoverImage(context.Context, *PostUserReportCoverImageRequest) (*PostUserReportCoverImageResponse, error)
	PostUserReportContentImage(context.Context, *PostUserReportContentImageRequest) (*PostUserReportContentImageResponse, error)
	PatchUserImageStatus(context.Context, *PatchUserImageStatusRequest) (*emptypb.Empty, error)
	PathUserImageName(context.Context, *PatchUserImageNameRequest) (*emptypb.Empty, error)
	CreateUserProfileImage(context.Context, *CreateUserProfileImageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMediaServer()
}

// UnimplementedMediaServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServer struct {
}

func (UnimplementedMediaServer) PostUserProfileImage(context.Context, *emptypb.Empty) (*PostUserProfileImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserProfileImage not implemented")
}
func (UnimplementedMediaServer) PostUserReportCoverImage(context.Context, *PostUserReportCoverImageRequest) (*PostUserReportCoverImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserReportCoverImage not implemented")
}
func (UnimplementedMediaServer) PostUserReportContentImage(context.Context, *PostUserReportContentImageRequest) (*PostUserReportContentImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserReportContentImage not implemented")
}
func (UnimplementedMediaServer) PatchUserImageStatus(context.Context, *PatchUserImageStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchUserImageStatus not implemented")
}
func (UnimplementedMediaServer) PathUserImageName(context.Context, *PatchUserImageNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PathUserImageName not implemented")
}
func (UnimplementedMediaServer) CreateUserProfileImage(context.Context, *CreateUserProfileImageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserProfileImage not implemented")
}
func (UnimplementedMediaServer) mustEmbedUnimplementedMediaServer() {}

// UnsafeMediaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServer will
// result in compilation errors.
type UnsafeMediaServer interface {
	mustEmbedUnimplementedMediaServer()
}

func RegisterMediaServer(s grpc.ServiceRegistrar, srv MediaServer) {
	s.RegisterService(&Media_ServiceDesc, srv)
}

func _Media_PostUserProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).PostUserProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_PostUserProfileImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).PostUserProfileImage(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_PostUserReportCoverImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserReportCoverImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).PostUserReportCoverImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_PostUserReportCoverImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).PostUserReportCoverImage(ctx, req.(*PostUserReportCoverImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_PostUserReportContentImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserReportContentImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).PostUserReportContentImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_PostUserReportContentImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).PostUserReportContentImage(ctx, req.(*PostUserReportContentImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_PatchUserImageStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchUserImageStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).PatchUserImageStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_PatchUserImageStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).PatchUserImageStatus(ctx, req.(*PatchUserImageStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_PathUserImageName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchUserImageNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).PathUserImageName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_PathUserImageName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).PathUserImageName(ctx, req.(*PatchUserImageNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Media_CreateUserProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserProfileImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServer).CreateUserProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Media_CreateUserProfileImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServer).CreateUserProfileImage(ctx, req.(*CreateUserProfileImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Media_ServiceDesc is the grpc.ServiceDesc for Media service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Media_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "media.Media",
	HandlerType: (*MediaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostUserProfileImage",
			Handler:    _Media_PostUserProfileImage_Handler,
		},
		{
			MethodName: "PostUserReportCoverImage",
			Handler:    _Media_PostUserReportCoverImage_Handler,
		},
		{
			MethodName: "PostUserReportContentImage",
			Handler:    _Media_PostUserReportContentImage_Handler,
		},
		{
			MethodName: "PatchUserImageStatus",
			Handler:    _Media_PatchUserImageStatus_Handler,
		},
		{
			MethodName: "PathUserImageName",
			Handler:    _Media_PathUserImageName_Handler,
		},
		{
			MethodName: "CreateUserProfileImage",
			Handler:    _Media_CreateUserProfileImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "media.proto",
}
