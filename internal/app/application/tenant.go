package application

import (
	"context"
	"errors"

	"github.com/maurofran/iam/internal/app/application/command"
	"github.com/maurofran/iam/internal/app/domain/model"
)

// ErrTenantNotFound is returned when no tenant is found.
var ErrTenantNotFound = errors.New("tenant not found")

// TenantService is the application service for tenants.
type TenantService struct {
	TenantRepository    model.TenantRepository           `inject:""`
	ProvisioningService *model.TenantProvisioningService `inject:""`
}

// ProvisionTenant will provision a new tenant with data for supplied command.
func (ts *TenantService) ProvisionTenant(ctx context.Context, cmd command.ProvisionTenant) (string, error) {
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
func (ts *TenantService) ActivateTenant(ctx context.Context, cmd command.ActivateTenant) error {
	tenant, err := loadTenant(ts.TenantRepository, cmd.TenantID)
	if err != nil {
		return err
	}
	tenant.Activate()
	return ts.TenantRepository.Update(tenant)
}

// DeactivateTenant will deactivate a tenant.
func (ts *TenantService) DeactivateTenant(ctx context.Context, cmd command.DeactivateTenant) error {
	tenant, err := loadTenant(ts.TenantRepository, cmd.TenantID)
	if err != nil {
		return err
	}
	tenant.Deactivate()
	return ts.TenantRepository.Update(tenant)
}

// OfferInvitation will offer an invitation for tenant.
func (ts *TenantService) OfferInvitation(ctx context.Context, cmd command.OfferInvitation) (string, error) {
	tenant, err := loadTenant(ts.TenantRepository, cmd.TenantID)
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
func (ts *TenantService) WithdrawInvitation(ctx context.Context, cmd command.WithdrawInvitation) error {
	tenant, err := loadTenant(ts.TenantRepository, cmd.TenantID)
	if err != nil {
		return err
	}
	err = tenant.WithdrawInvitation(cmd.InvitationIdentifier)
	if err != nil {
		return err
	}
	return ts.TenantRepository.Update(tenant)
}

func loadTenant(repo model.TenantRepository, tenantID string) (*model.Tenant, error) {
	aTenantID, err := model.MakeTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	tenant, err := repo.TenantWithID(aTenantID)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, ErrTenantNotFound
	}
	return tenant, nil
}
