package iam

import (
	"fmt"
	"time"
)

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

func (e Enablement) String() string {
	return fmt.Sprintf(
		"Enablement [enabled=%t, startDate=%v, endDate=%v]",
		e.Enabled,
		e.StartDate,
		e.EndDate,
	)
}

// IndefiniteEnablement will return a new indefinite enablement.
func IndefiniteEnablement() Enablement {
	return Enablement{
		Enabled:   true,
		StartDate: time.Time{},
		EndDate:   time.Time{},
	}
}
