// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/monitoring/v3/alert_service.proto

package monitoringpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AlertPolicyServiceClient is the client API for AlertPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertPolicyServiceClient interface {
	// Lists the existing alerting policies for the workspace.
	ListAlertPolicies(ctx context.Context, in *ListAlertPoliciesRequest, opts ...grpc.CallOption) (*ListAlertPoliciesResponse, error)
	// Gets a single alerting policy.
	GetAlertPolicy(ctx context.Context, in *GetAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
	// Creates a new alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	CreateAlertPolicy(ctx context.Context, in *CreateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
	// Deletes an alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	DeleteAlertPolicy(ctx context.Context, in *DeleteAlertPolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates an alerting policy. You can either replace the entire policy with
	// a new one or replace only certain fields in the current alerting policy by
	// specifying the fields to be updated via `updateMask`. Returns the
	// updated alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	UpdateAlertPolicy(ctx context.Context, in *UpdateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error)
}

type alertPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertPolicyServiceClient(cc grpc.ClientConnInterface) AlertPolicyServiceClient {
	return &alertPolicyServiceClient{cc}
}

func (c *alertPolicyServiceClient) ListAlertPolicies(ctx context.Context, in *ListAlertPoliciesRequest, opts ...grpc.CallOption) (*ListAlertPoliciesResponse, error) {
	out := new(ListAlertPoliciesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.v3.AlertPolicyService/ListAlertPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) GetAlertPolicy(ctx context.Context, in *GetAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.v3.AlertPolicyService/GetAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) CreateAlertPolicy(ctx context.Context, in *CreateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.v3.AlertPolicyService/CreateAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) DeleteAlertPolicy(ctx context.Context, in *DeleteAlertPolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.v3.AlertPolicyService/DeleteAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertPolicyServiceClient) UpdateAlertPolicy(ctx context.Context, in *UpdateAlertPolicyRequest, opts ...grpc.CallOption) (*AlertPolicy, error) {
	out := new(AlertPolicy)
	err := c.cc.Invoke(ctx, "/mockgcp.monitoring.v3.AlertPolicyService/UpdateAlertPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertPolicyServiceServer is the server API for AlertPolicyService service.
// All implementations must embed UnimplementedAlertPolicyServiceServer
// for forward compatibility
type AlertPolicyServiceServer interface {
	// Lists the existing alerting policies for the workspace.
	ListAlertPolicies(context.Context, *ListAlertPoliciesRequest) (*ListAlertPoliciesResponse, error)
	// Gets a single alerting policy.
	GetAlertPolicy(context.Context, *GetAlertPolicyRequest) (*AlertPolicy, error)
	// Creates a new alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	CreateAlertPolicy(context.Context, *CreateAlertPolicyRequest) (*AlertPolicy, error)
	// Deletes an alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	DeleteAlertPolicy(context.Context, *DeleteAlertPolicyRequest) (*empty.Empty, error)
	// Updates an alerting policy. You can either replace the entire policy with
	// a new one or replace only certain fields in the current alerting policy by
	// specifying the fields to be updated via `updateMask`. Returns the
	// updated alerting policy.
	//
	// Design your application to single-thread API calls that modify the state of
	// alerting policies in a single project. This includes calls to
	// CreateAlertPolicy, DeleteAlertPolicy and UpdateAlertPolicy.
	UpdateAlertPolicy(context.Context, *UpdateAlertPolicyRequest) (*AlertPolicy, error)
	mustEmbedUnimplementedAlertPolicyServiceServer()
}

// UnimplementedAlertPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAlertPolicyServiceServer struct {
}

func (UnimplementedAlertPolicyServiceServer) ListAlertPolicies(context.Context, *ListAlertPoliciesRequest) (*ListAlertPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertPolicies not implemented")
}
func (UnimplementedAlertPolicyServiceServer) GetAlertPolicy(context.Context, *GetAlertPolicyRequest) (*AlertPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertPolicy not implemented")
}
func (UnimplementedAlertPolicyServiceServer) CreateAlertPolicy(context.Context, *CreateAlertPolicyRequest) (*AlertPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertPolicy not implemented")
}
func (UnimplementedAlertPolicyServiceServer) DeleteAlertPolicy(context.Context, *DeleteAlertPolicyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertPolicy not implemented")
}
func (UnimplementedAlertPolicyServiceServer) UpdateAlertPolicy(context.Context, *UpdateAlertPolicyRequest) (*AlertPolicy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertPolicy not implemented")
}
func (UnimplementedAlertPolicyServiceServer) mustEmbedUnimplementedAlertPolicyServiceServer() {}

// UnsafeAlertPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertPolicyServiceServer will
// result in compilation errors.
type UnsafeAlertPolicyServiceServer interface {
	mustEmbedUnimplementedAlertPolicyServiceServer()
}

func RegisterAlertPolicyServiceServer(s grpc.ServiceRegistrar, srv AlertPolicyServiceServer) {
	s.RegisterService(&AlertPolicyService_ServiceDesc, srv)
}

func _AlertPolicyService_ListAlertPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlertPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).ListAlertPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.v3.AlertPolicyService/ListAlertPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).ListAlertPolicies(ctx, req.(*ListAlertPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_GetAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).GetAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.v3.AlertPolicyService/GetAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).GetAlertPolicy(ctx, req.(*GetAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_CreateAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).CreateAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.v3.AlertPolicyService/CreateAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).CreateAlertPolicy(ctx, req.(*CreateAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_DeleteAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).DeleteAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.v3.AlertPolicyService/DeleteAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).DeleteAlertPolicy(ctx, req.(*DeleteAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertPolicyService_UpdateAlertPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertPolicyServiceServer).UpdateAlertPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.monitoring.v3.AlertPolicyService/UpdateAlertPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertPolicyServiceServer).UpdateAlertPolicy(ctx, req.(*UpdateAlertPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertPolicyService_ServiceDesc is the grpc.ServiceDesc for AlertPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.monitoring.v3.AlertPolicyService",
	HandlerType: (*AlertPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAlertPolicies",
			Handler:    _AlertPolicyService_ListAlertPolicies_Handler,
		},
		{
			MethodName: "GetAlertPolicy",
			Handler:    _AlertPolicyService_GetAlertPolicy_Handler,
		},
		{
			MethodName: "CreateAlertPolicy",
			Handler:    _AlertPolicyService_CreateAlertPolicy_Handler,
		},
		{
			MethodName: "DeleteAlertPolicy",
			Handler:    _AlertPolicyService_DeleteAlertPolicy_Handler,
		},
		{
			MethodName: "UpdateAlertPolicy",
			Handler:    _AlertPolicyService_UpdateAlertPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/monitoring/v3/alert_service.proto",
}
