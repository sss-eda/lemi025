package lemi025

// Config TODO
type Config struct {
	StationType   string
	StationNumber uint8
}

// ReadConfigInput TODO
type ReadConfigInput struct{}

// ReadConfigStrategy TODO
type ReadConfigStrategy func(ReadConfigInput) error

// Read TODO
func (config *Config) Read(
	strategy ReadConfigStrategy,
) ReadConfigStrategy {
	return func(input ReadConfigInput) error {
		return strategy(input)
	}
}
