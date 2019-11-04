package events

// NetworkPerformanceEvent struct for NetworkPerformanceEvent
type NetworkPerformanceEvent struct {
	Data      NetworkPerformanceEventData `json:"data,omitempty"`
	EventType EventType                   `json:"event_type,omitempty"`
}

// NewNetworkPerformanceEvent constructor
func NewNetworkPerformanceEvent() NetworkPerformanceEvent {
	return NetworkPerformanceEvent{
		Data:      NetworkPerformanceEventData{},
		EventType: NetworkPerformanceEventType,
	}
}

func (e NetworkPerformanceEvent) data() interface{} {
	return e.Data
}

func (e NetworkPerformanceEvent) eventType() EventType {
	return e.EventType
}
