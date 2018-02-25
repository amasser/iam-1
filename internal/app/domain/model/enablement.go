package model

import (
	"errors"
	"time"
)

// Enablement is the value object used to provide data about User enable/disable status.
//
// Its default value (a disabled indefinite enablement) is a valid value object and can be used as is.
type Enablement struct {
	Enabled   bool      `bson:"enabled"`
	StartDate time.Time `bson:"startDate,omitempty"`
	EndDate   time.Time `bson:"endDate,omitempty"`
}

// IndefiniteEnablement will make a new indefinite enabled-status enablement.
func IndefiniteEnablement() Enablement {
	return Enablement{Enabled: true}
}

// MakeEnablement will make a new enablement with supplied status and start and end time.
// The function will fail with an error if the startDate is greater than endDate or if only one of startDate
// and endDate are provided.
func MakeEnablement(enabled bool, startDate, endDate time.Time) (Enablement, error) {
	if startDate.After(endDate) {
		return Enablement{}, errors.New("Invalid startDate and endDate provided")
	}
	return Enablement{enabled, startDate, endDate}, nil
}

// ActuallyEnabled will check if the enablement is actually enabled.
func (e Enablement) ActuallyEnabled() bool {
	return e.Enabled && !e.TimeExpired()
}

// TimeExpired will check if the validity time of this enablement is expired.
func (e Enablement) TimeExpired() bool {
	expired := false
	if !e.StartDate.IsZero() && !e.EndDate.IsZero() {
		now := time.Now()
		expired = now.After(e.StartDate) && now.Before(e.EndDate)
	}
	return expired
}
