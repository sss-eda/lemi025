package lemi025

type Station interface {
	// ReadConfig(ReadConfigCommandPayload) error
	// ReadTime(ReadTimeCommandPayload) error
	// SetTime(SetTimeCommandPaylaod) error
	OnDataFrameAcquired(func(DataFrameAcquiredEvent))
	OnConfigRead(func(ConfigReadEvent))
	OnTimeRead(func(TimeReadEvent))
}

// // Station TODO
// type Station struct {
// 	Number         StationNumber
// 	Time           Time
// 	Coefficients1  Coefficients1
// 	Coefficients2  Coefficients2
// 	GPSCoordinates GPSCoordinates
// 	Recording      SystemStatus
// 	FLASHStatus    FLASHStatus
// 	DACx           DAC
// 	DACy           DAC
// 	DACz           DAC
// 	// history        []Event
// }

// // NewStation TODO
// func NewStation(adapter Adapter) (*Station, error) {
// 	station := Station{}

// 	err := adapter.ReadConfig()
// 	if err != nil {
// 		return err
// 	}

// 	return &station, nil
// }

// // // NewStationFromEvents TODO
// // func NewStationFromEvents(
// // 	events []Event,
// // ) (*Station, error) {
// // 	station, err := NewStation()
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	for _, event := range events {
// // 		station.Apply(event)
// // 	}

// // 	return station, nil
// // }

// // ReadConfig TODO
// func (station *Station) ReadConfig() error {
// 	if station.Recording {
// 		return fmt.Errorf("station is recording")
// 	}

// 	event := ReadConfigCommandReceived{}

// 	return gateway()
// }

// // SetTime TODO
// func (station *Station) SetTime(handle func(SetTimeCommandPayload) error) error {
// 	return nil
// }

// // StationDriver TODO
// type StationDriver interface {
// 	ReadConfig() error
// 	ReadTime() error
// 	SetTime(Time) error
// }

// // StationsService TODO
// type StationsService interface {
// }
