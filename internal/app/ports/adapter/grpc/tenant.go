package grpc

import (
	"time"

	"github.com/maurofran/iam/internal/app/application"
	"github.com/maurofran/iam/internal/app/application/command"
	context "golang.org/x/net/context"
)

// TenantServer is the GRPC server for tenant server.
type TenantServer struct {
	TenantService *application.TenantService `inject:""`
}

// ProvisionTenant will expose the provision tenant via GRPC.
func (ts *TenantServer) ProvisionTenant(ctx context.Context, req *ProvisionTenantRequest) (*ProvisionTenantResponse, error) {
	cmd := command.ProvisionTenant{
		TenantName:             req.TenantName,
		TenantDescription:      req.TenantDescription,
		AdministratorFirstName: req.AdministratorName.FirstName,
		AdministratorLastName:  req.AdministratorName.LastName,
		EmailAddress:           req.EmailAddress,
		PrimaryTelephone:       req.PrimaryTelephone,
		SecondaryTelephone:     req.SecondaryTelephone,
		AddressStreetName:      req.PostalAddress.StreetName,
		AddressBuildingNumber:  req.PostalAddress.BuildingNumber,
		AddressPostalCode:      req.PostalAddress.PostalCode,
		AddressTown:            req.PostalAddress.Town,
		AddressStateProvince:   req.PostalAddress.StateProvince,
		AddressCountryCode:     req.PostalAddress.CountryCode,
	}
	res, err := ts.TenantService.ProvisionTenant(cmd)
	if err != nil {
		return nil, err
	}
	return &ProvisionTenantResponse{TenantId: res}, nil
}

// ActivateTenant will expose the activate tenant via GRPC
func (ts *TenantServer) ActivateTenant(ctx context.Context, req *ActivateTenantRequest) (*ActivateTenantResponse, error) {
	cmd := command.ActivateTenant{
		TenantID: req.TenantId,
	}
	if err := ts.TenantService.ActivateTenant(cmd); err != nil {
		return nil, err
	}
	return &ActivateTenantResponse{Activated: true}, nil
}

// DeactivateTenant will expose the deactivate tenant via GRPC
func (ts *TenantServer) DeactivateTenant(ctx context.Context, req *DeactivateTenantRequest) (*DeactivateTenantResponse, error) {
	cmd := command.DeactivateTenant{
		TenantID: req.TenantId,
	}
	if err := ts.TenantService.DeactivateTenant(cmd); err != nil {
		return nil, err
	}
	return &DeactivateTenantResponse{Deactivated: true}, nil
}

// OfferInvitation will exposes the offer invitation via GRPC
func (ts *TenantServer) OfferInvitation(ctx context.Context, req *OfferInvitationRequest) (*OfferInvitationResponse, error) {
	cmd := command.OfferInvitation{
		TenantID:    req.TenantId,
		Description: req.Description,
	}
	if req.StartDate != 0 {
		cmd.ValidFrom = time.Unix(req.StartDate, 0)
	}
	if req.EndDate != 0 {
		cmd.ValidTo = time.Unix(req.EndDate, 0)
	}
	res, err := ts.TenantService.OfferInvitation(cmd)
	if err != nil {
		return nil, err
	}
	return &OfferInvitationResponse{InvitationId: res}, nil
}

// WithdrawInvitation will expose the withdraw invitation via GRPC
func (ts *TenantServer) WithdrawInvitation(ctx context.Context, req *WithdrawInvitationRequest) (*WithdrawInvitationResponse, error) {
	cmd := command.WithdrawInvitation{
		TenantID:             req.TenantId,
		InvitationIdentifier: req.InvitationId,
	}
	if err := ts.TenantService.WithdrawInvitation(cmd); err != nil {
		return nil, err
	}
	return &WithdrawInvitationResponse{Withdrawn: true}, nil
}
