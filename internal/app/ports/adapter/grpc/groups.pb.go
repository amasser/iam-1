// Code generated by protoc-gen-go. DO NOT EDIT.
// source: groups.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	groups.proto
	roles.proto
	tenants.proto

It has these top-level messages:
	ProvisionGroupRequest
	ProvisionGroupResponse
	AddGroupToGroupRequest
	AddGroupToGroupResponse
	RemoveGroupFromGroupRequest
	RemoveGroupFromGroupResponse
	AddUserToGroupRequest
	AddUserToGroupResponse
	RemoveUserFromGroupRequest
	RemoveUserFromGroupResponse
	ProvisionRoleRequest
	ProvisionRoleResponse
	AssignUserToRoleRequest
	AssignUserToRoleResponse
	UnassignUserFromRoleRequest
	UnassignUserFromRoleResponse
	AssignGroupToRoleRequest
	AssignGroupToRoleResponse
	UnassignGroupFromRoleRequest
	UnassignGroupFromRoleResponse
	ProvisionTenantRequest
	ProvisionTenantResponse
	ActivateTenantRequest
	ActivateTenantResponse
	DeactivateTenantRequest
	DeactivateTenantResponse
	OfferInvitationRequest
	OfferInvitationResponse
	WithdrawInvitationRequest
	WithdrawInvitationResponse
	FullName
	PostalAddress
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProvisionGroupRequest struct {
	TenantId    string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
	GroupName   string `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *ProvisionGroupRequest) Reset()                    { *m = ProvisionGroupRequest{} }
func (m *ProvisionGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*ProvisionGroupRequest) ProtoMessage()               {}
func (*ProvisionGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProvisionGroupRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *ProvisionGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *ProvisionGroupRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ProvisionGroupResponse struct {
}

func (m *ProvisionGroupResponse) Reset()                    { *m = ProvisionGroupResponse{} }
func (m *ProvisionGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*ProvisionGroupResponse) ProtoMessage()               {}
func (*ProvisionGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type AddGroupToGroupRequest struct {
	TenantId        string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
	GroupName       string `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	NestedGroupName string `protobuf:"bytes,3,opt,name=nested_group_name,json=nestedGroupName" json:"nested_group_name,omitempty"`
}

func (m *AddGroupToGroupRequest) Reset()                    { *m = AddGroupToGroupRequest{} }
func (m *AddGroupToGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*AddGroupToGroupRequest) ProtoMessage()               {}
func (*AddGroupToGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddGroupToGroupRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *AddGroupToGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *AddGroupToGroupRequest) GetNestedGroupName() string {
	if m != nil {
		return m.NestedGroupName
	}
	return ""
}

type AddGroupToGroupResponse struct {
}

