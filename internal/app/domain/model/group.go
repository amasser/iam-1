package model

import (
	"github.com/maurofran/kit/assert"
)

// GroupRepository is the interface representing the virtual collection of groups.
type GroupRepository interface {
	Add(*Group) error
	Update(*Group) error
	Remove(*Group) error
	AllGroups(TenantID) ([]*Group, error)
	GroupNamed(TenantID, string) (*Group, error)
}

// Group is the aggregate root for groups.
type Group struct {
	TenantID    TenantID
	Name        string
	Description string
	Members     []GroupMember
}

// NewGroup create a new instance of group aggregate root.
func NewGroup(tenantID TenantID, name, description string) (*Group, error) {
	if err := assert.NotZero(tenantID, "tenantID"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(name, "name"); err != nil {
		return nil, err
	}
	return &Group{tenantID, name, description, nil}, nil
}

// AddGroup is used to add the supplied group as a member of this one.
// The function call can fail either if the supplied group does not belongs to the
// tenant of receiver group or if a group recursion is detected.
func (g *Group) AddGroup(other *Group, memberService GroupMemberService) error {
	if err := assert.Equals(other.TenantID, g.TenantID, "other.TenantID"); err != nil {
		return err
	}
	member, err := memberService.IsMemberGroup(other, g.toGroupMember())
	if err != nil {
		return err
	}
	if err := assert.Condition(!member, "group recursion detected"); err != nil {
		return err
	}
	g.Members = append(g.Members, other.toGroupMember())
	return nil
}

// AddUser is used to add the supplied user as a member of this group.
// The function call can fail either if the user is not enabled or if the user does not
// belongs to the tenant of the receiver group.
func (g *Group) AddUser(user *User) error {
	if err := assert.Equals(user.TenantID, g.TenantID, "user.TenantID"); err != nil {
		return err
	}
	if err := assert.IsTrue(user.Enabled(), "user.Enabled"); err != nil {
		return err
	}
	g.Members = append(g.Members, user.toGroupMember())
	return nil
}

// IsMember will check if supplied user is member of this group (either directly or through a nested group).
// The function call can fail either if the user is not enabled or if the user does not
// belongs to the tenant of the receiver group.
func (g *Group) IsMember(user *User, memberService GroupMemberService) (bool, error) {
	if err := assert.Equals(user.TenantID, g.TenantID, "user.TenantID"); err != nil {
		return false, err
	}
	if err := assert.IsTrue(user.Enabled(), "user.Enabled"); err != nil {
		return false, err
	}
	userMember := user.toGroupMember()
	for _, member := range g.Members {
		if member == userMember {
			return memberService.ConfirmUser(g, user)
		}
	}
	return memberService.IsUserInNestedGroup(g, user)
}

// RemoveGroup will remove the given group from the receiver one.
// The function call can fail either if the supplied group does not belongs to the
// tenant of receiver group.
func (g *Group) RemoveGroup(other *Group) error {
	if err := assert.Equals(other.TenantID, g.TenantID, "other.TenantID"); err != nil {
		return err
	}
	groupMember := other.toGroupMember()
	for i, member := range g.Members {
		if member == groupMember {
			var m []GroupMember
			copy(m, g.Members)
			m[i] = m[len(m)-1]
			m = m[:len(m)-1]
			g.Members = m
			return nil
		}
	}
	return nil
}

// RemoveUser will remove the given user from the receiver group.
// The function call fail if the tenant of the user is not the same of the receiver group.
func (g *Group) RemoveUser(user *User) error {
	if err := assert.Equals(user.TenantID, g.TenantID, "user.TenantID"); err != nil {
		return err
	}
	userMember := user.toGroupMember()
	for i, member := range g.Members {
		if member == userMember {
			var m []GroupMember
			copy(m, g.Members)
			m[i] = m[len(m)-1]
			m = m[:len(m)-1]
			g.Members = m
			return nil
		}
	}
	return nil
}

func (g *Group) toGroupMember() GroupMember {
	return GroupMember{MemberTypeGroup, g.Name}
}
