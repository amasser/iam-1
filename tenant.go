package iam

import (
	"time"
)

// TenantID is the value object for a tenant identifier.
type TenantID string

// Tenant is the aggregate root object for the tenant.
type Tenant struct {
	ID          TenantID
	Name        string
	Description string
	Active      bool
	Invitations Invitations
}

// Invitations is the collection of invitation entities.
type Invitations []*Invitation

// Invitation is the entity for an invitation.
type Invitation struct {
	ID          string
	Description string
	StartingOn  time.Time
	Until       time.Time
}
