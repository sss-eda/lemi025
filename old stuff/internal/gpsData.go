package lemi025

// GPSData TODO
type GPSData struct {
	Latitude  Latitude
	Longitude Longitude
	Altitude  Altitude
}

// Latitude TODO
type Latitude [5]byte

// Longitude TODO
type Longitude [6]byte

// Altitude TODO
type Altitude [3]byte

// ReadGPSDataCommand TODO
type ReadGPSDataCommand struct{}
