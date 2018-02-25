package model

import (
	"time"

	"github.com/maurofran/iam/internal/pkg/aggregate"
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
	aggregate.Root
	TenantID   TenantID
	Username   string
	Password   string
	Enablement Enablement
	Person     *Person
}

// NewUser will create a new user with supplied data.
//
// The function will return an error if any of supplied data is not valid.
func newUser(tenantID TenantID, username, password string, enablement Enablement, person *Person) (*User, error) {
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
	user := &User{TenantID: tenantID, Username: username, Enablement: enablement, Person: person}
	if err := user.protectPassword("", password); err != nil {
		return nil, err
	}
	user.RegisterEvent(&UserRegistered{
		EventVersion: 1,
		OccurredOn:   time.Now(),
		TenantID:     user.TenantID,
		Username:     user.Username,
		EmailAddress: user.Person.EmailAddress(),
		FullName:     user.Person.Name,
	})
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
	if err := u.protectPassword(currentPassword, newPassword); err != nil {
		return err
	}
	u.RegisterEvent(UserPasswordChanged{
		EventVersion: 1,
		OccurredOn:   time.Now(),
		TenantID:     u.TenantID,
		Username:     u.Username,
	})
	return nil
}

// ChangePersonalContactInformation will change the personal contact information.
func (u *User) ChangePersonalContactInformation(contactInformation ContactInformation) error {
	changed, err := u.Person.changeContactInformation(contactInformation)
	if err != nil && changed {
		u.RegisterEvent(PersonContactInformationChanged{
			EventVersion:       1,
			OccurredOn:         time.Now(),
			TenantID:           u.TenantID,
			Username:           u.Username,
			ContactInformation: u.Person.ContactInformation,
		})
	}
	return err
}

// ChangePersonalName will change the personal name.
func (u *User) ChangePersonalName(name FullName) error {
	changed, err := u.Person.changeName(name)
	if err != nil && changed {
		u.RegisterEvent(PersonNameChanged{
			EventVersion: 1,
			OccurredOn:   time.Now(),
			TenantID:     u.TenantID,
			Username:     u.Username,
			Name:         u.Person.Name,
		})
	}
	return err
}

// DefineEnablement the enablement of this user.
func (u *User) DefineEnablement(enablement Enablement) error {
	if u.Enablement != enablement {
		u.Enablement = enablement
		u.RegisterEvent(UserEnablementChanged{
			EventVersion: 1,
			OccurredOn:   time.Now(),
			TenantID:     u.TenantID,
			Username:     u.Username,
			Enablement:   u.Enablement,
		})
	}
	return nil
}

// Enabled will retrieve the user actual enablement status.
func (u *User) Enabled() bool {
	return u.Enablement.ActuallyEnabled()
}

func (u *User) toGroupMember() GroupMember {
	return GroupMember{MemberTypeUser, u.Username}
}

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

// UserRegistered is the event raised when a new user is registered.
type UserRegistered struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
	Username     string
	EmailAddress EmailAddress
	FullName     FullName
}

// UserPasswordChanged is the event raised when the password of a user changed.
type UserPasswordChanged struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
	Username     string
}

// UserEnablementChanged is the event raised when the user enablement status changed.
type UserEnablementChanged struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
	Username     string
	Enablement   Enablement
}
