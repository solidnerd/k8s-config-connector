// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/pubsublite/v1/publisher.proto

package pubsublitepb

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

// PublisherServiceClient is the client API for PublisherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublisherServiceClient interface {
	// Establishes a stream with the server for publishing messages. Once the
	// stream is initialized, the client publishes messages by sending publish
	// requests on the stream. The server responds with a PublishResponse for each
	// PublishRequest sent by the client, in the same order that the requests
	// were sent. Note that multiple PublishRequests can be in flight
	// simultaneously, but they will be processed by the server in the order that
	// they are sent by the client on a given stream.
	Publish(ctx context.Context, opts ...grpc.CallOption) (PublisherService_PublishClient, error)
}

type publisherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherServiceClient(cc grpc.ClientConnInterface) PublisherServiceClient {
	return &publisherServiceClient{cc}
}

func (c *publisherServiceClient) Publish(ctx context.Context, opts ...grpc.CallOption) (PublisherService_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &PublisherService_ServiceDesc.Streams[0], "/mockgcp.cloud.pubsublite.v1.PublisherService/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherServicePublishClient{stream}
	return x, nil
}

type PublisherService_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type publisherServicePublishClient struct {
	grpc.ClientStream
}

func (x *publisherServicePublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *publisherServicePublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherServiceServer is the server API for PublisherService service.
// All implementations must embed UnimplementedPublisherServiceServer
// for forward compatibility
type PublisherServiceServer interface {
	// Establishes a stream with the server for publishing messages. Once the
	// stream is initialized, the client publishes messages by sending publish
	// requests on the stream. The server responds with a PublishResponse for each
	// PublishRequest sent by the client, in the same order that the requests
	// were sent. Note that multiple PublishRequests can be in flight
	// simultaneously, but they will be processed by the server in the order that
	// they are sent by the client on a given stream.
	Publish(PublisherService_PublishServer) error
	mustEmbedUnimplementedPublisherServiceServer()
}

// UnimplementedPublisherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublisherServiceServer struct {
}

func (UnimplementedPublisherServiceServer) Publish(PublisherService_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedPublisherServiceServer) mustEmbedUnimplementedPublisherServiceServer() {}

// UnsafePublisherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublisherServiceServer will
// result in compilation errors.
type UnsafePublisherServiceServer interface {
	mustEmbedUnimplementedPublisherServiceServer()
}

func RegisterPublisherServiceServer(s grpc.ServiceRegistrar, srv PublisherServiceServer) {
	s.RegisterService(&PublisherService_ServiceDesc, srv)
}

func _PublisherService_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PublisherServiceServer).Publish(&publisherServicePublishServer{stream})
}

type PublisherService_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type publisherServicePublishServer struct {
	grpc.ServerStream
}

func (x *publisherServicePublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *publisherServicePublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherService_ServiceDesc is the grpc.ServiceDesc for PublisherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublisherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.pubsublite.v1.PublisherService",
	HandlerType: (*PublisherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _PublisherService_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mockgcp/cloud/pubsublite/v1/publisher.proto",
}
