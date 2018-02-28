package application

import (
	"errors"

	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// ErrGroupNotFound is returned when a group was not found.
var ErrGroupNotFound = errors.New("no group found")

// GroupService is the application service used to manage groups.
type GroupService struct {
	TenantRepository   model.TenantRepository    `inject:""`
	GroupRepository    model.GroupRepository     `inject:""`
	UserRepository     model.UserRepository      `inject:""`
	GroupMemberService *model.GroupMemberService `inject:""`
}

// ProvisionGroup will provision a new group.
func (gs *GroupService) ProvisionGroup(cmd command.ProvisionGroup) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := gs.TenantRepository.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	group, err := tenant.ProvisionGroup(cmd.Name, cmd.Description)
	if err != nil {
		return err
	}
	return gs.GroupRepository.Add(group)
}

// AddGroupToGroup will add a group to an existing group.
func (gs *GroupService) AddGroupToGroup(cmd command.AddGroupToGroup) error {
	group, err := gs.loadGroup(cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	otherGroup, err := gs.loadGroup(cmd.TenantID, cmd.ChildGroupName)
	if err != nil {
		return err
	}
	if err := group.AddGroup(otherGroup, gs.GroupMemberService); err != nil {
		return err
	}
	return gs.GroupRepository.Update(group)
}

// AddUserToGroup will ad a user to an existing group.
func (gs *GroupService) AddUserToGroup(cmd command.AddUserToGroup) error {
	group, err := gs.loadGroup(cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	user, err := gs.loadUser(cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	if err := group.AddUser(user); err != nil {
		return err
	}
	return gs.GroupRepository.Update(group)
}

// RemoveGroupFromGroup will remove a group from an existing group.
func (gs *GroupService) RemoveGroupFromGroup(cmd command.RemoveGroupFromGroup) error {
	group, err := gs.loadGroup(cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	otherGroup, err := gs.loadGroup(cmd.TenantID, cmd.ChildGroupName)
	if err != nil {
		return err
	}
	if _, err := group.RemoveGroup(otherGroup); err != nil {
		return err
	}
	return gs.GroupRepository.Update(group)
}

// RemoveUserFromGroup will remove a user from a group.
func (gs *GroupService) RemoveUserFromGroup(cmd command.RemoveUserFromGroup) error {
	group, err := gs.loadGroup(cmd.TenantID, cmd.GroupName)
	if err != nil {
		return err
	}
	user, err := gs.loadUser(cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	if _, err := group.RemoveUser(user); err != nil {
		return err
	}
	return gs.GroupRepository.Update(group)
}

func (gs *GroupService) loadGroup(tenantID, name string) (*model.Group, error) {
	theTenantID, err := model.MakeTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	group, err := gs.GroupRepository.GroupNamed(theTenantID, name)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, ErrGroupNotFound
	}
	return group, nil
}

func (gs *GroupService) loadUser(tenantID, username string) (*model.User, error) {
	theTenantID, err := model.MakeTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	user, err := gs.UserRepository.UserWithUsername(theTenantID, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
