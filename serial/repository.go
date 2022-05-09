package serial

import "fmt"

type repository struct {
	stations map[string]*Station
}

// LoadStation TODO
func (repo *repository) LoadStation(
	stationID string,
) (*Station, error) {
	station, ok := repo.stations[stationID]
	if !ok {
		return nil, fmt.Errorf("station %s not found", stationID)
	}

	return station, nil
}

// SaveStation
func (repo *repository) SaveStation(
	stationID string,
	station *Station,
) error {
	repo.stations[stationID] = station

	return nil
}
