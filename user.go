package iam

import (
	"time"
)

// User is the aggregate root representing a user.
type User struct {
	ID         int64
	Version    int32
	TenantID   TenantID
	Username   string
	Password   string
	Enablement Enablement
	Person     *Person
}

// Enablement is the value object for a user enablement status.
type Enablement struct {
	Enabled   bool
	StartDate time.Time
	EndDate   time.Time
}
