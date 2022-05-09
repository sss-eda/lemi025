package lemi025

// EventKind TODO
type EventKind string

const (
	// DatumAcquiredEventKind TODO
	DatumAcquiredEventKind EventKind = "DatumAcquired"
	// ConfigReadEventKind TODO
	ConfigReadEventKind EventKind = "ConfigRead"
	// TimeReadEventKind TODO
	TimeReadEventKind EventKind = "TimeRead"
	// TimeSetEventKind TODO
	TimeSetEventKind EventKind = "TimeSet"
)

// Event TODO
type Event interface {
	Message
	Kind() EventKind
}

// BaseEvent TODO
type BaseEvent struct {
	BaseMessage
	kind EventKind
}

// Kind TODO
func (event *BaseEvent) Kind() EventKind {
	return event.kind
}

// DatumAcquiredEvent TODO
type DatumAcquiredEvent struct{}

// Kind TODO
func (event DatumAcquiredEvent) Kind() EventKind {
	return DatumAcquiredEventKind
}

// ConfigReadEventPayload TODO
type ConfigReadEventPayload struct {
	StationType   string
	StationNumber uint8
}

// Kind TODO
func (event ConfigReadEventPayload) Kind() EventKind {
	return ConfigReadEventKind
}

// TimeReadEventPayload TODO
type TimeReadEventPayload struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// Kind TODO
func (event TimeReadEventPayload) Kind() EventKind {
	return TimeReadEventKind
}

// TimeSetEventPayload TODO
type TimeSetEventPayload struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// Kind TODO
func (event TimeSetEventPayload) Kind() EventKind {
	return TimeSetEventKind
}
