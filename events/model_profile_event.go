package events

// ProfileEvent struct for ProfileEvent
type ProfileEvent struct {
	Data      ProfileEventData `json:"data,omitempty"`
	EventType EventType        `json:"event_type,omitempty"`
}

// NewProfileEvent constructor
func NewProfileEvent() ProfileEvent {
	return ProfileEvent{
		Data:      ProfileEventData{},
		EventType: ProfileEventType,
	}
}

func (e ProfileEvent) data() interface{} {
	return e.Data
}

func (e ProfileEvent) eventType() EventType {
	return e.EventType
}
