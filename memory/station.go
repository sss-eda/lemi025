package memory

import "github.com/sss-eda/lemi025"

// StationNumber TODO
type StationNumber uint8

// Station TODO
type Station struct {
	Time           lemi025.Time
	Coefficients1  lemi025.Coefficients1
	Coefficients2  lemi025.Coefficients2
	GPSCoordinates lemi025.GPSCoordinates
	SystemStatus   lemi025.SystemStatus
	FLASHStatus    lemi025.FLASHStatus
	DACx           lemi025.DAC
	DACy           lemi025.DAC
	DACz           lemi025.DAC
}

// func (station *Station) ToLemi025() *lemi025.Station {
// 	return &lemi025.Station{
// 		Name: station.Name,
// 		Baud: station.Baud,
// 	}
// }

// StationRepository TODO
type StationRepository struct {
	stations map[StationNumber]*Station
}

// NewStationRepository TODO
func NewStationRepository() (*StationRepository, error) {
	return &StationRepository{}, nil
}

// Add TODO
func (repository *StationRepository) Add(station *lemi025.Station) error {
	_, ok := repository.stations[StationNumber(station.Number)]
	if ok {
		return ErrStationAlreadyExists{StationNumber(station.Number)}
	}

	repository.stations[StationNumber(station.Number)] = &Station{
		Time:           station.Time,
		Coefficients1:  station.Coefficients1,
		Coefficients2:  station.Coefficients2,
		GPSCoordinates: station.GPSCoordinates,
		SystemStatus:   station.SystemStatus,
		FLASHStatus:    station.FLASHStatus,
		DACx:           station.DACx,
		DACy:           station.DACy,
		DACz:           station.DACz,
	}

	return nil
}

// FindByNumber TODO
func (repository *StationRepository) FindByNumber(number lemi025.StationNumber) (*Station, error) {
	station, ok := repository.stations[StationNumber(number)]
	if !ok {
		return nil, ErrStationNotFound{StationNumber(number)}
	}

	return station, nil
}
