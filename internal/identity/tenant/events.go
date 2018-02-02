package tenant

import (
	"time"

	"github.com/maurofran/kit/domain"
)

const (
	tenantProvisioned = "tenant:provisioned"
	tenantActivated   = "tenant:activated"
	tenantDeactivated = "tenant:deactivated"
)

func provisioned(id ID, name string) domain.Event {
	return tenantEvent{EventType: tenantProvisioned, EventTime: time.Now(), EventVersion: 1, ID: id.Value(), Name: name}
}

func activated(id ID) domain.Event {
	return tenantEvent{EventType: tenantActivated, EventTime: time.Now(), EventVersion: 1, ID: id.Value()}
}

func deactivated(id ID) domain.Event {
	return tenantEvent{EventType: tenantDeactivated, EventTime: time.Now(), EventVersion: 1, ID: id.Value()}
}

type tenantEvent struct {
	EventType    string    `json:"_type"`
	EventTime    time.Time `json:"occurredOn"`
	EventVersion int       `json:"version"`
	ID           string    `json:"id"`
	Name         string    `json:"name,omitempty"`
}

func (ev tenantEvent) Type() string {
	return ev.EventType
}

func (ev tenantEvent) OccurredOn() time.Time {
	return ev.EventTime
}

func (ev tenantEvent) Version() int {
	return ev.EventVersion
}
