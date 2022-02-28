package datapoints

// API TODO
type API interface {
	AddDataPoint(ID, *DataPoint) error
}
