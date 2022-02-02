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

// ReadTimeInput TODO
type ReadTimeInput struct{}

// ReadTimeStrategy TODO
type ReadTimeStrategy func(ReadTimeInput) error

// Read TODO
func (time *Time) Read(
	strategy ReadTimeStrategy,
) func(ReadTimeInput) error {
	return func(input ReadTimeInput) error {
		return strategy(input)
	}
}

// SetTimeInput TODO
type SetTimeInput struct{}

// SetTimeStrategy TODO
type SetTimeStrategy func(SetTimeInput) error

// Set TODO
func (time *Time) Set(
	strategy SetTimeStrategy,
) func(SetTimeInput) error {
	return func(input SetTimeInput) error {
		return strategy(input)
	}
}
