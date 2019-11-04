package events

//CustomEvent struct, use NewCustomEvent() instead
type CustomEvent struct {
	Data      CustomEventData `json:"data,omitempty"`
	EventType EventType       `json:"event_type,omitempty"`
}

// NewCustomEvent constructor
func NewCustomEvent() CustomEvent {
	return CustomEvent{
		Data: CustomEventData{
			CustomEventType: OtherCustomEventType,
		},
		EventType: CustomEventEventType,
	}
}

func (e CustomEvent) data() interface{} {
	return e.Data
}

func (e CustomEvent) eventType() EventType {
	return e.EventType
}
