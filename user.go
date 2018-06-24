package iam

// Users is the type for a collection of users.
type Users []*User

// User is the aggregate root representing a user.
type User struct {
	TenantID   TenantID   `bson:"tenantId"`
	Username   string     `bson:"username"`
	Password   string     `bson:"password"`
	Enablement Enablement `bson:"enablement"`
	Person     *Person    `bson:"person"`
}

// NewUser will create a new user with supplied initial data.
func NewUser(tenantID, username, password string, person *Person) (*User, Events, error) {
	u := &User{
		TenantID:   TenantID(tenantID),
		Username:   username,
		Password:   "",
		Enablement: IndefiniteEnablement(),
		Person:     &Person{},
	}
	if person != nil {
		u.Person.FullName = person.FullName
		u.Person.ContactInformation = person.ContactInformation
	}
	if _, err := u.ChangePassword("", password); err != nil {
		return nil, nil, err
	}
	events := Events{EventWithPayload(&UserRegistered{
		TenantID:     u.TenantID,
		Username:     u.Username,
		FullName:     u.Person.FullName,
		EmailAddress: u.Person.ContactInformation.EmailAddress,
	})}
	return u, events, nil
}

// ChangePassword will change the new password.
func (u *User) ChangePassword(current, changed string) (Events, error) {
	enc, err := encrypt(current)
	if err != nil {
		return nil, err
	}
	if enc != u.Password {
		return nil, &Error{
			Code:    EINVALID,
			Message: "Current password non confirmed.",
			Op:      "ChangePassword",
		}
	}
	if changed == u.Username {
		return nil, &Error{
			Code:    EINVALID,
			Message: "Changed password must be different from username.",
			Op:      "ChangePassword",
		}
	}
	chenc, err := encrypt(changed)
	if err != nil {
		return nil, err
	}
	if chenc == enc {
		return nil, &Error{
			Code:    EINVALID,
			Message: "Changed password must be different from current one.",
			Op:      "ChangePassword",
		}
	}
	// Check that new password is not equal to current.
	u.Password, err = encrypt(changed)
	if err != nil {
		return nil, err
	}
	return Events{EventWithPayload(&UserPasswordChanged{
		TenantID: u.TenantID,
		Username: u.Username,
	})}, nil
}

// ChangeContactInformation will change the contact information of a user.
func (u *User) ChangeContactInformation(contactInformation ContactInformation) Events {
	u.Person.ContactInformation = contactInformation

	return Events{EventWithPayload(&PersonContactInformationChanged{
		TenantID:           u.TenantID,
		Username:           u.Username,
		ContactInformation: contactInformation,
	})}
}

// ChangeName will change the name of a user.
func (u *User) ChangeName(fullName FullName) Events {
	u.Person.FullName = fullName

	return Events{EventWithPayload(&PersonNameChanged{
		TenantID: u.TenantID,
		Username: u.Username,
		FullName: fullName,
	})}
}

// DefineEnablement will define the user enablement.
func (u *User) DefineEnablement(enablement Enablement) Events {
	u.Enablement = enablement

	return Events{EventWithPayload(&UserEnablementChanged{
		TenantID:   u.TenantID,
		Username:   u.Username,
		Enablement: enablement,
	})}
}

// IsEnabled will check if the user is actually enabled.
func (u *User) IsEnabled() bool {
	return u.Enablement.IsEnabled()
}

// UserRegistered is the event raised when the user is registed.
type UserRegistered struct {
	TenantID     TenantID
	Username     string
	FullName     FullName
	EmailAddress EmailAddress
}

// UserPasswordChanged is the event raised when the password for a user was changed.
type UserPasswordChanged struct {
	TenantID TenantID
	Username string
}

// UserEnablementChanged is the event raised when the user enablement changes.
type UserEnablementChanged struct {
	TenantID   TenantID
	Username   string
	Enablement Enablement
}

// PersonContactInformationChanged is the event raised when person contact information changed.
type PersonContactInformationChanged struct {
	TenantID           TenantID
	Username           string
	ContactInformation ContactInformation
}

// PersonNameChanged is the event raised when person name was changed.
type PersonNameChanged struct {
	TenantID TenantID
	Username string
	FullName FullName
}

// UserRepository is the interace for user repository.
type UserRepository interface {
	Add(*User) error
	Update(*User) error
	Remove(*User) error
	UserWithUsername(TenantID, string) (*User, error)
	AllSimilarlyNamedUsers(TenantID, string, string) (Users, error)
}
