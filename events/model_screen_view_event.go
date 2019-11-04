package events

// ScreenViewEvent struct for ScreenViewEvent
type ScreenViewEvent struct {
	Data      ScreenViewEventData `json:"data,omitempty"`
	EventType EventType           `json:"event_type,omitempty"`
}

// NewScreenViewEvent constructor
func NewScreenViewEvent() ScreenViewEvent {
	return ScreenViewEvent{
		Data:      ScreenViewEventData{},
		EventType: ScreenViewEventType,
	}
}

func (e ScreenViewEvent) data() interface{} {
	return e.Data
}

func (e ScreenViewEvent) eventType() EventType {
	return e.EventType
}
