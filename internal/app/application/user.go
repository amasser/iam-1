package application

import (
	"context"
	"errors"

	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// ErrUserNotFound is the error returned when a user for supplied data was not found.
var ErrUserNotFound = errors.New("no user found")

// UserService is the object used to manage user services.
type UserService struct {
	TenantRepository      model.TenantRepository       `inject:""`
	UserRepository        model.UserRepository         `inject:""`
	AuthenticationService *model.AuthenticationService `inject:""`
}

// RegisterUser will register a user.
func (us *UserService) RegisterUser(ctx context.Context, cmd command.RegisterUser) error {
	tenant, err := loadTenant(us.TenantRepository, cmd.TenantID)
	if err != nil {
		return err
	}
	enablement, err := model.MakeEnablement(cmd.Enabled, cmd.StartDate, cmd.EndDate)
	if err != nil {
		return err
	}
	fullName, err := model.MakeFullName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}
	contactInformation, err := makeContactInformation(cmd.EmailAddress, cmd.AddressStreetName, cmd.AddressBuildingNumber,
		cmd.AddressPostalCode, cmd.AddressTown, cmd.AddressStateProvince, cmd.AddressCountryCode,
		cmd.PrimaryTelephone, cmd.SecondaryTelephone)
	if err != nil {
		return err
	}
	person, err := model.NewPerson(fullName, contactInformation)
	if err != nil {
		return err
	}
	user, err := tenant.RegisterUser(
		cmd.InvitationIdentifier,
		cmd.Username,
		cmd.Password,
		enablement,
		person,
	)
	if err != nil {
		return err
	}
	return us.UserRepository.Add(user)
}

// AuthenticateUser will authenticate a user.
func (us *UserService) AuthenticateUser(ctx context.Context, cmd command.AuthenticateUser) (model.UserDescriptor, error) {
	theTenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return model.UserDescriptor{}, err
	}
	return us.AuthenticationService.Authenticate(theTenantID, cmd.Username, cmd.Password)
}

// ChangeContactInformation will change the contact information.
func (us *UserService) ChangeContactInformation(ctx context.Context, cmd command.ChangeContactInformation) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	contactInformation, err := makeContactInformation(cmd.EmailAddress, cmd.AddressStreetName, cmd.AddressBuildingNumber,
		cmd.AddressPostalCode, cmd.AddressTown, cmd.AddressStateProvince, cmd.AddressCountryCode,
		cmd.PrimaryTelephone, cmd.SecondaryTelephone)
	if err != nil {
		return err
	}
	if err = user.ChangePersonalContactInformation(contactInformation); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangeEmailAddress will change e-mail address of user.
func (us *UserService) ChangeEmailAddress(ctx context.Context, cmd command.ChangeEmailAddress) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	emailAddress, err := model.MakeEmailAddress(cmd.EmailAddress)
	if err != nil {
		return err
	}
	contactInformation, err := user.Person.ContactInformation.WithEmailAddress(emailAddress)
	if err != nil {
		return err
	}
	if err := user.ChangePersonalContactInformation(contactInformation); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangePostalAddress will change postal address of user.
func (us *UserService) ChangePostalAddress(ctx context.Context, cmd command.ChangePostalAddress) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	postalAddress, err := model.MakePostalAddress(cmd.AddressStreetName, cmd.AddressBuildingNumber,
		cmd.AddressPostalCode, cmd.AddressTown, cmd.AddressStateProvince, cmd.AddressCountryCode)
	if err != nil {
		return err
	}
	contactInformation, err := user.Person.ContactInformation.WithPostalAddress(postalAddress)
	if err != nil {
		return err
	}
	if err := user.ChangePersonalContactInformation(contactInformation); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangePrimaryTelephone will change the primary telephone of user.
func (us *UserService) ChangePrimaryTelephone(ctx context.Context, cmd command.ChangePrimaryTelephone) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	primaryTelephone, err := model.MakeTelephone(cmd.PrimaryTelephone)
	if err != nil {
		return err
	}
	contactInformation, err := user.Person.ContactInformation.WithPrimaryTelephone(primaryTelephone)
	if err != nil {
		return err
	}
	if err := user.ChangePersonalContactInformation(contactInformation); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangeSecondaryTelephone will change the secondary telephone of user.
func (us *UserService) ChangeSecondaryTelephone(ctx context.Context, cmd command.ChangeSecondaryTelephone) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	secondaryTelephone, err := model.MakeTelephone(cmd.SecondaryTelephone)
	if err != nil {
		return err
	}
	contactInformation, err := user.Person.ContactInformation.WithSecondaryTelephone(secondaryTelephone)
	if err != nil {
		return err
	}
	if err := user.ChangePersonalContactInformation(contactInformation); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangeUserPassword will change user password.
func (us *UserService) ChangeUserPassword(ctx context.Context, cmd command.ChangeUserPassword) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	if err := user.ChangePassword(cmd.CurrentPassword, cmd.ChangedPassword); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// ChangeUserPersonalName will change the user personal name.
func (us *UserService) ChangeUserPersonalName(ctx context.Context, cmd command.ChangeUserPersonalName) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	fullName, err := model.MakeFullName(cmd.FirstName, cmd.LastName)
	if err != nil {
		return err
	}
	if err := user.ChangePersonalName(fullName); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

// DefineUserEnablement will define the user enablement.
func (us *UserService) DefineUserEnablement(ctx context.Context, cmd command.DefineUserEnablement) error {
	user, err := loadUser(us.UserRepository, cmd.TenantID, cmd.Username)
	if err != nil {
		return err
	}
	enablement, err := model.MakeEnablement(cmd.Enabled, cmd.StartDate, cmd.EndDate)
	if err != nil {
		return err
	}
	if err := user.DefineEnablement(enablement); err != nil {
		return err
	}
	return us.UserRepository.Update(user)
}

func loadUser(repo model.UserRepository, tenantID, username string) (*model.User, error) {
	theTenantID, err := model.MakeTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	user, err := repo.UserWithUsername(theTenantID, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
