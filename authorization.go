package iam

// AuthorizationService is the interface for authorization.
type AuthorizationService interface {
	IsUsernameInRole(tenantID TenantID, username, roleName string) (bool, error)
	IsUserInRole(user *User, roleName string) (bool, error)
}
