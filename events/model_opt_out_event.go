package events

// OptOutEvent struct for OptOutEvent
type OptOutEvent struct {
	Data      OptOutEventData `json:"data,omitempty"`
	EventType EventType       `json:"event_type,omitempty"`
}

// NewOptOutEvent constructor
func NewOptOutEvent() OptOutEvent {
	return OptOutEvent{
		Data:      OptOutEventData{},
		EventType: OptOutEventType,
	}
}

func (e OptOutEvent) data() interface{} {
	return e.Data
}

func (e OptOutEvent) eventType() EventType {
	return e.EventType
}
