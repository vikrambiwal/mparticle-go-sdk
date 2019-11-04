package events

// BreadcrumbEvent struct
type BreadcrumbEvent struct {
	Data      BreadcrumbEventData `json:"data,omitempty"`
	EventType EventType           `json:"event_type,omitempty"`
}

// NewBreadcrumbEvent constructor
func NewBreadcrumbEvent() BreadcrumbEvent {
	return BreadcrumbEvent{
		Data:      BreadcrumbEventData{},
		EventType: BreadcrumbEventType,
	}
}

func (e BreadcrumbEvent) data() interface{} {
	return e.Data
}

func (e BreadcrumbEvent) eventType() EventType {
	return e.EventType
}
