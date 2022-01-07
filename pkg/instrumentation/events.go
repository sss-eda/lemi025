package instrumentation

// TimeSynchronisedEvent TODO
type TimeSynchronisedEvent struct {
	Year   uint16 `json:"year"`
	Month  uint8  `json:"month"`
	Day    uint8  `json:"day"`
	Hour   uint8  `json:"hour"`
	Minute uint8  `json:"minute"`
	Second uint8  `json:"second"`
}

// Type TODO
func (event *TimeSynchronisedEvent) Type() string {
	return "Time Synchronised"
}
