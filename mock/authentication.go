package mock

import "github.com/maurofran/iam"

// AuthenticationService is the mock implementation of authentication service interface.
type AuthenticationService struct {
	AuthenticateFn      func(iam.TenantID, string, string) (*iam.User, error)
	AuthenticateInvoked bool
}

// Authenticate function mock the authenticate service call.
func (a *AuthenticationService) Authenticate(tenantID iam.TenantID, username, password string) (*iam.User, error) {
	a.AuthenticateInvoked = true
	return a.AuthenticateFn(tenantID, username, password)
}
