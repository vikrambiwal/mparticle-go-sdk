package events

// PushRegistrationEvent struct for PushRegistrationEvent
type PushRegistrationEvent struct {
	Data      PushRegistrationEventData `json:"data,omitempty"`
	EventType EventType                 `json:"event_type,omitempty"`
}

// NewPushRegistrationEvent constructor
func NewPushRegistrationEvent() PushRegistrationEvent {
	return PushRegistrationEvent{
		Data:      PushRegistrationEventData{},
		EventType: PushRegistrationEventType,
	}
}

func (e PushRegistrationEvent) data() interface{} {
	return e.Data
}

func (e PushRegistrationEvent) eventType() EventType {
	return e.EventType
}
