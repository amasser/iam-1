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
type GroupMemberService struct {
	GroupRepo GroupRepository
	UserRepo  UserRepository
}

// ConfirmUser check that user belongs to same tenant of group and it's active.
func (gms *GroupMemberService) ConfirmUser(group *Group, user *User) (bool, error) {
	confirmed, err := gms.UserRepo.UserWithUsername(group.TenantID, user.Username)
	if err != nil {
		return false, err
	}
	return confirmed != nil && confirmed.Enabled(), nil
}

// IsMemberGroup check if supplied member is a member of supplied group
func (gms *GroupMemberService) IsMemberGroup(group *Group, member GroupMember) (bool, error) {
	for _, m := range group.Members {
		if m.IsGroup() {
			if m == member {
				return true, nil
			}
			nestedGroup, err := gms.GroupRepo.GroupNamed(group.TenantID, m.Name)
			if err != nil {
				return false, err
			}
			if nestedGroup != nil {
				isMember, err := gms.IsMemberGroup(nestedGroup, member)
				if err != nil {
					return false, err
				}
				if isMember {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

// IsUserInNestedGroup will check that supplied user is in a nested group of supplied group.
func (gms *GroupMemberService) IsUserInNestedGroup(group *Group, user *User) (bool, error) {
	for _, m := range group.Members {
		if m.IsGroup() {
			nestedGroup, err := gms.GroupRepo.GroupNamed(group.TenantID, m.Name)
			if err != nil {
				return false, err
			}
			if nestedGroup != nil {
				isMember, err := nestedGroup.IsMember(user, gms)
				if err != nil {
					return false, err
				}
				if isMember {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
