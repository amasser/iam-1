package command

// ProvisionGroup will provision a new group.
type ProvisionGroup struct {
	TenantID    string `json:"tenantId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// AddGroupToGroup will add a group to another group.
type AddGroupToGroup struct {
	TenantID       string `json:"tenantId"`
	GroupName      string `json:"groupName"`
	ChildGroupName string `json:"childGroupName"`
}

// AddUserToGroup will add a user to a group.
type AddUserToGroup struct {
	TenantID  string `json:"tenantId"`
	GroupName string `json:"groupName"`
	Username  string `json:"username"`
}

// RemoveGroupFromGroup will remove group from another group.
type RemoveGroupFromGroup struct {
	TenantID       string `json:"tenantId"`
	GroupName      string `json:"groupName"`
	ChildGroupName string `json:"childGroupName"`
}

// RemoveUserFromGroup will remove a user from a group.
type RemoveUserFromGroup struct {
	TenantID  string `json:"tenantId"`
	GroupName string `json:"groupName"`
	Username  string `json:"username"`
}
