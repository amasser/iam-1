package iam

// Group is the aggregate root object representing a group.
type Group struct {
	TenantID    TenantID
	Name        string
	Description string
	Members     GroupMembers
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
