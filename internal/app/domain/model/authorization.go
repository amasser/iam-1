package model

// AuthorizationService is the domain service used to manage authorization.
type AuthorizationService struct {
	UserRepo  UserRepository
	GroupRepo GroupRepository
	RoleRepo  RoleRepository
}

// IsUserOfTenantInRole verify if the user of supplied tenant is in role.
func (as *AuthorizationService) IsUserOfTenantInRole(tenantID TenantID, username, role string) (bool, error) {
	user, err := as.UserRepo.UserWithUsername(tenantID, username)
	if err != nil {
		return false, err
	}
	if user != nil {
		return as.IsUserInRole(user, role)
	}
	return false, nil
}

// IsUserInRole verify if the user is in role.
func (as *AuthorizationService) IsUserInRole(user *User, role string) (bool, error) {
	if user.Enabled() {
		role, err := as.RoleRepo.RoleNamed(user.TenantID, role)
		if err != nil {
			return false, err
		}
		if role != nil {
			gms := &GroupMemberService{UserRepo: as.UserRepo, GroupRepo: as.GroupRepo}
			return role.IsInRole(user, gms)
		}
	}
	return false, nil
}
