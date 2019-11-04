package events

// CrashReportEvent struct
type CrashReportEvent struct {
	Data      CrashReportEventData `json:"data,omitempty"`
	EventType EventType            `json:"event_type,omitempty"`
}

// NewCrashReportEvent constructor
func NewCrashReportEvent() CrashReportEvent {
	return CrashReportEvent{
		Data:      CrashReportEventData{},
		EventType: CrashReportEventType,
	}
}

func (e CrashReportEvent) data() interface{} {
	return e.Data
}

func (e CrashReportEvent) eventType() EventType {
	return e.EventType
}
