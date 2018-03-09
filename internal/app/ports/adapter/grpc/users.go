package grpc

import (
	"time"

	"github.com/maurofran/iam/internal/app/application"
	"github.com/maurofran/iam/internal/app/application/command"
	context "golang.org/x/net/context"
)

// UserServer is GRPC service for users.
type UserServer struct {
	UserService *application.UserService `inject:""`
}

// RegisterUser will expose the register user through GRPC
func (us *UserServer) RegisterUser(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	cmd := command.RegisterUser{
		TenantID:              req.TenantId,
		InvitationIdentifier:  req.InvitationId,
		Username:              req.Username,
		Password:              req.Password,
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		Enabled:               req.Enabled,
		StartDate:             time.Unix(req.StartDate, 0),
		EndDate:               time.Unix(req.EndDate, 0),
		EmailAddress:          req.EmailAddress,
		PrimaryTelephone:      req.PrimaryTelephone,
		SecondaryTelephone:    req.SecondaryTelephone,
		AddressStreetName:     req.AddressStreetName,
		AddressBuildingNumber: req.AddressBuildingNumber,
		AddressPostalCode:     req.AddressPostalCode,
		AddressTown:           req.AddressTown,
		AddressStateProvince:  req.AddressStateProvince,
		AddressCountryCode:    req.AddressCountryCode,
	}
	if err := us.UserService.RegisterUser(ctx, cmd); err != nil {
		return nil, err
	}
	return &RegisterUserResponse{}, nil
}

// AuthenticateUser will expose the authenticate user through GRPC
func (us *UserServer) AuthenticateUser(ctx context.Context, req *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	cmd := command.AuthenticateUser{
		TenantID: req.TenantId,
		Username: req.Username,
		Password: req.Password,
	}
	user, err := us.UserService.AuthenticateUser(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return &AuthenticateUserResponse{
		TenantId:     user.TenantID.String(),
		Username:     user.Username,
		FirstName:    user.Name.FirstName,
		LastName:     user.Name.LastName,
		EmailAddress: user.EmailAddress.String(),
	}, nil
}

// ChangeContactInformation exposes the change contact information function through GRPC
func (us *UserServer) ChangeContactInformation(ctx context.Context, req *ChangeContactInformationRequest) (*ChangeContactInformationResponse, error) {
	cmd := command.ChangeContactInformation{
		TenantID:              req.TenantId,
		Username:              req.Username,
		EmailAddress:          req.EmailAddress,
		PrimaryTelephone:      req.PrimaryTelephone,
		SecondaryTelephone:    req.SecondaryTelephone,
		AddressStreetName:     req.AddressStreetName,
		AddressBuildingNumber: req.AddressBuildingNumber,
		AddressPostalCode:     req.AddressPostalCode,
		AddressTown:           req.AddressTown,
		AddressStateProvince:  req.AddressStateProvince,
		AddressCountryCode:    req.AddressCountryCode,
	}
	if err := us.UserService.ChangeContactInformation(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangeContactInformationResponse{}, nil
}

// ChangeEmailAddress exposes the change email address function through GRPC
func (us *UserServer) ChangeEmailAddress(ctx context.Context, req *ChangeEmailAddressRequest) (*ChangeEmailAddressResponse, error) {
	cmd := command.ChangeEmailAddress{
		TenantID:     req.TenantId,
		Username:     req.Username,
		EmailAddress: req.EmailAddress,
	}
	if err := us.UserService.ChangeEmailAddress(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangeEmailAddressResponse{}, nil
}

// ChangePostalAddress exposes the change postal address function through GRPC
func (us *UserServer) ChangePostalAddress(ctx context.Context, req *ChangePostalAddressRequest) (*ChangePostalAddressResponse, error) {
	cmd := command.ChangePostalAddress{
		TenantID:              req.TenantId,
		Username:              req.Username,
		AddressStreetName:     req.AddressStreetName,
		AddressBuildingNumber: req.AddressBuildingNumber,
		AddressPostalCode:     req.AddressPostalCode,
		AddressTown:           req.AddressTown,
		AddressStateProvince:  req.AddressStateProvince,
		AddressCountryCode:    req.AddressCountryCode,
	}
	if err := us.UserService.ChangePostalAddress(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangePostalAddressResponse{}, nil
}

// ChangePrimaryTelephone exposes the change primary telephone function through GRPC
func (us *UserServer) ChangePrimaryTelephone(ctx context.Context, req *ChangePrimaryTelephoneRequest) (*ChangePrimaryTelephoneResponse, error) {
	cmd := command.ChangePrimaryTelephone{
		TenantID:         req.TenantId,
		Username:         req.Username,
		PrimaryTelephone: req.PrimaryTelephone,
	}
	if err := us.UserService.ChangePrimaryTelephone(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangePrimaryTelephoneResponse{}, nil
}

// ChangeSecondaryTelephone exposes the change secondary telephone function through GRPC
func (us *UserServer) ChangeSecondaryTelephone(ctx context.Context, req *ChangeSecondaryTelephoneRequest) (*ChangeSecondaryTelephoneResponse, error) {
	cmd := command.ChangeSecondaryTelephone{
		TenantID:           req.TenantId,
		Username:           req.Username,
		SecondaryTelephone: req.SecondaryTelephone,
	}
	if err := us.UserService.ChangeSecondaryTelephone(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangeSecondaryTelephoneResponse{}, nil
}

// ChangeUserPassword exposes the change password function through GRPC
func (us *UserServer) ChangeUserPassword(ctx context.Context, req *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	cmd := command.ChangeUserPassword{
		TenantID:        req.TenantId,
		Username:        req.Username,
		CurrentPassword: req.CurrentPassword,
		ChangedPassword: req.ChangedPassword,
	}
	if err := us.UserService.ChangeUserPassword(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangeUserPasswordResponse{}, nil
}

// ChangeUserPersonalName exposes the change password function through GRPC
func (us *UserServer) ChangeUserPersonalName(ctx context.Context, req *ChangeUserPersonalNameRequest) (*ChangeUserPersonalNameResponse, error) {
	cmd := command.ChangeUserPersonalName{
		TenantID:  req.TenantId,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	if err := us.UserService.ChangeUserPersonalName(ctx, cmd); err != nil {
		return nil, err
	}
	return &ChangeUserPersonalNameResponse{}, nil
}

// DefineUserEnablement exposes the define user enablement function through GRPC
func (us *UserServer) DefineUserEnablement(ctx context.Context, req *DefineUserEnablementRequest) (*DefineUserEnablementResponse, error) {
	cmd := command.DefineUserEnablement{
		TenantID:  req.TenantId,
		Username:  req.Username,
		Enabled:   req.Enabled,
		StartDate: time.Unix(req.StartDate, 0),
		EndDate:   time.Unix(req.EndDate, 0),
	}
	if err := us.UserService.DefineUserEnablement(ctx, cmd); err != nil {
		return nil, err
	}
	return &DefineUserEnablementResponse{}, nil
}
