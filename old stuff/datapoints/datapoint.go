package datapoints

import "time"

// DataPoint TODO
type DataPoint struct {
	time time.Time
	x    int
	y    int
	z    int
}

// New TODO
func New() (*DataPoint, error) {
	return &DataPoint{
		time: time.Now(),
		x:    0,
		y:    1,
		z:    2,
	}, nil
}
