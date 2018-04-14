package iam

// Role is the aggregate root object managing roles.
type Role struct {
	TenantID        TenantID
	Name            string
	Description     string
	SupportsNesting bool
	Group           *Group
}

// Roles is the collection of roles
type Roles []*Role

// RoleRepository is the repository of roles.
type RoleRepository interface {
	Add(*Role) error
	Update(*Role) error
	Remove(*Role) error
	RoleNamed(TenantID, string) (*Role, error)
	AllRoles(TenantID) (Roles, error)
}
