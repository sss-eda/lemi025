package lemi025

// Time TODO
type Time struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// ReadTimeCommand TODO
type ReadTimeCommand struct{}

// ReadTimeStrategy TODO
type ReadTimeStrategy func(ReadTimeCommand) error

// Read TODO - This is in the problem
func (time *Time) Read(
	strategy ReadTimeStrategy,
) func(ReadTimeCommand) error {
	return func(command ReadTimeCommand) error {
		return strategy(command)
	}
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

// OnRead TODO
func (time *Time) OnRead(
	event TimeReadEvent,
) {
	time.Year = event.Year
	time.Month = event.Month
	time.Day = event.Day
	time.Hour = event.Hour
	time.Minute = event.Minute
	time.Second = event.Second
}

// SetTimeCommand TODO
type SetTimeCommand struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// SetTimeStrategy TODO
type SetTimeStrategy func(SetTimeCommand) error

// Set TODO
func (time *Time) Set(
	strategy SetTimeStrategy,
) func(SetTimeCommand) error {
	return func(command SetTimeCommand) error {
		return strategy(command)
	}
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

// OnSet TODO
func (time *Time) OnSet(
	event TimeSetEvent,
) {
	time.Year = event.Year
	time.Month = event.Month
	time.Day = event.Day
	time.Hour = event.Hour
	time.Minute = event.Minute
	time.Second = event.Second
}

// Year TODO
type Year uint16

// Month TODO
type Month uint8

// Day TODO
type Day uint8

// Hour TODO
type Hour uint8

// Minute TODO
type Minute uint8

// Second TODO
type Second uint8
