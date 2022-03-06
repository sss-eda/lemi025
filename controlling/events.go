package controlling

// Events TODO
type Events interface {
	ConfigReadEvent | TimeReadEvent | TimeSetEvent | DataFrameAcquiredEvent
}

// // EventHandlerFunc TODO
// type EventHandlerFunc[Event Events] func(*Event)

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	StationNumber uint8
}

// TimeReadEvent TODO
type TimeReadEvent struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// TimeSetEvent TODO
type TimeSetEvent struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// DataFrameAcquiredEvent TODO
type DataFrameAcquiredEvent struct{}
