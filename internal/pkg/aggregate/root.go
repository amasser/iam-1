package aggregate

// Root is the object used to provide aggregate roots.
type Root struct {
	events []interface{}
}

// Events will retrieve the events for the receiver aggregate root.
func (r *Root) Events() []interface{} {
	return r.events
}

// RegisterEvent will register a new event.
func (r *Root) RegisterEvent(event interface{}) {
	r.events = append(r.events, event)
}

// ClearEvents will remove the events.
func (r *Root) ClearEvents() {
	r.events = nil
}
