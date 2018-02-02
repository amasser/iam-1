package tenant

import (
	"github.com/google/uuid"
	"github.com/maurofran/kit/assert"
	"github.com/maurofran/kit/domain"
)

// Tenant is the aggregate root object used to manage a tenant.
type Tenant struct {
	*domain.AggregateRoot
	ID          ID
	Name        string
	Description string
	Active      bool
	Invitations []*Invitation
}

// New is used to create a new tenant instance.
// It returns an erro if any of supplied parameter does not comply with restrictions.
func New(id ID, name, description string, active bool) (*Tenant, error) {
	if err := assert.NotZero(id, "id"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(name, "name"); err != nil {
		return nil, err
	}
	if err := assert.MaxLength(name, 70, "name"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(description, "description"); err != nil {
		return nil, err
	}
	tenant := &Tenant{
		AggregateRoot: &domain.AggregateRoot{},
		ID:            id,
		Name:          name,
		Description:   description,
		Active:        active,
	}
	tenant.RegisterEvent(provisioned(id, name))
	return tenant, nil
}

// Activate this tenant if not active.
func (t *Tenant) Activate() {
	if !t.Active {
		t.Active = true
		t.RegisterEvent(activated(t.ID))
	}
}

// Deactivate this tenant if active.
func (t *Tenant) Deactivate() {
	if t.Active {
		t.Active = false
		t.RegisterEvent(deactivated(t.ID))
	}
}

// AllAvailableInvitations will retrieve a slice of available invitation descriptors.
func (t *Tenant) AllAvailableInvitations() ([]InvitationDescriptor, error) {
	if err := assert.State(t.Active, "active"); err != nil {
		return nil, err
	}
	var result []InvitationDescriptor
	for _, invitation := range t.Invitations {
		if invitation.IsAvailable() {
			result = append(result, invitation.toDescriptor(t.ID))
		}
	}
	return result, nil
}

// AllUnavailableInvitations will retrieve a slice of unavailable invitation descriptors.
func (t *Tenant) AllUnavailableInvitations() ([]InvitationDescriptor, error) {
	if err := assert.State(t.Active, "active"); err != nil {
		return nil, err
	}
	var result []InvitationDescriptor
	for _, invitation := range t.Invitations {
		if !invitation.IsAvailable() {
			result = append(result, invitation.toDescriptor(t.ID))
		}
	}
	return result, nil
}

// IsInvitationAvailableThrough will check if an invitation is available.
func (t *Tenant) IsInvitationAvailableThrough(invitationIdentifier string) (bool, error) {
	if err := assert.State(t.Active, "active"); err != nil {
		return false, err
	}
	idx, _ := t.invitationFor(invitationIdentifier)
	return idx >= 0, nil
}

// OfferInvitation will offer a new invitation and return to the caller.
func (t *Tenant) OfferInvitation(description string) (*Invitation, error) {
	if err := assert.State(t.Active, "active"); err != nil {
		return nil, err
	}
	exists, err := t.IsInvitationAvailableThrough(description)
	if err != nil {
		return nil, err
	}
	if err := assert.StateNot(exists, "description exists"); err != nil {
		return nil, err
	}
	invitation, err := newInvitation(uuid.New().String(), description)
	if err != nil {
		return nil, err
	}
	invitations := append(t.Invitations, invitation)
	t.Invitations = invitations
	return invitation, nil
}

// WithdrawInvitation will remove the invitation identified by supplied identifier
func (t *Tenant) WithdrawInvitation(identifier string) error {
	if err := assert.State(t.Active, "active"); err != nil {
		return err
	}
	if idx, _ := t.invitationFor(identifier); idx >= 0 {
		var inv []*Invitation
		copy(inv, t.Invitations)
		inv[idx] = inv[len(inv)-1]
		inv[len(inv)-1] = nil
		t.Invitations = inv[:len(inv)-1]
	}
	return nil
}

// GetInvitation will find the invitation for supplied identifier.
func (t *Tenant) GetInvitation(identifier string) (*Invitation, error) {
	if err := assert.State(t.Active, "active"); err != nil {
		return nil, err
	}
	_, invitation := t.invitationFor(identifier)
	return invitation, nil
}

func (t *Tenant) invitationFor(identifier string) (int, *Invitation) {
	for idx, invitation := range t.Invitations {
		if invitation.IsIdentifiedBy(identifier) {
			return idx, invitation
		}
	}
	return -1, nil
}
