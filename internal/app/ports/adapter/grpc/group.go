package grpc

import (
	"context"

	"github.com/maurofran/iam/internal/app/application/command"

	"github.com/maurofran/iam/internal/app/application"
)

// GroupServer is the GRPC interface to group server.
type GroupServer struct {
	GroupService *application.GroupService `inject:""`
}

// ProvisionGroup exposes the provision group function through GRPC
func (gs *GroupServer) ProvisionGroup(ctx context.Context, req *ProvisionGroupRequest) (*ProvisionGroupResponse, error) {
	cmd := command.ProvisionGroup{
		TenantID:    req.TenantId,
		Name:        req.GroupName,
		Description: req.Description,
	}
	if err := gs.GroupService.ProvisionGroup(ctx, cmd); err != nil {
		return nil, err
	}
	return &ProvisionGroupResponse{}, nil
}

// AddGroupToGroup exposes the add group to group function through GRPC.
func (gs *GroupServer) AddGroupToGroup(ctx context.Context, req *AddGroupToGroupRequest) (*AddGroupToGroupResponse, error) {
	cmd := command.AddGroupToGroup{
		TenantID:       req.TenantId,
		GroupName:      req.GroupName,
		ChildGroupName: req.NestedGroupName,
	}
	if err := gs.GroupService.AddGroupToGroup(ctx, cmd); err != nil {
		return nil, err
	}
	return &AddGroupToGroupResponse{}, nil
}

// RemoveGroupFromGroup exposes the remove group from group function through GRPC.
func (gs *GroupServer) RemoveGroupFromGroup(ctx context.Context, req *RemoveGroupFromGroupRequest) (*RemoveGroupFromGroupResponse, error) {
	cmd := command.RemoveGroupFromGroup{
		TenantID:       req.TenantId,
		GroupName:      req.GroupName,
		ChildGroupName: req.NestedGroupName,
	}
	if err := gs.GroupService.RemoveGroupFromGroup(ctx, cmd); err != nil {
		return nil, err
	}
	return &RemoveGroupFromGroupResponse{}, nil
}

// AddUserToGroup exposes the add user to group function through GRPC.
func (gs *GroupServer) AddUserToGroup(ctx context.Context, req *AddUserToGroupRequest) (*AddUserToGroupResponse, error) {
	cmd := command.AddUserToGroup{
		TenantID:  req.TenantId,
		GroupName: req.GroupName,
		Username:  req.Username,
	}
	if err := gs.GroupService.AddUserToGroup(ctx, cmd); err != nil {
		return nil, err
	}
	return &AddUserToGroupResponse{}, nil
}

// RemoveUserFromGroup exposes the remove user from group function through GRPC.
func (gs *GroupServer) RemoveUserFromGroup(ctx context.Context, req *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error) {
	cmd := command.RemoveUserFromGroup{
		TenantID:  req.TenantId,
		GroupName: req.GroupName,
		Username:  req.Username,
	}
	if err := gs.GroupService.RemoveUserFromGroup(ctx, cmd); err != nil {
		return nil, err
	}
	return &RemoveUserFromGroupResponse{}, nil
}
