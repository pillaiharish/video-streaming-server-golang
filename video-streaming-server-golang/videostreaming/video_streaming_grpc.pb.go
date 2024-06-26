// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: video_streaming.proto

package videostreaming

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VideoStreamingService_StreamVideo_FullMethodName = "/videostreaming.VideoStreamingService/StreamVideo"
	VideoStreamingService_ListVideos_FullMethodName  = "/videostreaming.VideoStreamingService/ListVideos"
)

// VideoStreamingServiceClient is the client API for VideoStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStreamingServiceClient interface {
	StreamVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStreamingService_StreamVideoClient, error)
	ListVideos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VideoList, error)
}

type videoStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStreamingServiceClient(cc grpc.ClientConnInterface) VideoStreamingServiceClient {
	return &videoStreamingServiceClient{cc}
}

func (c *videoStreamingServiceClient) StreamVideo(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (VideoStreamingService_StreamVideoClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoStreamingService_ServiceDesc.Streams[0], VideoStreamingService_StreamVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &videoStreamingServiceStreamVideoClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoStreamingService_StreamVideoClient interface {
	Recv() (*VideoChunk, error)
	grpc.ClientStream
}

type videoStreamingServiceStreamVideoClient struct {
	grpc.ClientStream
}

func (x *videoStreamingServiceStreamVideoClient) Recv() (*VideoChunk, error) {
	m := new(VideoChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoStreamingServiceClient) ListVideos(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*VideoList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VideoList)
	err := c.cc.Invoke(ctx, VideoStreamingService_ListVideos_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoStreamingServiceServer is the server API for VideoStreamingService service.
// All implementations must embed UnimplementedVideoStreamingServiceServer
// for forward compatibility
type VideoStreamingServiceServer interface {
	StreamVideo(*VideoRequest, VideoStreamingService_StreamVideoServer) error
	ListVideos(context.Context, *Empty) (*VideoList, error)
	mustEmbedUnimplementedVideoStreamingServiceServer()
}

// UnimplementedVideoStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStreamingServiceServer struct {
}

func (UnimplementedVideoStreamingServiceServer) StreamVideo(*VideoRequest, VideoStreamingService_StreamVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamVideo not implemented")
}
func (UnimplementedVideoStreamingServiceServer) ListVideos(context.Context, *Empty) (*VideoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideos not implemented")
}
func (UnimplementedVideoStreamingServiceServer) mustEmbedUnimplementedVideoStreamingServiceServer() {}

// UnsafeVideoStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStreamingServiceServer will
// result in compilation errors.
type UnsafeVideoStreamingServiceServer interface {
	mustEmbedUnimplementedVideoStreamingServiceServer()
}

func RegisterVideoStreamingServiceServer(s grpc.ServiceRegistrar, srv VideoStreamingServiceServer) {
	s.RegisterService(&VideoStreamingService_ServiceDesc, srv)
}

func _VideoStreamingService_StreamVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VideoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoStreamingServiceServer).StreamVideo(m, &videoStreamingServiceStreamVideoServer{ServerStream: stream})
}

type VideoStreamingService_StreamVideoServer interface {
	Send(*VideoChunk) error
	grpc.ServerStream
}

type videoStreamingServiceStreamVideoServer struct {
	grpc.ServerStream
}

func (x *videoStreamingServiceStreamVideoServer) Send(m *VideoChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _VideoStreamingService_ListVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoStreamingServiceServer).ListVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoStreamingService_ListVideos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoStreamingServiceServer).ListVideos(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoStreamingService_ServiceDesc is the grpc.ServiceDesc for VideoStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "videostreaming.VideoStreamingService",
	HandlerType: (*VideoStreamingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVideos",
			Handler:    _VideoStreamingService_ListVideos_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamVideo",
			Handler:       _VideoStreamingService_StreamVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "video_streaming.proto",
}
