package receiving

import (
	"context"
	"time"
)

// MessageTypes TODO
type MessageTypes int

const (
	// ConfigReadEventMessage TODO
	ConfigReadEventMessage MessageTypes = iota
	// TimeReadEventMessage TODO
	TimeReadEventMessage
)

// Events TODO
type Events interface {
	ConfigReadEvent | TimeReadEvent
}

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	StationNumber uint8
}

// Apply TODO
func (event ConfigReadEvent) Apply(
	ctx context.Context,
	config *Config,
) {
	config.StationNumber = event.StationNumber
}

// TimeReadEvent TODO
type TimeReadEvent struct {
	Time time.Time
}
