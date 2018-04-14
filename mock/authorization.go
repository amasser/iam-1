package mock

import "github.com/maurofran/iam"

// AuthorizationService is the mock authorization service implementaation.
type AuthorizationService struct {
	IsUsernameInRoleFn      func(iam.TenantID, string, string) (bool, error)
	IsUsernameInRoleInvoked bool
	IsUserInRoleFn          func(*iam.User, string) (bool, error)
	IsUserInRoleInvoked     bool
}

// IsUsernameInRole is the mock implementation of service method.
func (a *AuthorizationService) IsUsernameInRole(tenantID iam.TenantID, username, roleName string) (bool, error) {
	a.IsUsernameInRoleInvoked = true
	return a.IsUsernameInRoleFn(tenantID, username, roleName)
}

// IsUserInRole is the mock implementation of service method.
func (a *AuthorizationService) IsUserInRole(user *iam.User, roleName string) (bool, error) {
	a.IsUserInRoleInvoked = true
	return a.IsUserInRoleFn(user, roleName)
}
