package iam

import (
	"time"
)

// TenantID is the value object for a tenant identifier.
type TenantID string

// Tenant is the aggregate root object for the tenant.
type Tenant struct {
	ID          TenantID    `bson:"tenantId"`
	Name        string      `bson:"name"`
	Description string      `bson:"description,omitempty"`
	Active      bool        `bson:"active"`
	Invitations Invitations `bson:"invitations"`
}

// TenantRepository is the interface for tenants.
type TenantRepository interface {
	Add(*Tenant) error
	Update(*Tenant) error
	Remove(*Tenant) error
	TenantNamed(string) (*Tenant, error)
	TenantOfID(TenantID) (*Tenant, error)
}

// Invitations is the collection of invitation entities.
type Invitations []*Invitation

// Invitation is the entity for an invitation.
type Invitation struct {
	ID          string    `bson:"invitationId"`
	Description string    `bson:"description"`
	StartingOn  time.Time `bson:"startingOn"`
	Until       time.Time `bson:"until"`
}
