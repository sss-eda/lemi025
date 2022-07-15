package lemi025

type Station struct {
	number        StationNumber
	time          Time
	coefficients1 Coefficients1
	coefficients2 Coefficients2
	// gpsData       GPSData
	// systemStatus  SystemStatus
	// flashData     FLASHData
	// dacX          DAC
	// dacY          DAC
	// dacZ          DAC
}

func (station *Station) Number() (StationNumber, error) {
	return station.number, nil
}

func (station *Station) Time() (Time, error) {
	return station.time, nil
}

func (station *Station) Coefficients1() (Coefficients1, error) {
	return station.coefficients1, nil
}

func (station *Station) Coefficients2() (Coefficients2, error) {
	return station.coefficients2, nil
}
