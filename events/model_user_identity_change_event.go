package events

// UserIdentityChangeEvent struct for UserIdentityChangeEvent
type UserIdentityChangeEvent struct {
	Data      UserIdentityChangeEventData `json:"data,omitempty"`
	EventType EventType                   `json:"event_type,omitempty"`
}

// NewUserIdentityChangeEvent constructor
func NewUserIdentityChangeEvent() UserIdentityChangeEvent {
	return UserIdentityChangeEvent{
		Data:      UserIdentityChangeEventData{},
		EventType: UserIdentityChangeEventType,
	}
}

func (e UserIdentityChangeEvent) data() interface{} {
	return e.Data
}

func (e UserIdentityChangeEvent) eventType() EventType {
	return e.EventType
}
