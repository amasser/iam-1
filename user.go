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
	StartDate time.Time `bson:"startDate"`
	EndDate   time.Time `bson:"endDate"`
}
