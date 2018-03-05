package grpc

import (
	"context"

	"github.com/maurofran/iam/internal/app/application/command"

	"github.com/maurofran/iam/internal/app/application"
)

// RoleServer will provide the GRPC server used to manager roles.
type RoleServer struct {
	RoleService *application.RoleService `inject=""`
}

// ProvisionRole will provide the provisioning of role through GRPC
func (rs *RoleServer) ProvisionRole(ctx context.Context, req *ProvisionRoleRequest) (*ProvisionRoleResponse, error) {
	cmd := command.ProvisionRole{
		TenantID:        req.TenantId,
		RoleName:        req.Name,
		Description:     req.Description,
		SupportsNesting: req.SupportsNesting,
	}
	if err := rs.RoleService.ProvisionRole(ctx, cmd); err != nil {
		return nil, err
	}
	return &ProvisionRoleResponse{}, nil
}

// AssignUserToRole will provide the assign user to role through GRPC
func (rs *RoleServer) AssignUserToRole(ctx context.Context, req *AssignUserToRoleRequest) (*AssignUserToRoleResponse, error) {
	cmd := command.AssignUserToRole{
		TenantID: req.TenantId,
		RoleName: req.RoleName,
		Username: req.Username,
	}
	if err := rs.RoleService.AssignUserToRole(ctx, cmd); err != nil {
		return nil, err
	}
	return &AssignUserToRoleResponse{}, nil
}

// UnassignUserFromRole will provide the unassign user from role through GRPC
func (rs *RoleServer) UnassignUserFromRole(ctx context.Context, req *UnassignUserFromRoleRequest) (*UnassignUserFromRoleResponse, error) {
	cmd := command.UnassignUserFromRole{
		TenantID: req.TenantId,
		RoleName: req.RoleName,
		Username: req.Username,
	}
	if err := rs.RoleService.UnassignUserFromRole(ctx, cmd); err != nil {
		return nil, err
	}
	return &UnassignUserFromRoleResponse{}, nil
}

// AssignGroupToRole will provide the assign group to role through GRPC
func (rs *RoleServer) AssignGroupToRole(ctx context.Context, req *AssignGroupToRoleRequest) (*AssignGroupToRoleResponse, error) {
	cmd := command.AssignGroupToRole{
		TenantID:  req.TenantId,
		RoleName:  req.RoleName,
		GroupName: req.GroupName,
	}
	if err := rs.RoleService.AssignGroupToRole(ctx, cmd); err != nil {
		return nil, err
	}
	return &AssignGroupToRoleResponse{}, nil
}

// UnassignGroupFromRole will provide the unassign group from role throguh GRPC
func (rs *RoleServer) UnassignGroupFromRole(ctx context.Context, req *UnassignGroupFromRoleRequest) (*UnassignGroupFromRoleResponse, error) {
	cmd := command.UnassignGroupFromRole{
		TenantID:  req.TenantId,
		RoleName:  req.RoleName,
		GroupName: req.GroupName,
	}
	if err := rs.RoleService.UnassignGroupFromRole(ctx, cmd); err != nil {
		return nil, err
	}
	return &UnassignGroupFromRoleResponse{}, nil
}
