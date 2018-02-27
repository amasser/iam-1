package command

// ProvisionRole will provision a new role.
type ProvisionRole struct {
	TenantID        string `json:"tenantId"`
	RoleName        string `json:"roleName"`
	Description     string `json:"description"`
	SupportsNesting bool   `json:"supportsNesting"`
}

// AssignUserToRole will assign a user to a role.
type AssignUserToRole struct {
	TenantID string `json:"tenantId"`
	RoleName string `json:"roleName"`
	Username string `json:"username"`
}

// UnassignUserFromRole will remove a user from a role.
type UnassignUserFromRole struct {
	TenantID string `json:"tenantId"`
	RoleName string `json:"roleName"`
	Username string `json:"username"`
}

// AssignGroupToRole will assign a group to a role.
type AssignGroupToRole struct {
	TenantID  string `json:"tenantId"`
	RoleName  string `json:"roleName"`
	GroupName string `json:"groupName"`
}

// UnassignGroupFromRole will unassign a group from a role.
type UnassignGroupFromRole struct {
	TenantID  string `json:"tenantId"`
	RoleName  string `json:"roleName"`
	GroupName string `json:"groupName"`
}
