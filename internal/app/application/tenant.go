package application

import (
	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// TenantService is the application service for tenants.
type TenantService struct {
	TenantRepository    model.TenantRepository           `inject:""`
	ProvisioningService *model.TenantProvisioningService `inject:""`
}

// ProvisionTenant will provision a new tenant with data for supplied command.
func (ts *TenantService) ProvisionTenant(cmd command.ProvisionTenant) (string, error) {
	var err error
	fullName, err := model.MakeFullName(cmd.AdministratorFirstName, cmd.AdministratorLastName)
	if err != nil {
		return "", err
	}
	ci, err := makeContactInformation(cmd.EmailAddress, cmd.AddressStreetName, cmd.AddressBuildingNumber,
		cmd.AddressPostalCode, cmd.AddressTown, cmd.AddressStateProvince, cmd.AddressCountryCode,
		cmd.PrimaryTelephone, cmd.SecondaryTelephone)
	if err != nil {
		return "", err
	}

	tenant, err := ts.ProvisioningService.ProvisionTenant(
		cmd.TenantName,
		cmd.TenantDescription,
		fullName,
		ci.EmailAddress,
		ci.PostalAddress,
		ci.PrimaryTelephone,
		ci.SecondaryTelephone,
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
	tenant, err := ts.TenantRepository.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	tenant.Activate()
	return ts.TenantRepository.Update(tenant)
}

// DeactivateTenant will deactivate a tenant.
func (ts *TenantService) DeactivateTenant(cmd command.DeactivateTenant) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := ts.TenantRepository.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	tenant.Deactivate()
	return ts.TenantRepository.Update(tenant)
}

// OfferInvitation will offer an invitation for tenant.
func (ts *TenantService) OfferInvitation(cmd command.OfferInvitation) (string, error) {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return "", err
	}
	tenant, err := ts.TenantRepository.TenantWithID(tenantID)
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
	err = ts.TenantRepository.Update(tenant)
	return invitation.InvitationID, err
}

// WithdrawInvitation will withdraw an invitation for tenant.
func (ts *TenantService) WithdrawInvitation(cmd command.WithdrawInvitation) error {
	tenantID, err := model.MakeTenantID(cmd.TenantID)
	if err != nil {
		return err
	}
	tenant, err := ts.TenantRepository.TenantWithID(tenantID)
	if err != nil {
		return err
	}
	err = tenant.WithdrawInvitation(cmd.InvitationIdentifier)
	if err != nil {
		return err
	}
	return ts.TenantRepository.Update(tenant)
}
