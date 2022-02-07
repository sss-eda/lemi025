package domain

// Config TODO
type Config struct {
	StationNumber uint8
}

// ReadConfigStrategy TODO
type ReadConfigStrategy func(ReadConfigCommand) error

// Read TODO
func (config *Config) Read(
	strategy ReadConfigStrategy,
) ReadConfigStrategy {
	return func(command ReadConfigCommand) error {
		return strategy(command)
	}
}

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	StationNumber uint8
}

// OnRead TODO
func (config *Config) OnRead(
	event ConfigReadEvent,
) {
	config.StationNumber = event.StationNumber
}
