package lemi025

import "time"

// CommandPayloads TODO
type CommandPayloads interface {
	ReadConfigCommandPayload | ReadTimeCommandPayload | SetTimeCommandPayload
}

// ReadConfigCommandPayload TODO
type ReadConfigCommandPayload struct{}

// ReadTimeCommandPayload TODO
type ReadTimeCommandPayload struct{}

// SetTimeCommandPayload TODO
type SetTimeCommandPayload struct {
	Time time.Time
}
