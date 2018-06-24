package iam

import (
	"reflect"
	"time"
)

// Events is a slice of events.
type Events []*Event

// Event is the base type for domain event
type Event struct {
	Version   int
	Timestamp time.Time
	Type      string
	Payload   interface{}
}

// EventWithPayload will return a new event with given payload.
func EventWithPayload(payload interface{}) *Event {
	t := reflect.TypeOf(payload)
	return &Event{
		Version:   1,
		Timestamp: time.Now(),
		Type:      t.Name(),
		Payload:   payload,
	}
}
