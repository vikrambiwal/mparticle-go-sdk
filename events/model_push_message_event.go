package events

// PushMessageEvent struct for PushMessageEvent
type PushMessageEvent struct {
	Data      PushMessageEventData `json:"data,omitempty"`
	EventType EventType            `json:"event_type,omitempty"`
}

// NewPushMessageEvent constructor
func NewPushMessageEvent() PushMessageEvent {
	return PushMessageEvent{
		Data:      PushMessageEventData{},
		EventType: PushMessageEventType,
	}
}

func (e PushMessageEvent) data() interface{} {
	return e.Data
}

func (e PushMessageEvent) eventType() EventType {
	return e.EventType
}
