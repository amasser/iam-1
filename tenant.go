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
	ID          string
	Description string
	StartingOn  time.Time
	Until       time.Time
}
