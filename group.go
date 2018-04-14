package iam

// Group is the aggregate root object representing a group.
type Group struct {
	TenantID    TenantID
	Name        string
	Description string
	Members     GroupMembers
}

// Groups is a collection of group.
type Groups []*Group

// GroupRepository is the interface for group management repository.
type GroupRepository interface {
	Add(*Group) error
	Update(*Group) error
	Remove(*Group) error
	GroupNamed(TenantID, string) (*Group, error)
	AllGroups(TenantID) (Groups, error)
}

// GroupMemberType is an enum type for group member.
type GroupMemberType int

// GroupMember is the value object representing a group member.
type GroupMember struct {
	Type GroupMemberType
	Name string
}

// GroupMembers is the collection of group members
type GroupMembers []*GroupMember
