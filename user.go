package iam

import (
	"time"
)

// Users is the type for a collection of users.
type Users []*User

// User is the aggregate root representing a user.
type User struct {
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
