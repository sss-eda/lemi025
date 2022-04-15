package lemi025

import (
	"fmt"
	"strconv"
)

// StationNumber - Value Object
// This struct should be always valid.
type StationNumber struct {
	sn uint8
}

// NewStationNumber - Creates a new StationNumber value object.
func NewStationNumber(
	stationNumber uint8,
) (StationNumber, error) {
	// There aren't any invalid values for the required parameters.
	return StationNumber{
		stationNumber,
	}, nil
}

// MarshalText - Serializes a StationNumber value object.
func (stationNumber StationNumber) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", stationNumber.sn)), nil
}

// UnmarshalText - Deserializes a StationNumber value object.
func (stationNumber StationNumber) UnmarshalText(data []byte) error {
	sn, err := strconv.ParseUint(string(data), 10, 8)
	if err != nil {
		// Return the error and leave VO unchanged.
		return err
	}

	stationNumber.sn = uint8(sn)

	return nil
}
