package model

import (
	"github.com/maurofran/iam/internal/pkg/password"
	"github.com/maurofran/kit/assert"
)

// UserRepository is the interface for managing.
type UserRepository interface {
	// Add a user to virtual collection.
	Add(*User) error
	// Updates a user already in virtual collection.
	Update(*User) error
	// Removes a user from virtual collection.
	Remove(*User) error
	// AllUsers will retrieve all the users of virtual collection for tenant.
	AllUsers(TenantID) ([]*User, error)
	// UserWithUsername will retrieve a user for a username.
	UserWithUsername(TenantID, string) (*User, error)
}

// User is the aggregate root used to provide a user.
type User struct {
	TenantID   TenantID
	Username   string
	Password   string
	Enablement Enablement
	Person     *Person
}

// NewUser will create a new user with supplied data.
//
// The function will return an error if any of supplied data is not valid.
func (u *User) NewUser(tenantID TenantID, username, password string, enablement Enablement, person *Person) (*User, error) {
	if err := assert.NotZero(tenantID, "tenantID"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(username, "username"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(password, "password"); err != nil {
		return nil, err
	}
	if err := assert.NotNil(person, "person"); err != nil {
		return nil, err
	}
	user := &User{tenantID, username, "", enablement, person}
	if err := user.protectPassword("", password); err != nil {
		return nil, err
	}
	// TODO Raise event
	return user, nil
}

// ChangePassword will change the password, after checking for current one.
func (u *User) ChangePassword(currentPassword, newPassword string) error {
	if err := assert.NotEmpty(currentPassword, "currentPassword"); err != nil {
		return err
	}
	encrypted, err := password.Encrypt(currentPassword)
	if err != nil {
		return err
	}
	if err := assert.Condition(u.Password == encrypted, "currentPassword not confirmed"); err != nil {
		return err
	}
	return u.protectPassword(currentPassword, newPassword)
}

// ChangePersonalContactInformation will change the personal contact information.
func (u *User) ChangePersonalContactInformation(contactInformation ContactInformation) error {
	_, err := u.Person.changeContactInformation(contactInformation)
	// TODO Raise event
	return err
}

// ChangePersonalName will change the personal name.
func (u *User) ChangePersonalName(name FullName) error {
	_, err := u.Person.changeName(name)
	// TODO Raise event
	return err
}

// DefineEnablement the enablement of this user.
func (u *User) DefineEnablement(enablement Enablement) error {
	u.Enablement = enablement
	// TODO Raise event
	return nil
}

// Enabled will retrieve the user actual enablement status.
func (u *User) Enabled() bool {
	return u.Enablement.ActuallyEnabled()
}

// TODO toGroupMember

func (u *User) protectPassword(currentPassword, newPassword string) error {
	if err := assert.NotEquals(currentPassword, newPassword, "currentPassword"); err != nil {
		return err
	}
	if err := assert.Condition(!password.IsWeak(newPassword), "newPassword is weak"); err != nil {
		return err
	}
	if err := assert.NotEquals(newPassword, u.Username, "newPassword"); err != nil {
		return err
	}
	encrypted, err := password.Encrypt(newPassword)
	if err != nil {
		return err
	}
	u.Password = encrypted
	return nil
}
