package lemi025

import "fmt"

// The station number of a station can't be changed
type StationNumber struct {
	value uint8
}

// func (stationNumber *StationNumber)

func (stationNumber StationNumber) String() string {
	return fmt.Sprintf("%d", stationNumber.value)
}
