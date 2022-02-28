package datapoints

import "context"

// Repository TODO
type Repository interface {
	Load(context.Context, ID) (*DataPoint, error)
	Save(context.Context, ID, *DataPoint) error
}
