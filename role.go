package iam

// Role is the aggregate root object managing roles.
type Role struct {
	TenantID        TenantID
	Name            string
	Description     string
	SupportsNesting bool
	Group           *Group
}
