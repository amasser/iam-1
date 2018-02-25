package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/maurofran/kit/assert"
	"github.com/pkg/errors"
)

// Invitation is the entity object for tenant invitations.
type Invitation struct {
	InvitationID string    `bson:"invitationId"`
	Description  string    `bson:"description"`
	StartingOn   time.Time `bson:"startingOn,omitempty"`
	Until        time.Time `bson:"until,omitempty"`
}

func newInvitation(description string) (*Invitation, error) {
	if err := assert.NotEmpty(description, "description"); err != nil {
		return nil, err
	}
	randomID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate random invitationID")
	}
	return &Invitation{InvitationID: randomID.String(), Description: description}, nil
}

// Available will check if the receiver invitation is available.
func (i *Invitation) Available() bool {
	if i.StartingOn.IsZero() && i.Until.IsZero() {
		return true
	}
	now := time.Now()
	return now.After(i.StartingOn) && now.Before(i.Until)
}

func (i *Invitation) identifiedBy(descriptor string) bool {
	return i.InvitationID == descriptor || i.Description == descriptor
}

func (i *Invitation) redefineAsOpenEnded() {
	i.StartingOn = time.Time{}
	i.Until = time.Time{}
}

func (i *Invitation) redefineAs(startingOn, until time.Time) error {
	if err := assert.NotZero(startingOn, "startingOn"); err != nil {
		return err
	}
	if err := assert.NotZero(until, "until"); err != nil {
		return err
	}
	if err := assert.Condition(startingOn.Before(until), "startingOn after until"); err != nil {
		return err
	}
	i.StartingOn = startingOn
	i.Until = until
	return nil
}

func (i *Invitation) toDescriptor(tenantID TenantID) InvitationDescriptor {
	return InvitationDescriptor{
		TenantID:     tenantID,
		InvitationID: i.InvitationID,
		Description:  i.Description,
		StartingOn:   i.StartingOn,
		Until:        i.Until,
	}
}

// InvitationDescriptor is a value object used to provide a read-only version of invitation.
type InvitationDescriptor struct {
	TenantID     TenantID
	InvitationID string
	Description  string
	StartingOn   time.Time
	Until        time.Time
}
