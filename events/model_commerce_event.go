package events

// CommerceEvent struct
type CommerceEvent struct {
	Data      CommerceEventData `json:"data,omitempty"`
	EventType EventType         `json:"event_type,omitempty"`
}

// NewCommerceEvent constructor
func NewCommerceEvent() CommerceEvent {
	return CommerceEvent{
		Data:      CommerceEventData{},
		EventType: CommerceEventType,
	}
}

func (e CommerceEvent) data() interface{} {
	return e.Data
}

func (e CommerceEvent) eventType() EventType {
	return e.EventType
}
