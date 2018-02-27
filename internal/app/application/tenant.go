package application

import (
	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// TenantService is the application service for tenants.
type TenantService struct {
	TenantRepo          model.TenantRepository
	ProvisioningService *model.TenantProvisioningService
}

// ProvisionTenant will provision a new tenant with data for supplied command.
func (ts *TenantService) ProvisionTenant(cmd command.ProvisionTenant) (string, error) {
	var err error
	fullName, err := model.MakeFullName(cmd.AdministratorFirstName, cmd.AdministratorLastName)
	if err != nil {
		return "", err
	}
	emailAddress, err := model.MakeEmailAddress(cmd.EmailAddress)
	if err != nil {
		return "", err
	}
	postalAddress, err := model.MakePostalAddress(cmd.AddressStreetAddress, cmd.AddressBuildingNumber,
		cmd.AddressPostalCode, cmd.AddressTown, cmd.AddressStateProvince, cmd.AddressCountryCode)
	if err != nil {
		return "", err
	}
	primaryTelephone, err := model.MakeTelephone(cmd.PrimaryTelephone)
	if err != nil {
		return "", err
	}
	secondaryTelephone := model.Telephone("")
	if cmd.SecondaryTelephone != "" {
		secondaryTelephone, err = model.MakeTelephone(cmd.SecondaryTelephone)
		if err != nil {
			return "", err
		}
	}
	tenant, err := ts.ProvisioningService.ProvisionTenant(
		cmd.TenantName,
		cmd.TenantDescription,
		fullName,
		emailAddress,
		postalAddress,
		primaryTelephone,
		secondaryTelephone,
	)
	if err != nil {
		return "", err
	}
	return string(tenant.ID), nil
}

// ActivateTenant will activate a tenant.
func (ts *TenantService) ActivateTenant(cmd command.ActivateTenant) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := ts.TenantRepo.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	tenant.Activate()
	return ts.TenantRepo.Update(tenant)
}

// DeactivateTenant will deactivate a tenant.
func (ts *TenantService) DeactivateTenant(cmd command.DeactivateTenant) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := ts.TenantRepo.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	tenant.Deactivate()
	return ts.TenantRepo.Update(tenant)
}

// OfferInvitation will offer an invitation for tenant.
func (ts *TenantService) OfferInvitation(cmd command.OfferInvitation) (string, error) {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return "", err
	}
	tenant, err := ts.TenantRepo.TenantWithID(tenantID)
	if err != nil {
		return "", err
	}
	invitation, err := tenant.OfferInvitation(cmd.Description)
	if err != nil {
		return "", err
	}
	if !cmd.ValidFrom.IsZero() && !cmd.ValidTo.IsZero() {
		if err := invitation.RedefineAs(cmd.ValidFrom, cmd.ValidTo); err != nil {
			tenant.WithdrawInvitation(invitation.InvitationID)
			return "", err
		}
	}
	err = ts.TenantRepo.Update(tenant)
	return invitation.InvitationID, err
}

// WithdrawInvitation will withdraw an invitation for tenant.
func (ts *TenantService) WithdrawInvitation(cmd command.WithdrawInvitation) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := ts.TenantRepo.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	err = tenant.WithdrawInvitation(cmd.InvitationIdentifier)
	if err != nil {
		return err
	}
	return ts.TenantRepo.Update(tenant)
}
