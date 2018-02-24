package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/maurofran/iam/internal/app/domain/model/event"
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
	aggregate.Root
	ID          TenantID
	Name        string
	Description string
	Active      bool
	Invitations []*Invitation
}

// NewTenant will create a new tenant with supplied data.
func NewTenant(name, description string, active bool) (*Tenant, error) {
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
	tenant.RegisterEvent(&event.TenantProvisioned{
		EventVersion: 1,
		OccurredOn:   time.Now().Unix(),
		TenantId:     string(id),
		Name:         name,
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
		t.RegisterEvent(&event.TenantActivated{
			EventVersion: 1,
			OccurredOn:   time.Now().Unix(),
			TenantId:     string(t.ID),
		})
	}
}

// Deactivate will make this tenant inactive.
func (t *Tenant) Deactivate() {
	if t.Active {
		t.Active = false
		t.RegisterEvent(&event.TenantDeactivated{
			EventVersion: 1,
			OccurredOn:   time.Now().Unix(),
			TenantId:     string(t.ID),
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
		// TODO Raise event
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
	// TODO Raise event
}

// ProvisionRole will provision a new role for this tenant.
func (t *Tenant) ProvisionRole(name, description string, supportsNesting bool) (interface{}, error) {
	if err := assertTenantActive(t); err != nil {
		return nil, err
	}
	return newRole(t.ID, name, description, supportsNesting)
	// TODO Raise event
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
