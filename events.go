package lemi025

// ConfigReadEventPayload TODO
type ConfigReadEventPayload struct {
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
