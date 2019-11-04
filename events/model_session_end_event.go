package events

// SessionEndEvent struct for SessionEndEvent
type SessionEndEvent struct {
	Data      SessionEndEventData `json:"data,omitempty"`
	EventType EventType           `json:"event_type,omitempty"`
}

// NewSessionEndEvent constructor
func NewSessionEndEvent() SessionEndEvent {
	return SessionEndEvent{
		Data:      SessionEndEventData{},
		EventType: SessionEndEventType,
	}
}

func (e SessionEndEvent) data() interface{} {
	return e.Data
}

func (e SessionEndEvent) eventType() EventType {
	return e.EventType
}
