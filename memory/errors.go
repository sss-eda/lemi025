package memory

import "strconv"

// ErrStationAlreadyExists TODO
type ErrStationAlreadyExists struct {
	StationNumber StationNumber
}

// Error TODO
func (err ErrStationAlreadyExists) Error() string {
	return "Station " + strconv.Itoa(int(err.StationNumber)) + " already exists"
}

// ErrStationNotFound TODO
type ErrStationNotFound struct {
	StationNumber StationNumber
}

// Error TODO
func (err ErrStationNotFound) Error() string {
	return "Station " + strconv.Itoa(int(err.StationNumber)) + " not found"
}
