package datapoint

import "context"

// Repository TODO
type Repository interface {
	Load(context.Context, ID) (*Aggregate, error)
	Save(context.Context, ID, *Aggregate) error
}
