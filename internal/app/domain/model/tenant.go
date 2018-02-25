package model

import (
	"time"

	"github.com/maurofran/iam/internal/pkg/event"
	"github.com/maurofran/iam/internal/pkg/password"

	"github.com/google/uuid"
	"github.com/maurofran/iam/internal/pkg/aggregate"
	"github.com/maurofran/kit/assert"
	"github.com/pkg/errors"
)

// TenantID is the value object used to provide unique tenant identifier.
type TenantID string

// MakeTenantID will build a new tenant identifier, returning an error if it's empty.
func MakeTenantID(value string) (TenantID, error) {
	if err := assert.NotEmpty(value, "value"); err != nil {
		return "", err
	}
	return TenantID(value), nil
}

// randomTenantID will return a new random tenant identifier randomly genenrated.
func randomTenantID() (TenantID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "unable to generated random tenantId")
	}
	return TenantID(id.String()), nil
}

// IsZero will check if tenant identifier is zero value.
func (t TenantID) IsZero() bool {
	return t == ""
}

// TenantRepository is the interface of virtual collection of tenants.
type TenantRepository interface {
	Add(*Tenant) error
	Update(*Tenant) error
	Remove(*Tenant) error
	TenantWithID(TenantID) (*Tenant, error)
	TenantWithName(string) (*Tenant, error)
}

// Tenant is the aggregate root representing a tenant.
type Tenant struct {
	aggregate.Root `bson:"-"`
	ID             TenantID      `bson:"tenantId"`
	Name           string        `bson:"name"`
	Description    string        `bson:"description"`
	Active         bool          `bson:"active"`
	Invitations    []*Invitation `bson:"invitations"`
}

