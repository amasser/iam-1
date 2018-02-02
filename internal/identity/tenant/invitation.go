package tenant

import (
	"fmt"
	"time"

	"github.com/maurofran/kit/assert"
)

// Invitation is the entity object used to supply an invitation to a tenant.
type Invitation struct {
	InvitationID string    `bson:"invitationId"`
	Description  string    `bson:"description"`
	StartingOn   time.Time `bson:"startingOn,omitempty"`
	Until        time.Time `bson:"until,omitempty"`
}

// newInvitation is used to create a new Invitation object.
// It return an error if either invitationID or description are an empty string
func newInvitation(invitationID, description string) (*Invitation, error) {
	if err := assert.NotEmpty(invitationID, "invitationID"); err != nil {
		return nil, err
	}
	if err := assert.NotEmpty(description, "description"); err != nil {
		return nil, err
	}
	return &Invitation{
		InvitationID: invitationID,
		Description:  description,
	}, nil
}

// IsAvailable will check if receiver invitation is actually available.
func (i *Invitation) IsAvailable() bool {
	if i.StartingOn.IsZero() && i.Until.IsZero() {
		return true
	}
	now := time.Now().Unix()
	return i.StartingOn.Unix() <= now && now <= i.Until.Unix()
}

// IsIdentifiedBy will check if receiver invitation is identified by supplied identifier.
func (i *Invitation) IsIdentifiedBy(invitationIdentifier string) bool {
	return i.InvitationID == invitationIdentifier || i.Description == invitationIdentifier
}

// RedefineAsOpenEnded will change the invitation in order to not have a starting and until date.
func (i *Invitation) RedefineAsOpenEnded() {
	i.StartingOn = time.Time{}
	i.Until = time.Time{}
}

// RedefineAs will change the invitation in order to make it available inside the supplied period.
// It returns an error if startingOn or until time are zero or if startingOn is greater than until.
func (i *Invitation) RedefineAs(startingOn, until time.Time) error {
	if err := assert.NotZero(startingOn, "startingOn"); err != nil {
		return err
	}
	if err := assert.NotZero(until, "until"); err != nil {
		return err
	}
	if err := assert.Condition(startingOn.Before(until), "startingOn must occurs before until"); err != nil {
		return err
	}
	i.StartingOn = startingOn
	i.Until = until
	return nil
}

func (i *Invitation) toDescriptor(tenantID ID) InvitationDescriptor {
	return InvitationDescriptor{
		TenantID:     tenantID,
		InvitationID: i.InvitationID,
		Description:  i.Description,
		StartingOn:   i.StartingOn,
		Until:        i.Until,
	}
}

func (i *Invitation) String() string {
	return fmt.Sprintf(
		"Invitation [InvitationID=%s, Description=%s, StartingOn=%s, Until=%s]",
		i.InvitationID,
		i.Description,
		i.StartingOn,
		i.Until,
	)
}

// InvitationDescriptor is a value object used to provide invitations as read-only objects.
type InvitationDescriptor struct {
	TenantID     ID
	InvitationID string
	Description  string
	StartingOn   time.Time
	Until        time.Time
}
