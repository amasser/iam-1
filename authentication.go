package iam

// AuthenticationService is the service for an authentication.
type AuthenticationService interface {
	Authenticate(tenantID TenantID, username, password string) (*User, error)
}