func (m *AddGroupToGroupResponse) Reset()                    { *m = AddGroupToGroupResponse{} }
func (m *AddGroupToGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*AddGroupToGroupResponse) ProtoMessage()               {}
func (*AddGroupToGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RemoveGroupFromGroupRequest struct {
	TenantId        string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
	GroupName       string `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	NestedGroupName string `protobuf:"bytes,3,opt,name=nested_group_name,json=nestedGroupName" json:"nested_group_name,omitempty"`
}

func (m *RemoveGroupFromGroupRequest) Reset()                    { *m = RemoveGroupFromGroupRequest{} }
func (m *RemoveGroupFromGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveGroupFromGroupRequest) ProtoMessage()               {}
func (*RemoveGroupFromGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveGroupFromGroupRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *RemoveGroupFromGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *RemoveGroupFromGroupRequest) GetNestedGroupName() string {
	if m != nil {
		return m.NestedGroupName
	}
	return ""
}

type RemoveGroupFromGroupResponse struct {
}

func (m *RemoveGroupFromGroupResponse) Reset()                    { *m = RemoveGroupFromGroupResponse{} }
func (m *RemoveGroupFromGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveGroupFromGroupResponse) ProtoMessage()               {}
func (*RemoveGroupFromGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AddUserToGroupRequest struct {
	TenantId  string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *AddUserToGroupRequest) Reset()                    { *m = AddUserToGroupRequest{} }
func (m *AddUserToGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserToGroupRequest) ProtoMessage()               {}
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddUserToGroupRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *AddUserToGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *AddUserToGroupRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AddUserToGroupResponse struct {
}

func (m *AddUserToGroupResponse) Reset()                    { *m = AddUserToGroupResponse{} }
func (m *AddUserToGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserToGroupResponse) ProtoMessage()               {}
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type RemoveUserFromGroupRequest struct {
	TenantId  string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId" json:"tenant_id,omitempty"`
	GroupName string `protobuf:"bytes,2,opt,name=group_name,json=groupName" json:"group_name,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *RemoveUserFromGroupRequest) Reset()                    { *m = RemoveUserFromGroupRequest{} }
func (m *RemoveUserFromGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveUserFromGroupRequest) ProtoMessage()               {}
func (*RemoveUserFromGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoveUserFromGroupRequest) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *RemoveUserFromGroupRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *RemoveUserFromGroupRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RemoveUserFromGroupResponse struct {
}

func (m *RemoveUserFromGroupResponse) Reset()                    { *m = RemoveUserFromGroupResponse{} }
func (m *RemoveUserFromGroupResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveUserFromGroupResponse) ProtoMessage()               {}
func (*RemoveUserFromGroupResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*ProvisionGroupRequest)(nil), "iam.ProvisionGroupRequest")
	proto.RegisterType((*ProvisionGroupResponse)(nil), "iam.ProvisionGroupResponse")
	proto.RegisterType((*AddGroupToGroupRequest)(nil), "iam.AddGroupToGroupRequest")
	proto.RegisterType((*AddGroupToGroupResponse)(nil), "iam.AddGroupToGroupResponse")
	proto.RegisterType((*RemoveGroupFromGroupRequest)(nil), "iam.RemoveGroupFromGroupRequest")
	proto.RegisterType((*RemoveGroupFromGroupResponse)(nil), "iam.RemoveGroupFromGroupResponse")
	proto.RegisterType((*AddUserToGroupRequest)(nil), "iam.AddUserToGroupRequest")
	proto.RegisterType((*AddUserToGroupResponse)(nil), "iam.AddUserToGroupResponse")
	proto.RegisterType((*RemoveUserFromGroupRequest)(nil), "iam.RemoveUserFromGroupRequest")
	proto.RegisterType((*RemoveUserFromGroupResponse)(nil), "iam.RemoveUserFromGroupResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for GroupService service

type GroupServiceClient interface {
	ProvisionGroup(ctx context.Context, in *ProvisionGroupRequest, opts ...grpc1.CallOption) (*ProvisionGroupResponse, error)
	AddGroupToGroup(ctx context.Context, in *AddGroupToGroupRequest, opts ...grpc1.CallOption) (*AddGroupToGroupResponse, error)
	RemoveGroupFromGroup(ctx context.Context, in *RemoveGroupFromGroupRequest, opts ...grpc1.CallOption) (*RemoveGroupFromGroupResponse, error)
	AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc1.CallOption) (*AddUserToGroupResponse, error)
	RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc1.CallOption) (*RemoveUserFromGroupResponse, error)
}

type groupServiceClient struct {
	cc *grpc1.ClientConn
}

func NewGroupServiceClient(cc *grpc1.ClientConn) GroupServiceClient {
	return &groupServiceClient{cc}
}

func (c *groupServiceClient) ProvisionGroup(ctx context.Context, in *ProvisionGroupRequest, opts ...grpc1.CallOption) (*ProvisionGroupResponse, error) {
	out := new(ProvisionGroupResponse)
	err := grpc1.Invoke(ctx, "/iam.GroupService/ProvisionGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddGroupToGroup(ctx context.Context, in *AddGroupToGroupRequest, opts ...grpc1.CallOption) (*AddGroupToGroupResponse, error) {
	out := new(AddGroupToGroupResponse)
	err := grpc1.Invoke(ctx, "/iam.GroupService/AddGroupToGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) RemoveGroupFromGroup(ctx context.Context, in *RemoveGroupFromGroupRequest, opts ...grpc1.CallOption) (*RemoveGroupFromGroupResponse, error) {
	out := new(RemoveGroupFromGroupResponse)
	err := grpc1.Invoke(ctx, "/iam.GroupService/RemoveGroupFromGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc1.CallOption) (*AddUserToGroupResponse, error) {
	out := new(AddUserToGroupResponse)
	err := grpc1.Invoke(ctx, "/iam.GroupService/AddUserToGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc1.CallOption) (*RemoveUserFromGroupResponse, error) {
	out := new(RemoveUserFromGroupResponse)
	err := grpc1.Invoke(ctx, "/iam.GroupService/RemoveUserFromGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceServer interface {
	ProvisionGroup(context.Context, *ProvisionGroupRequest) (*ProvisionGroupResponse, error)
	AddGroupToGroup(context.Context, *AddGroupToGroupRequest) (*AddGroupToGroupResponse, error)
	RemoveGroupFromGroup(context.Context, *RemoveGroupFromGroupRequest) (*RemoveGroupFromGroupResponse, error)
	AddUserToGroup(context.Context, *AddUserToGroupRequest) (*AddUserToGroupResponse, error)
	RemoveUserFromGroup(context.Context, *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error)
}

func RegisterGroupServiceServer(s *grpc1.Server, srv GroupServiceServer) {
	s.RegisterService(&_GroupService_serviceDesc, srv)
}

func _GroupService_ProvisionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).ProvisionGroup(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.GroupService/ProvisionGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).ProvisionGroup(ctx, req.(*ProvisionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddGroupToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddGroupToGroup(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.GroupService/AddGroupToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddGroupToGroup(ctx, req.(*AddGroupToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_RemoveGroupFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).RemoveGroupFromGroup(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.GroupService/RemoveGroupFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).RemoveGroupFromGroup(ctx, req.(*RemoveGroupFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_AddUserToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).AddUserToGroup(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.GroupService/AddUserToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).AddUserToGroup(ctx, req.(*AddUserToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupService_RemoveUserFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServiceServer).RemoveUserFromGroup(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iam.GroupService/RemoveUserFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServiceServer).RemoveUserFromGroup(ctx, req.(*RemoveUserFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "iam.GroupService",
	HandlerType: (*GroupServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "ProvisionGroup",
			Handler:    _GroupService_ProvisionGroup_Handler,
		},
		{
			MethodName: "AddGroupToGroup",
			Handler:    _GroupService_AddGroupToGroup_Handler,
		},
		{
			MethodName: "RemoveGroupFromGroup",
			Handler:    _GroupService_RemoveGroupFromGroup_Handler,
		},
		{
			MethodName: "AddUserToGroup",
			Handler:    _GroupService_AddUserToGroup_Handler,
		},
		{
			MethodName: "RemoveUserFromGroup",
			Handler:    _GroupService_RemoveUserFromGroup_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "groups.proto",
}

func init() { proto.RegisterFile("groups.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x6f, 0x2f, 0x37, 0x04, 0xce, 0x25, 0x90, 0x3b, 0x57, 0xb0, 0x4e, 0x41, 0x6b, 0x57,
	0xc6, 0x05, 0x0b, 0x7d, 0x02, 0x5c, 0x48, 0x0c, 0x09, 0x31, 0xa8, 0x1b, 0x12, 0x43, 0x2a, 0x73,
	0x42, 0x66, 0xd1, 0x4e, 0x9d, 0x99, 0xb2, 0x76, 0xe3, 0x13, 0xf9, 0x82, 0x86, 0x69, 0x69, 0x4a,
	0x1d, 0x59, 0x91, 0xb8, 0xe4, 0xff, 0xcf, 0x9c, 0xf9, 0xe6, 0x9c, 0x9f, 0x42, 0x6b, 0x25, 0x45,
	0x9a, 0xa8, 0x61, 0x22, 0x85, 0x16, 0xa4, 0xc6, 0xc3, 0x28, 0x48, 0xa1, 0x7b, 0x2f, 0xc5, 0x9a,
	0x2b, 0x2e, 0xe2, 0xf1, 0xc6, 0x9d, 0xe1, 0x6b, 0x8a, 0x4a, 0x13, 0x0f, 0x9a, 0x1a, 0xe3, 0x30,
	0xd6, 0x0b, 0xce, 0x5c, 0xc7, 0x77, 0x2e, 0x9a, 0xb3, 0x46, 0x26, 0xdc, 0x31, 0x32, 0x00, 0x30,
	0xad, 0x16, 0x71, 0x18, 0xa1, 0xfb, 0xdb, 0xb8, 0x4d, 0xa3, 0x4c, 0xc3, 0x08, 0x89, 0x0f, 0x7f,
	0x19, 0xaa, 0xa5, 0xe4, 0x89, 0xe6, 0x22, 0x76, 0x6b, 0xc6, 0x2f, 0x4b, 0x81, 0x0b, 0xbd, 0xea,
	0xb5, 0x2a, 0x11, 0xb1, 0xc2, 0xe0, 0xcd, 0x81, 0xde, 0x88, 0x31, 0x23, 0x3e, 0x8a, 0x83, 0x21,
	0x5d, 0xc2, 0xbf, 0x18, 0x95, 0x46, 0xb6, 0x28, 0x55, 0x65, 0x60, 0x9d, 0xcc, 0x18, 0x6f, 0x6b,
	0x83, 0x13, 0x38, 0xfe, 0x42, 0x90, 0xd3, 0xbd, 0x3b, 0xe0, 0xcd, 0x30, 0x12, 0x6b, 0x34, 0xfa,
	0xad, 0x14, 0xd1, 0x8f, 0x20, 0x9e, 0x42, 0xdf, 0x8e, 0x91, 0x73, 0x0a, 0xe8, 0x8e, 0x18, 0x7b,
	0x52, 0x28, 0x0f, 0x38, 0x43, 0x0a, 0x8d, 0x54, 0xa1, 0x2c, 0x71, 0x15, 0xbf, 0x37, 0x0b, 0xad,
	0x5e, 0x98, 0xa3, 0x68, 0xa0, 0x19, 0xea, 0xc6, 0x3c, 0xe8, 0xc0, 0xf6, 0xf1, 0x0c, 0xb6, 0x7b,
	0xaa, 0xdc, 0x9a, 0x41, 0x5d, 0x7d, 0xd4, 0xa0, 0x65, 0x94, 0x07, 0x94, 0x6b, 0xbe, 0x44, 0x32,
	0x81, 0xf6, 0x6e, 0x20, 0x09, 0x1d, 0xf2, 0x30, 0x1a, 0x5a, 0xff, 0x1c, 0xd4, 0xb3, 0x7a, 0xf9,
	0x83, 0x7f, 0x91, 0x29, 0x74, 0x2a, 0x01, 0x22, 0xd9, 0x09, 0x7b, 0xb0, 0x69, 0xdf, 0x6e, 0x16,
	0xfd, 0x9e, 0xe1, 0xc8, 0xb6, 0x6d, 0xe2, 0x9b, 0x73, 0x7b, 0xf2, 0x48, 0xcf, 0xf7, 0x54, 0x14,
	0xed, 0x27, 0xd0, 0xde, 0xdd, 0x5d, 0xfe, 0x76, 0x6b, 0x82, 0xa8, 0x67, 0xf5, 0x8a, 0x66, 0x73,
	0xf8, 0x6f, 0x19, 0x3c, 0x39, 0x2b, 0x81, 0xd8, 0x82, 0x40, 0xfd, 0xef, 0x0b, 0xb6, 0xbd, 0x6f,
	0xea, 0xf3, 0x3f, 0x2b, 0x99, 0x2c, 0x5f, 0xea, 0xe6, 0x03, 0x76, 0xfd, 0x19, 0x00, 0x00, 0xff,
	0xff, 0x83, 0x71, 0x27, 0x61, 0xd0, 0x04, 0x00, 0x00,
}
