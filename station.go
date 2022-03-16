package lemi025

import "time"

// Station TODO
type Station interface {
	ReadConfig() error
	ReadTime() error
	SetTime(time.Time) error
}

// NewStation TODO
func NewStation(
	adapterType string,
) (Station, error) {
	var station Station
	switch adapterType {
	case "serial":
		station = &serial.Station{}
case "":
}

// StationNumber TODO
type StationNumber uint8

// StationRepository TODO
type StationRepository interface {
	RegisterStation(StationNumber, Station) error
	GetStationByNumber(StationNumber) (Station, error)
}

// Transmitter TODO
type Transmitter interface {
	ReadConfig(ReadConfigCommand) error
	ReadTime(ReadTimeCommand) error
}

// Receiver TODO
type Receiver interface {
	DataFrameAcquired(DataFrameAcquiredEvent) error
	ConfigRead(ConfigReadEvent) error
	TimeRead(TimeReadEvent) error
}

// Lister TODO
type Lister interface {
}
