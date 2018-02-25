package model

import "github.com/maurofran/iam/internal/pkg/password"

// AuthenticationService is the domain service used to perform authentication.
type AuthenticationService struct {
	TenantRepo TenantRepository
	UserRepo   UserRepository
}

// Authenticate will perform authentication of user with supplied credentials.
func (as *AuthenticationService) Authenticate(tenantID TenantID, username, plainPassword string) (UserDescriptor, error) {
	desc := UserDescriptor{}
	tenant, err := as.TenantRepo.TenantWithID(tenantID)
	if err != nil {
		return desc, err
	}
	if tenant != nil && tenant.Active {
		encrypted, err := password.Encrypt(plainPassword)
		if err != nil {
			return desc, err
		}
		user, err := as.UserRepo.UserWithCredentials(tenantID, username, encrypted)
		if err != nil {
			return desc, err
		}
		if user != nil && user.Enabled() {
			return user.toDescriptor(), nil
		}
	}
	return desc, nil
}