// NewTenant will create a new tenant with supplied data.
func newTenant(name, description string, active bool) (*Tenant, error) {
	id, err := randomTenantID()
	if err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(name, "name"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(description, "description"); err != nil {
		return nil, err
	}
	tenant := &Tenant{ID: id, Name: name, Description: description, Active: active}
	tenant.RegisterEvent(TenantProvisioned{
		EventVersion: 1,
		OccurredOn:   time.Now(),
		TenantID:     tenant.ID,
		Name:         tenant.Name,
	})
	return tenant, nil
}

func assertTenantActive(t *Tenant) error {
	return assertTenantActive(t)
}

// Activate will make this tenant active.
func (t *Tenant) Activate() {
	if !t.Active {
		t.Active = true
		t.RegisterEvent(TenantActivated{
			EventVersion: 1,
			OccurredOn:   time.Now(),
			TenantID:     t.ID,
		})
	}
}

// Deactivate will make this tenant inactive.
func (t *Tenant) Deactivate() {
	if t.Active {
		t.Active = false
		t.RegisterEvent(TenantDeactivated{
			EventVersion: 1,
			OccurredOn:   time.Now(),
			TenantID:     t.ID,
		})
	}
}

// OfferInvitation will create a new invitation into this tenant.
func (t *Tenant) OfferInvitation(description string) (*Invitation, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	exists, err := t.IsInvitationAvailableThrough(description)
	if err != nil {
		return nil, err
	}
	if err := assert.Condition(!exists, "invitation already exists"); err != nil {
		return nil, err
	}
	invitation, err := newInvitation(description)
	if err != nil {
		return nil, err
	}
	t.Invitations = append(t.Invitations, invitation)
	return invitation, nil
}

// InvitationFor will retrieve the invitation for supplied identifier or nil.
func (t *Tenant) InvitationFor(identifier string) (*Invitation, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return t.invitationFor(identifier), nil
}

// WithdrawInvitation will remove the invitation for supplied identifier.
func (t *Tenant) WithdrawInvitation(identifier string) error {
	if err := assertTenantActive(t); err != nil {
		return err
	}
	for i, invitation := range t.Invitations {
		if invitation.identifiedBy(identifier) {
			var res []*Invitation
			copy(res, t.Invitations)

			res[i] = res[len(res)-1]
			res[len(res)-1] = nil
			res = res[:len(res)-1]

			t.Invitations = res
		}
	}
	return nil
}

// IsInvitationAvailableThrough will check if an invitation is available through the supplied identifier
func (t *Tenant) IsInvitationAvailableThrough(identifier string) (bool, error) {
	if err := assertTenantActive(t); err != nil {
		return false, err
	}
	return t.invitationFor(identifier) != nil, nil
}

// AllAvailableInvitations will retrieve a slice of all available invitation descriptors.
func (t *Tenant) AllAvailableInvitations() ([]InvitationDescriptor, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return t.allInvitationsFor(true), nil
}

// AllUnavailableInvitations will retrieve a slice of all unavailable invitation descriptors.
func (t *Tenant) AllUnavailableInvitations() ([]InvitationDescriptor, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return t.allInvitationsFor(false), nil
}

// RegisterUser will register a new user for this tenant.
func (t *Tenant) RegisterUser(identifier, username, password string, enablement Enablement, person *Person) (*User, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	available, err := t.IsInvitationAvailableThrough(identifier)
	if err != nil {
		return nil, err
	}
	if !available {
		user, err := newUser(t.ID, username, password, enablement, person)
		if err != nil {
			return nil, err
		}
		if err = t.WithdrawInvitation(identifier); err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

// ProvisionGroup will provision a new group.
func (t *Tenant) ProvisionGroup(name, description string) (*Group, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return newGroup(t.ID, name, description)
}

// ProvisionRole will provision a new role for this tenant.
func (t *Tenant) ProvisionRole(name, description string, supportsNesting bool) (*Role, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return newRole(t.ID, name, description, supportsNesting)
}

func (t *Tenant) invitationFor(identifier string) *Invitation {
	for _, i := range t.Invitations {
		if i.identifiedBy(identifier) {
			return i
		}
	}
	return nil
}

func (t *Tenant) allInvitationsFor(status bool) []InvitationDescriptor {
	res := make([]InvitationDescriptor, 0, len(t.Invitations))
	for _, i := range t.Invitations {
		if i.Available() == status {
			res = append(res, i.toDescriptor(t.ID))
		}
	}
	return res
}

// TenantProvisioned is the event raised when a new tenant is provisioned.
type TenantProvisioned struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
	Name         string
}

// TenantActivated is the event raised when a tenant is activated.
type TenantActivated struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
}

// TenantDeactivated is the event raised when a tenant is deactivated.
type TenantDeactivated struct {
	EventVersion int
	OccurredOn   time.Time
	TenantID     TenantID
}

// TenantAdministratorRegistered is the event raised when a tenant administrator is created.
type TenantAdministratorRegistered struct {
	EventVersion      int
	OccurredOn        time.Time
	TenantID          TenantID
	TenantName        string
	Username          string
	TemporaryPassword string
	EmailAddress      EmailAddress
	AdministratorName FullName
}

// TenantProvisioningService is the domain service used to provision tenants.
type TenantProvisioningService struct {
	TenantRepo TenantRepository
	RoleRepo   RoleRepository
	UserRepo   UserRepository
}

// ProvisionTenant will provision a new tenant with supplied data.
func (tps *TenantProvisioningService) ProvisionTenant(
	tenantName,
	tenantDescription string,
	administratorName FullName,
	emailAddress EmailAddress,
	postalAddress PostalAddress,
	primaryTelephone Telephone,
	secondaryTelephone Telephone,
) (*Tenant, error) {
	tenant, err := newTenant(tenantName, tenantDescription, true)
	if err != nil {
		return nil, err
	}
	if err := tps.TenantRepo.Add(tenant); err != nil {
		return nil, err
	}
	if err := tps.registerAdministratorFor(tenant, administratorName, emailAddress, postalAddress, primaryTelephone, secondaryTelephone); err != nil {
		// Compensating transaction.
		tps.TenantRepo.Remove(tenant)
		return nil, err
	}
	return tenant, nil
}

func (tps *TenantProvisioningService) registerAdministratorFor(
	tenant *Tenant,
	administratorName FullName,
	emailAddress EmailAddress,
	postalAddress PostalAddress,
	primaryTelephone Telephone,
	secondaryTelephone Telephone,
) error {
	invitation, err := tenant.OfferInvitation("init")
	if err != nil {
		return err
	}
	temporaryPassword := password.Generate()
	person, err := newPerson(administratorName, ContactInformation{
		EmailAddress:       emailAddress,
		PostalAddress:      postalAddress,
		PrimaryTelephone:   primaryTelephone,
		SecondaryTelephone: secondaryTelephone,
	})
	if err != nil {
		return err
	}
	admin, err := tenant.RegisterUser(invitation.InvitationID, "admin", temporaryPassword, IndefiniteEnablement(), person)
	if err != nil {
		return err
	}
	if err := tps.UserRepo.Add(admin); err != nil {
		return err
	}

	role, err := tenant.ProvisionRole("Administrator", "Default "+tenant.Name+" Administrator", false)
	if err != nil {
		return err
	}
	if err := role.AssignUser(admin); err != nil {
		return err
	}

	if err := tps.RoleRepo.Add(role); err != nil {
		return err
	}

	event.Publish(TenantAdministratorRegistered{
		EventVersion:      1,
		OccurredOn:        time.Now(),
		TenantID:          tenant.ID,
		TenantName:        tenant.Name,
		Username:          admin.Username,
		TemporaryPassword: temporaryPassword,
		EmailAddress:      emailAddress,
		AdministratorName: administratorName,
	})

	return nil
}
