package command

import "time"

// ProvisionTenant is the command issued to provision a new tenant.
type ProvisionTenant struct {
	TenantName             string `json:"tenantName"`
	TenantDescription      string `json:"tenantDescription"`
	AdministratorFirstName string `json:"administratorFirstName"`
	AdministratorLastName  string `json:"administratorLastName"`
	EmailAddress           string `json:"emailAddress"`
	PrimaryTelephone       string `json:"primaryTelephone"`
	SecondaryTelephone     string `json:"secondaryTelephone"`
	AddressStreetName      string `json:"addressStreetName"`
	AddressBuildingNumber  string `json:"addressStreetNumber"`
	AddressPostalCode      string `json:"addressPostalCode"`
	AddressTown            string `json:"addressTown"`
	AddressStateProvince   string `json:"addressStateProvince"`
	AddressCountryCode     string `json:"addressCountryCode"`
}

// ActivateTenant is the command issued to activate a tenant.
type ActivateTenant struct {
	TenantID string `json:"tenantId"`
}

// DeactivateTenant is the command issued to deactivate a tenant.
type DeactivateTenant struct {
	TenantID string `json:"tenantId"`
}

// OfferInvitation will offer an invitation.
type OfferInvitation struct {
	TenantID    string    `json:"tenantId"`
	Description string    `json:"description"`
	ValidFrom   time.Time `json:"validFrom,omitempty"`
	ValidTo     time.Time `json:"validTo,omitempty"`
}

// WithdrawInvitation will widthraw an invitation.
type WithdrawInvitation struct {
	TenantID             string `json:"tenantId"`
	InvitationIdentifier string `json:"invitationId"`
}
