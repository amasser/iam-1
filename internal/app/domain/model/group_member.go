package model

// MemberTypeGroup is the group member type constant for a group.
// MemberTypeUser is the group member type constant for a user.
const (
	MemberTypeGroup GroupMemberType = 1 + iota
	MemberTypeUser
)

var memberTypes = [...]string{"Group", "User"}

// GroupMemberType is the value of group member types.
type GroupMemberType int

// IsZero will check if receiver group member type is zero value.
func (gmt GroupMemberType) IsZero() bool {
	return gmt == 0
}

func (gmt GroupMemberType) String() string {
	return memberTypes[gmt-1]
}

// GroupMember is the value object used to represent a group membership.
//
// The default value of group member is not a valid value.
type GroupMember struct {
	Type GroupMemberType
	Name string
}

// IsGroup will check if the member is a group.
func (gm GroupMember) IsGroup() bool {
	return gm.Type == MemberTypeGroup
}

// IsUser will check if the number is a user.
func (gm GroupMember) IsUser() bool {
	return gm.Type == MemberTypeUser
}

// GroupMemberService is the interface for domain service used to confirm and verify group membership.
type GroupMemberService interface {
	ConfirmUser(*Group, *User) (bool, error)
	IsMemberGroup(*Group, GroupMember) (bool, error)
	IsUserInNestedGroup(*Group, *User) (bool, error)
}
