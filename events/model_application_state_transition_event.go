package events

// ApplicationStateTransitionEvent struct
type ApplicationStateTransitionEvent struct {
	Data      ApplicationStateTransitionEventData `json:"data,omitempty"`
	EventType string                              `json:"event_type,omitempty"`
}
