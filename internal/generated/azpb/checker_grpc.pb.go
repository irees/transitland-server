// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: checker.proto

package azpb

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

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckerClient interface {
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	TenantList(ctx context.Context, in *TenantListRequest, opts ...grpc.CallOption) (*TenantListResponse, error)
	Tenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error)
	TenantPermissions(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantPermissionsResponse, error)
	TenantSave(ctx context.Context, in *TenantSaveRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error)
	TenantAddPermission(ctx context.Context, in *TenantModifyPermissionRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error)
	TenantRemovePermission(ctx context.Context, in *TenantModifyPermissionRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error)
	TenantCreate(ctx context.Context, in *TenantCreateRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error)
	TenantCreateGroup(ctx context.Context, in *TenantCreateGroupRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error)
	GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error)
	Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	GroupPermissions(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupPermissionsResponse, error)
	GroupSave(ctx context.Context, in *GroupSaveRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error)
	GroupAddPermission(ctx context.Context, in *GroupModifyPermissionRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error)
	GroupRemovePermission(ctx context.Context, in *GroupModifyPermissionRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error)
	GroupSetTenant(ctx context.Context, in *GroupSetTenantRequest, opts ...grpc.CallOption) (*GroupSetTenantResponse, error)
	FeedList(ctx context.Context, in *FeedListRequest, opts ...grpc.CallOption) (*FeedListResponse, error)
	Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	FeedPermissions(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedPermissionsResponse, error)
	FeedSetGroup(ctx context.Context, in *FeedSetGroupRequest, opts ...grpc.CallOption) (*FeedSaveResponse, error)
	FeedVersionList(ctx context.Context, in *FeedVersionListRequest, opts ...grpc.CallOption) (*FeedVersionListResponse, error)
	FeedVersion(ctx context.Context, in *FeedVersionRequest, opts ...grpc.CallOption) (*FeedVersionResponse, error)
	FeedVersionPermissions(ctx context.Context, in *FeedVersionRequest, opts ...grpc.CallOption) (*FeedVersionPermissionsResponse, error)
	FeedVersionAddPermission(ctx context.Context, in *FeedVersionModifyPermissionRequest, opts ...grpc.CallOption) (*FeedVersionSaveResponse, error)
	FeedVersionRemovePermission(ctx context.Context, in *FeedVersionModifyPermissionRequest, opts ...grpc.CallOption) (*FeedVersionSaveResponse, error)
}

type checkerClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckerClient(cc grpc.ClientConnInterface) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) User(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantList(ctx context.Context, in *TenantListRequest, opts ...grpc.CallOption) (*TenantListResponse, error) {
	out := new(TenantListResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) Tenant(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantResponse, error) {
	out := new(TenantResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/Tenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantPermissions(ctx context.Context, in *TenantRequest, opts ...grpc.CallOption) (*TenantPermissionsResponse, error) {
	out := new(TenantPermissionsResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantSave(ctx context.Context, in *TenantSaveRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error) {
	out := new(TenantSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantAddPermission(ctx context.Context, in *TenantModifyPermissionRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error) {
	out := new(TenantSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantAddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantRemovePermission(ctx context.Context, in *TenantModifyPermissionRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error) {
	out := new(TenantSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantRemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantCreate(ctx context.Context, in *TenantCreateRequest, opts ...grpc.CallOption) (*TenantSaveResponse, error) {
	out := new(TenantSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) TenantCreateGroup(ctx context.Context, in *TenantCreateGroupRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error) {
	out := new(GroupSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/TenantCreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupList(ctx context.Context, in *GroupListRequest, opts ...grpc.CallOption) (*GroupListResponse, error) {
	out := new(GroupListResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) Group(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/Group", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupPermissions(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupPermissionsResponse, error) {
	out := new(GroupPermissionsResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupSave(ctx context.Context, in *GroupSaveRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error) {
	out := new(GroupSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupAddPermission(ctx context.Context, in *GroupModifyPermissionRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error) {
	out := new(GroupSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupAddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupRemovePermission(ctx context.Context, in *GroupModifyPermissionRequest, opts ...grpc.CallOption) (*GroupSaveResponse, error) {
	out := new(GroupSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupRemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) GroupSetTenant(ctx context.Context, in *GroupSetTenantRequest, opts ...grpc.CallOption) (*GroupSetTenantResponse, error) {
	out := new(GroupSetTenantResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/GroupSetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedList(ctx context.Context, in *FeedListRequest, opts ...grpc.CallOption) (*FeedListResponse, error) {
	out := new(FeedListResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) Feed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedPermissions(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedPermissionsResponse, error) {
	out := new(FeedPermissionsResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedSetGroup(ctx context.Context, in *FeedSetGroupRequest, opts ...grpc.CallOption) (*FeedSaveResponse, error) {
	out := new(FeedSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedSetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedVersionList(ctx context.Context, in *FeedVersionListRequest, opts ...grpc.CallOption) (*FeedVersionListResponse, error) {
	out := new(FeedVersionListResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedVersionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedVersion(ctx context.Context, in *FeedVersionRequest, opts ...grpc.CallOption) (*FeedVersionResponse, error) {
	out := new(FeedVersionResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedVersionPermissions(ctx context.Context, in *FeedVersionRequest, opts ...grpc.CallOption) (*FeedVersionPermissionsResponse, error) {
	out := new(FeedVersionPermissionsResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedVersionPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedVersionAddPermission(ctx context.Context, in *FeedVersionModifyPermissionRequest, opts ...grpc.CallOption) (*FeedVersionSaveResponse, error) {
	out := new(FeedVersionSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedVersionAddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) FeedVersionRemovePermission(ctx context.Context, in *FeedVersionModifyPermissionRequest, opts ...grpc.CallOption) (*FeedVersionSaveResponse, error) {
	out := new(FeedVersionSaveResponse)
	err := c.cc.Invoke(ctx, "/azpb.Checker/FeedVersionRemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
// All implementations must embed UnimplementedCheckerServer
// for forward compatibility
type CheckerServer interface {
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
	User(context.Context, *UserRequest) (*UserResponse, error)
	TenantList(context.Context, *TenantListRequest) (*TenantListResponse, error)
	Tenant(context.Context, *TenantRequest) (*TenantResponse, error)
	TenantPermissions(context.Context, *TenantRequest) (*TenantPermissionsResponse, error)
	TenantSave(context.Context, *TenantSaveRequest) (*TenantSaveResponse, error)
	TenantAddPermission(context.Context, *TenantModifyPermissionRequest) (*TenantSaveResponse, error)
	TenantRemovePermission(context.Context, *TenantModifyPermissionRequest) (*TenantSaveResponse, error)
	TenantCreate(context.Context, *TenantCreateRequest) (*TenantSaveResponse, error)
	TenantCreateGroup(context.Context, *TenantCreateGroupRequest) (*GroupSaveResponse, error)
	GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error)
	Group(context.Context, *GroupRequest) (*GroupResponse, error)
	GroupPermissions(context.Context, *GroupRequest) (*GroupPermissionsResponse, error)
	GroupSave(context.Context, *GroupSaveRequest) (*GroupSaveResponse, error)
	GroupAddPermission(context.Context, *GroupModifyPermissionRequest) (*GroupSaveResponse, error)
	GroupRemovePermission(context.Context, *GroupModifyPermissionRequest) (*GroupSaveResponse, error)
	GroupSetTenant(context.Context, *GroupSetTenantRequest) (*GroupSetTenantResponse, error)
	FeedList(context.Context, *FeedListRequest) (*FeedListResponse, error)
	Feed(context.Context, *FeedRequest) (*FeedResponse, error)
	FeedPermissions(context.Context, *FeedRequest) (*FeedPermissionsResponse, error)
	FeedSetGroup(context.Context, *FeedSetGroupRequest) (*FeedSaveResponse, error)
	FeedVersionList(context.Context, *FeedVersionListRequest) (*FeedVersionListResponse, error)
	FeedVersion(context.Context, *FeedVersionRequest) (*FeedVersionResponse, error)
	FeedVersionPermissions(context.Context, *FeedVersionRequest) (*FeedVersionPermissionsResponse, error)
	FeedVersionAddPermission(context.Context, *FeedVersionModifyPermissionRequest) (*FeedVersionSaveResponse, error)
	FeedVersionRemovePermission(context.Context, *FeedVersionModifyPermissionRequest) (*FeedVersionSaveResponse, error)
	mustEmbedUnimplementedCheckerServer()
}

// UnimplementedCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedCheckerServer struct {
}

func (UnimplementedCheckerServer) UserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedCheckerServer) User(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedCheckerServer) TenantList(context.Context, *TenantListRequest) (*TenantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantList not implemented")
}
func (UnimplementedCheckerServer) Tenant(context.Context, *TenantRequest) (*TenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tenant not implemented")
}
func (UnimplementedCheckerServer) TenantPermissions(context.Context, *TenantRequest) (*TenantPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantPermissions not implemented")
}
func (UnimplementedCheckerServer) TenantSave(context.Context, *TenantSaveRequest) (*TenantSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantSave not implemented")
}
func (UnimplementedCheckerServer) TenantAddPermission(context.Context, *TenantModifyPermissionRequest) (*TenantSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantAddPermission not implemented")
}
func (UnimplementedCheckerServer) TenantRemovePermission(context.Context, *TenantModifyPermissionRequest) (*TenantSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantRemovePermission not implemented")
}
func (UnimplementedCheckerServer) TenantCreate(context.Context, *TenantCreateRequest) (*TenantSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantCreate not implemented")
}
func (UnimplementedCheckerServer) TenantCreateGroup(context.Context, *TenantCreateGroupRequest) (*GroupSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenantCreateGroup not implemented")
}
func (UnimplementedCheckerServer) GroupList(context.Context, *GroupListRequest) (*GroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupList not implemented")
}
func (UnimplementedCheckerServer) Group(context.Context, *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Group not implemented")
}
func (UnimplementedCheckerServer) GroupPermissions(context.Context, *GroupRequest) (*GroupPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPermissions not implemented")
}
func (UnimplementedCheckerServer) GroupSave(context.Context, *GroupSaveRequest) (*GroupSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupSave not implemented")
}
func (UnimplementedCheckerServer) GroupAddPermission(context.Context, *GroupModifyPermissionRequest) (*GroupSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupAddPermission not implemented")
}
func (UnimplementedCheckerServer) GroupRemovePermission(context.Context, *GroupModifyPermissionRequest) (*GroupSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupRemovePermission not implemented")
}
func (UnimplementedCheckerServer) GroupSetTenant(context.Context, *GroupSetTenantRequest) (*GroupSetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupSetTenant not implemented")
}
func (UnimplementedCheckerServer) FeedList(context.Context, *FeedListRequest) (*FeedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedList not implemented")
}
func (UnimplementedCheckerServer) Feed(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedCheckerServer) FeedPermissions(context.Context, *FeedRequest) (*FeedPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedPermissions not implemented")
}
func (UnimplementedCheckerServer) FeedSetGroup(context.Context, *FeedSetGroupRequest) (*FeedSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedSetGroup not implemented")
}
func (UnimplementedCheckerServer) FeedVersionList(context.Context, *FeedVersionListRequest) (*FeedVersionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVersionList not implemented")
}
func (UnimplementedCheckerServer) FeedVersion(context.Context, *FeedVersionRequest) (*FeedVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVersion not implemented")
}
func (UnimplementedCheckerServer) FeedVersionPermissions(context.Context, *FeedVersionRequest) (*FeedVersionPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVersionPermissions not implemented")
}
func (UnimplementedCheckerServer) FeedVersionAddPermission(context.Context, *FeedVersionModifyPermissionRequest) (*FeedVersionSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVersionAddPermission not implemented")
}
func (UnimplementedCheckerServer) FeedVersionRemovePermission(context.Context, *FeedVersionModifyPermissionRequest) (*FeedVersionSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedVersionRemovePermission not implemented")
}
func (UnimplementedCheckerServer) mustEmbedUnimplementedCheckerServer() {}

// UnsafeCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckerServer will
// result in compilation errors.
type UnsafeCheckerServer interface {
	mustEmbedUnimplementedCheckerServer()
}

func RegisterCheckerServer(s grpc.ServiceRegistrar, srv CheckerServer) {
	s.RegisterService(&Checker_ServiceDesc, srv)
}

func _Checker_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).User(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantList(ctx, req.(*TenantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_Tenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Tenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/Tenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Tenant(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantPermissions(ctx, req.(*TenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantSave(ctx, req.(*TenantSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantAddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantAddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantAddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantAddPermission(ctx, req.(*TenantModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantRemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantRemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantRemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantRemovePermission(ctx, req.(*TenantModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantCreate(ctx, req.(*TenantCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_TenantCreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantCreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).TenantCreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/TenantCreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).TenantCreateGroup(ctx, req.(*TenantCreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupList(ctx, req.(*GroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_Group_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Group(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/Group",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Group(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupPermissions(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupSave(ctx, req.(*GroupSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupAddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupAddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupAddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupAddPermission(ctx, req.(*GroupModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupRemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupRemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupRemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupRemovePermission(ctx, req.(*GroupModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_GroupSetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupSetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).GroupSetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/GroupSetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).GroupSetTenant(ctx, req.(*GroupSetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedList(ctx, req.(*FeedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Feed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedPermissions(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedSetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedSetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedSetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedSetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedSetGroup(ctx, req.(*FeedSetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedVersionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVersionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedVersionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedVersionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedVersionList(ctx, req.(*FeedVersionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedVersion(ctx, req.(*FeedVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedVersionPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedVersionPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedVersionPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedVersionPermissions(ctx, req.(*FeedVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedVersionAddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVersionModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedVersionAddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedVersionAddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedVersionAddPermission(ctx, req.(*FeedVersionModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_FeedVersionRemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedVersionModifyPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).FeedVersionRemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azpb.Checker/FeedVersionRemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).FeedVersionRemovePermission(ctx, req.(*FeedVersionModifyPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Checker_ServiceDesc is the grpc.ServiceDesc for Checker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "azpb.Checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserList",
			Handler:    _Checker_UserList_Handler,
		},
		{
			MethodName: "User",
			Handler:    _Checker_User_Handler,
		},
		{
			MethodName: "TenantList",
			Handler:    _Checker_TenantList_Handler,
		},
		{
			MethodName: "Tenant",
			Handler:    _Checker_Tenant_Handler,
		},
		{
			MethodName: "TenantPermissions",
			Handler:    _Checker_TenantPermissions_Handler,
		},
		{
			MethodName: "TenantSave",
			Handler:    _Checker_TenantSave_Handler,
		},
		{
			MethodName: "TenantAddPermission",
			Handler:    _Checker_TenantAddPermission_Handler,
		},
		{
			MethodName: "TenantRemovePermission",
			Handler:    _Checker_TenantRemovePermission_Handler,
		},
		{
			MethodName: "TenantCreate",
			Handler:    _Checker_TenantCreate_Handler,
		},
		{
			MethodName: "TenantCreateGroup",
			Handler:    _Checker_TenantCreateGroup_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Checker_GroupList_Handler,
		},
		{
			MethodName: "Group",
			Handler:    _Checker_Group_Handler,
		},
		{
			MethodName: "GroupPermissions",
			Handler:    _Checker_GroupPermissions_Handler,
		},
		{
			MethodName: "GroupSave",
			Handler:    _Checker_GroupSave_Handler,
		},
		{
			MethodName: "GroupAddPermission",
			Handler:    _Checker_GroupAddPermission_Handler,
		},
		{
			MethodName: "GroupRemovePermission",
			Handler:    _Checker_GroupRemovePermission_Handler,
		},
		{
			MethodName: "GroupSetTenant",
			Handler:    _Checker_GroupSetTenant_Handler,
		},
		{
			MethodName: "FeedList",
			Handler:    _Checker_FeedList_Handler,
		},
		{
			MethodName: "Feed",
			Handler:    _Checker_Feed_Handler,
		},
		{
			MethodName: "FeedPermissions",
			Handler:    _Checker_FeedPermissions_Handler,
		},
		{
			MethodName: "FeedSetGroup",
			Handler:    _Checker_FeedSetGroup_Handler,
		},
		{
			MethodName: "FeedVersionList",
			Handler:    _Checker_FeedVersionList_Handler,
		},
		{
			MethodName: "FeedVersion",
			Handler:    _Checker_FeedVersion_Handler,
		},
		{
			MethodName: "FeedVersionPermissions",
			Handler:    _Checker_FeedVersionPermissions_Handler,
		},
		{
			MethodName: "FeedVersionAddPermission",
			Handler:    _Checker_FeedVersionAddPermission_Handler,
		},
		{
			MethodName: "FeedVersionRemovePermission",
			Handler:    _Checker_FeedVersionRemovePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checker.proto",
}
