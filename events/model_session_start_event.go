package events

// SessionStartEvent struct for SessionStartEvent
type SessionStartEvent struct {
	Data      SessionStartEventData `json:"data,omitempty"`
	EventType EventType             `json:"event_type,omitempty"`
}

// NewSessionStartEvent constructor
func NewSessionStartEvent() SessionStartEvent {
	return SessionStartEvent{
		Data:      SessionStartEventData{},
		EventType: SessionStartEventType,
	}
}

func (e SessionStartEvent) data() interface{} {
	return e.Data
}

func (e SessionStartEvent) eventType() EventType {
	return e.EventType
}
