package events

// UserAttributeChangeEvent struct for UserAttributeChangeEvent
type UserAttributeChangeEvent struct {
	Data      UserAttributeChangeEventData `json:"data,omitempty"`
	EventType EventType                    `json:"event_type,omitempty"`
}

// NewUserAttributeChangeEvent constructor
func NewUserAttributeChangeEvent() UserAttributeChangeEvent {
	return UserAttributeChangeEvent{
		Data:      UserAttributeChangeEventData{},
		EventType: UserAttributeChangeEventType,
	}
}

func (e UserAttributeChangeEvent) data() interface{} {
	return e.Data
}

func (e UserAttributeChangeEvent) eventType() EventType {
	return e.EventType
}
