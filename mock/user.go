package mock

import "github.com/maurofran/iam"

// UserRepository is struct for mock user repository
type UserRepository struct {
	AddFn                         func(*iam.User) error
	AddInvoked                    bool
	UpdateFn                      func(*iam.User) error
	UpdateInvoked                 bool
	RemoveFn                      func(*iam.User) error
	RemoveInvoked                 bool
	UserWithUsernameFn            func(iam.TenantID, string) (*iam.User, error)
	UserWithUsernameInvoked       bool
	UserWithCredentialsFn         func(iam.TenantID, string, string) (*iam.User, error)
	UserWithCredentialsInvoked    bool
	AllSimilarlyNamedUsersFn      func(iam.TenantID, string, string) (iam.Users, error)
	AllSimilarlyNamedUsersInvoked bool
}

// Add is the mock of add method.
func (u *UserRepository) Add(user *iam.User) error {
	u.AddInvoked = true
	return u.AddFn(user)
}

// Update is the mock of update method.
func (u *UserRepository) Update(user *iam.User) error {
	u.UpdateInvoked = true
	return u.UpdateFn(user)
}

// Remove is the mock of remove method
func (u *UserRepository) Remove(user *iam.User) error {
	u.RemoveInvoked = true
	return u.RemoveFn(user)
}

// UserWithUsername is the mock of find method.
func (u *UserRepository) UserWithUsername(tenantID iam.TenantID, username string) (*iam.User, error) {
	u.UserWithUsernameInvoked = true
	return u.UserWithUsernameFn(tenantID, username)
}

// UserWithCredentials is the mock of find method.
func (u *UserRepository) UserWithCredentials(tenantID iam.TenantID, username string, password string) (*iam.User, error) {
	u.UserWithCredentialsInvoked = true
	return u.UserWithCredentialsFn(tenantID, username, password)
}

// AllSimilarlyNamedUsers is the mock of find method.
func (u *UserRepository) AllSimilarlyNamedUsers(tenantID iam.TenantID, firstNamePrefix string, lastNamePrefix string) (iam.Users, error) {
	u.AllSimilarlyNamedUsersInvoked = true
	return u.AllSimilarlyNamedUsersFn(tenantID, firstNamePrefix, lastNamePrefix)
}
