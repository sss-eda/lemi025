package lemi025

// EventPayloads TODO
type EventPayloads interface {
	DatumAcquiredEventPayload | ConfigReadEventPayload | TimeReadEventPayload | TimeSetEventPayload
}

// DatumAcquiredEventPayload TODO
type DatumAcquiredEventPayload struct{}

// ConfigReadEventPayload TODO
type ConfigReadEventPayload struct {
	StationType   string
	StationNumber uint8
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

// TimeSetEventPayload TODO
type TimeSetEventPayload struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}
