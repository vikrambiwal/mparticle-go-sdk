package events

// Event interface
type Event interface {
	data() interface{}
	eventType() EventType
}
