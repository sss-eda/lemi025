package instrument

// Config TODO
type Config struct {
	StationType   string
	StationNumber StationNumber
}

// Update TODO
func (config Config) Read(
	NewStationType string,
	NewStationNumber StationNumber,
) (ConfigReadEvent, error) {
	return ConfigReadEvent{

		NewStationNumber: NewStationNumber,
	}, nil
}

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	InstrumentID     ID
	NewStationNumber StationNumber
}

// // ReadConfigStrategy TODO
// type ReadConfigStrategy func(ReadConfigCommand) error

// // Read TODO
// func (config *Config) Read(
// 	strategy ReadConfigStrategy,
// ) ReadConfigStrategy {
// 	return func(command ReadConfigCommand) error {
// 		return strategy(command)
// 	}
// }

// // ConfigReadEvent TODO
// type ConfigReadEvent struct {
// 	StationNumber uint8
// }

// // OnRead TODO
// func (config *Config) OnRead(
// 	event ConfigReadEvent,
// ) {
// 	config.StationNumber = event.StationNumber
// }
