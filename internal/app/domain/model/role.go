package model

import (
	"time"

	"github.com/maurofran/iam/internal/app/domain/model/event"
	"github.com/maurofran/iam/internal/pkg/aggregate"
	"github.com/maurofran/kit/assert"
)

const roleGroupPrefix = "ROLE-INTERNAL-GROUP: "

// RoleRepository is the interface for object managing roles virtual collection.
type RoleRepository interface {
	Add(*Role) error
	Update(*Role) error
	Remove(*Role) error
	AllRoles(TenantID) ([]*Role, error)
	RoleNamed(TenantID, string) (*Role, error)
}

// Role is the aggregate root object for roles.
type Role struct {
	aggregate.Root
	TenantID        TenantID
	Name            string
	Description     string
	SupportsNesting bool
	Group           *Group
}

func newRole(tenantID TenantID, name, description string, supportsNesting bool) (*Role, error) {
	if err := assert.NotZero(tenantID, "tenantID"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(name, "name"); err != nil {
		return nil, err
	}
	group, err := newGroup(tenantID, roleGroupPrefix+name, "")
	if err != nil {
		return nil, err
	}
	role := &Role{TenantID: tenantID, Name: name, Description: description, SupportsNesting: supportsNesting, Group: group}
	role.RegisterEvent(&event.RoleProvisioned{
		EventVersion: 1,
		OccurredOn:   time.Now().Unix(),
		TenantId:     string(tenantID),
		RoleName:     name,
	})
	return role, nil
}

// AssignGroup will assign the supplied group to this role.
func (r *Role) AssignGroup(group *Group, memberService GroupMemberService) error {
	if err := assert.State(r.SupportsNesting, "role does not supports group nesting"); err != nil {
		return err
	}
	if err := assert.Equals(group.TenantID, r.TenantID, "group.TenantID"); err != nil {
		return err
	}
	if err := r.Group.AddGroup(group, memberService); err != nil {
		return err
	}
	r.RegisterEvent(&event.GroupAssignedToRole{
		EventVersion: 1,
		OccurredOn:   time.Now().Unix(),
		TenantId:     string(r.TenantID),
		RoleName:     r.Name,
		GroupName:    group.Name,
	})
	return nil
}

// AssignUser will assign the supplied user to this role.
func (r *Role) AssignUser(user *User) error {
	if err := assert.Equals(user.TenantID, r.TenantID, "user.TenantID"); err != nil {
		return err
	}
	if err := r.Group.AddUser(user); err != nil {
		return err
	}
	r.RegisterEvent(&event.UserAssignedToRole{
		EventVersion: 1,
		OccurredOn:   time.Now().Unix(),
		TenantId:     string(r.TenantID),
		RoleName:     r.Name,
		Username:     user.Username,
	})
	return nil
}

// IsInRole check if supplied user belongs to this role
func (r *Role) IsInRole(user *User, memberService GroupMemberService) (bool, error) {
	return r.Group.IsMember(user, memberService)
}

// UnassignGroup will unassign the group from role.
func (r *Role) UnassignGroup(group *Group) error {
	if err := assert.State(r.SupportsNesting, "role does not supports group nesting"); err != nil {
		return err
	}
	removed, err := r.Group.RemoveGroup(group)
	if err != nil && removed {
		r.RegisterEvent(&event.GroupUnassignedFromRole{
			EventVersion: 1,
			OccurredOn:   time.Now().Unix(),
			TenantId:     string(r.TenantID),
			RoleName:     r.Name,
			GroupName:    group.Name,
		})
	}
	return err
}

// UnassignUser will unassign the user from role.
func (r *Role) UnassignUser(user *User) error {
	removed, err := r.Group.RemoveUser(user)
	if err != nil && removed {
		r.RegisterEvent(&event.UserUnassignedFromRole{
			EventVersion: 1,
			OccurredOn:   time.Now().Unix(),
			TenantId:     string(r.TenantID),
			RoleName:     r.Name,
			Username:     user.Username,
		})
	}
	return err
}
