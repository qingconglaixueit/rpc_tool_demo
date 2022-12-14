// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: tenant.proto

package tenant

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

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	GetTenantDetail(ctx context.Context, in *TenantDetailReq, opts ...grpc.CallOption) (*TenantDetailRsp, error)
	GetTenantList(ctx context.Context, in *TenantListReq, opts ...grpc.CallOption) (*TenantListRsp, error)
}

type tenantClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantClient(cc grpc.ClientConnInterface) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) GetTenantDetail(ctx context.Context, in *TenantDetailReq, opts ...grpc.CallOption) (*TenantDetailRsp, error) {
	out := new(TenantDetailRsp)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/get_tenant_detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantClient) GetTenantList(ctx context.Context, in *TenantListReq, opts ...grpc.CallOption) (*TenantListRsp, error) {
	out := new(TenantListRsp)
	err := c.cc.Invoke(ctx, "/tenant.Tenant/get_tenant_list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	GetTenantDetail(context.Context, *TenantDetailReq) (*TenantDetailRsp, error)
	GetTenantList(context.Context, *TenantListReq) (*TenantListRsp, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
}

func (UnimplementedTenantServer) GetTenantDetail(context.Context, *TenantDetailReq) (*TenantDetailRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantDetail not implemented")
}
func (UnimplementedTenantServer) GetTenantList(context.Context, *TenantListReq) (*TenantListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantList not implemented")
}
func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_GetTenantDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetTenantDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/get_tenant_detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetTenantDetail(ctx, req.(*TenantDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetTenantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServer).GetTenantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.Tenant/get_tenant_list",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServer).GetTenantList(ctx, req.(*TenantListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.Tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_tenant_detail",
			Handler:    _Tenant_GetTenantDetail_Handler,
		},
		{
			MethodName: "get_tenant_list",
			Handler:    _Tenant_GetTenantList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant.proto",
}
