package lemi025

import "time"

// Commands TODO
type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand
}

// ReadConfigCommand TODO
type ReadConfigCommand struct{}

// ReadTimeCommand TODO
type ReadTimeCommand struct{}

// SetTimeCommand TODO
type SetTimeCommand struct {
	Time time.Time
}
