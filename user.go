package iam

import (
	"time"
)

// Users is the type for a collection of users.
type Users []*User

// User is the aggregate root representing a user.
type User struct {
	TenantID   TenantID   `bson:"tenantId"`
	Username   string     `bson:"username"`
	Password   string     `bson:"password"`
	Enablement Enablement `bson:"enablement"`
	Person     *Person    `bson:"person"`
}

// UserRepository is the interace for user repository.
type UserRepository interface {
	Add(*User) error
	Update(*User) error
	Remove(*User) error
	UserWithUsername(TenantID, string) (*User, error)
	AllSimilarlyNamedUsers(TenantID, string, string) (Users, error)
}

// Enablement is the value object for a user enablement status.
type Enablement struct {
	Enabled   bool      `bson:"enabled"`
	StartDate time.Time `bson:"startDate,omitempty"`
	EndDate   time.Time `bson:"endDate,omitempty"`
}

// IsTimeExpired check if enablement time is expired.
func (e Enablement) IsTimeExpired() bool {
	if e.StartDate.IsZero() && e.EndDate.IsZero() {
		return false
	}
	now := time.Now()
	return now.Before(e.StartDate) || now.After(e.EndDate)
}

// IsEnabled will verify if enablement is actually enabled.
func (e Enablement) IsEnabled() bool {
	return e.Enabled && !e.IsTimeExpired()
}

// IndefiniteEnablement will return a new indefinite enablement.
func IndefiniteEnablement() Enablement {
	return Enablement{
		Enabled:   true,
		StartDate: time.Time{},
		EndDate:   time.Time{},
	}
}
