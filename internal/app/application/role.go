package application

import (
	"errors"

	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// ErrRoleNotFound is returned when no role was found
var ErrRoleNotFound = errors.New("role not found")

// RoleService is the service for roles.
type RoleService struct {
	TenantRepository   model.TenantRepository    `inject:""`
	RoleRepository     model.RoleRepository      `inject:""`
	GroupRepository    model.GroupRepository     `inject:""`
	UserRepository     model.UserRepository      `inject:""`
	GroupMemberService *model.GroupMemberService `inject:""`
}

// ProvisionRole will provision a new role.
func (rs *RoleService) ProvisionRole(cmd command.ProvisionRole) error {
	tenant, err := loadTenant(rs.TenantRepository, cmd.TenantID)
	if err != nil {
		return err
	}
	role, err := tenant.ProvisionRole(cmd.RoleName, cmd.Description, cmd.SupportsNesting)
	if err != nil {
		return err
	}
	return rs.RoleRepository.Add(role)
}

// AssignUserToRole will assign a user to a role.
func (rs *RoleService) AssignUserToRole(cmd command.AssignUserToRole) error {
	role, err := loadRole(rs.RoleRepository, cmd.TenantID, cmd.RoleName)
	if err != nil {
		return err
	}
	user, err := loadUser(rs.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	if err := role.AssignUser(user); err != nil {
		return err
	}
	return rs.RoleRepository.Update(role)
}

// UnassignUserFromRole will remove a user from a role.
func (rs *RoleService) UnassignUserFromRole(cmd command.UnassignUserFromRole) error {
	role, err := loadRole(rs.RoleRepository, cmd.TenantID, cmd.RoleName)
	if err != nil {
		return err
	}
	user, err := loadUser(rs.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	if err := role.UnassignUser(user); err != nil {
		return err
	}
	return rs.RoleRepository.Update(role)
}

// AssignGroupToRole will assign a group to a role.
func (rs *RoleService) AssignGroupToRole(cmd command.AssignGroupToRole) error {
	role, err := loadRole(rs.RoleRepository, cmd.TenantID, cmd.RoleName)
	if err != nil {
		return err
	}
	group, err := loadGroup(rs.GroupRepository, cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	if err := role.AssignGroup(group, rs.GroupMemberService); err != nil {
		return err
	}
	return rs.RoleRepository.Update(role)
}

// UnassignGroupFromRole will unassign a group from a role.
func (rs *RoleService) UnassignGroupFromRole(cmd command.UnassignGroupFromRole) error {
	role, err := loadRole(rs.RoleRepository, cmd.TenantID, cmd.RoleName)
	if err != nil {
		return err
	}
	group, err := loadGroup(rs.GroupRepository, cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	if err := role.UnassignGroup(group); err != nil {
		return err
	}
	return rs.RoleRepository.Update(role)
}

func loadRole(repo model.RoleRepository, tenantID, name string) (*model.Role, error) {
	theTenantID, err := model.MakeTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	role, err := repo.RoleNamed(theTenantID, name)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, ErrRoleNotFound
	}
	return role, nil
}
