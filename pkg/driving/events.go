package driving

type Event interface {
	Type() EventType
}

type EventType int

const (
	ConfigRead EventType = iota
	TimeRead
	TimeSet
)

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	StationNumber uint8 `json:"stationNumber"`
}

// Type TODO
func (event ConfigReadEvent) Type() EventType {
	return ConfigRead
}

// TimeReadEvent TODO
type TimeReadEvent struct {
	Year   uint16 `json:"year"`
	Month  uint16 `json:"month"`
	Day    uint16 `json:"day"`
	Hour   uint16 `json:"hour"`
	Minute uint16 `json:"minute"`
	Second uint16 `json:"second"`
}

// Type TODO
func (event TimeReadEvent) Type() EventType {
	return TimeRead
}

// TimeSetEvent TODO
type TimeSetEvent struct {
	Year   uint16 `json:"year"`
	Month  uint16 `json:"month"`
	Day    uint16 `json:"day"`
	Hour   uint16 `json:"hour"`
	Minute uint16 `json:"minute"`
	Second uint16 `json:"second"`
}

// Type TODO
func (event TimeSetEvent) Type() EventType {
	return TimeSet
}
