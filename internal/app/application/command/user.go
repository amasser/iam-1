package command

import "time"

// RegisterUser will register a new user.
type RegisterUser struct {
	TenantID              string    `json:"tenantId"`
	InvitationIdentifier  string    `json:"invitationId"`
	Username              string    `json:"username"`
	Password              string    `json:"password"`
	FirstName             string    `json:"firstName"`
	LastName              string    `json:"lastName"`
	Enabled               bool      `json:"enabled"`
	StartDate             time.Time `json:"startDate,omitempty"`
	EndDate               time.Time `json:"endDate,omitempty"`
	EmailAddress          string    `json:"emailAddress"`
	PrimaryTelephone      string    `json:"primaryTelephone"`
	SecondaryTelephone    string    `json:"secondaryTelephone"`
	AddressStreetAddress  string    `json:"addressStreetAddress"`
	AddressBuildingNumber string    `json:"addressStreetNumber"`
	AddressPostalCode     string    `json:"addressPostalCode"`
	AddressTown           string    `json:"addressTown"`
	AddressStateProvince  string    `json:"addressStateProvince"`
	AddressCountryCode    string    `json:"addressCountryCode"`
}

// AuthenticateUser will authenticate a user.
type AuthenticateUser struct {
	TenantID string `json:"tenantId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ChangeContactInformation will change the contact information.
type ChangeContactInformation struct {
	TenantID              string `json:"tenantId"`
	Username              string `json:"username"`
	EmailAddress          string `json:"emailAddress"`
	PrimaryTelephone      string `json:"primaryTelephone"`
	SecondaryTelephone    string `json:"secondaryTelephone"`
	AddressStreetAddress  string `json:"addressStreetAddress"`
	AddressBuildingNumber string `json:"addressStreetNumber"`
	AddressPostalCode     string `json:"addressPostalCode"`
	AddressTown           string `json:"addressTown"`
	AddressStateProvince  string `json:"addressStateProvince"`
	AddressCountryCode    string `json:"addressCountryCode"`
}

// ChangeEmailAddress will change the email address.
type ChangeEmailAddress struct {
	TenantID     string `json:"tenantId"`
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
}

// CahngePrimaryTelephone will change the primary telephone.
type CahngePrimaryTelephone struct {
	TenantID         string `json:"tenantId"`
	Username         string `json:"username"`
	PrimaryTelephone string `json:"primaryTelephone"`
}

// CahngeSecondaryTelephone will change the secondary telephone.
type CahngeSecondaryTelephone struct {
	TenantID           string `json:"tenantId"`
	Username           string `json:"username"`
	SecondaryTelephone string `json:"secondaryTelephone"`
}

// ChangePostalAddress will change the postal address.
type ChangePostalAddress struct {
	TenantID              string `json:"tenantId"`
	Username              string `json:"username"`
	AddressStreetAddress  string `json:"addressStreetAddress"`
	AddressBuildingNumber string `json:"addressStreetNumber"`
	AddressPostalCode     string `json:"addressPostalCode"`
	AddressTown           string `json:"addressTown"`
	AddressStateProvince  string `json:"addressStateProvince"`
	AddressCountryCode    string `json:"addressCountryCode"`
}

// ChangeUserPassword will change a user password.
type ChangeUserPassword struct {
	TenantID        string `json:"tenantId"`
	Username        string `json:"username"`
	CurentPassword  string `json:"currentPassword"`
	ChangedPassword string `json:"changedPassword"`
}

// ChangeUserPersonalName will change a user personal name.
type ChangeUserPersonalName struct {
	TenantID  string `json:"tenantId"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// DefineUserEnablement will define a user enablement.
type DefineUserEnablement struct {
	TenantID  string    `json:"tenantId"`
	Username  string    `json:"username"`
	Enabled   bool      `json:"enabled"`
	StartDate time.Time `json:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty"`
}
