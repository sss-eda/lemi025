package datapoint

import (
	"context"
)

// Add TODO
func Add(
	ctx context.Context,
	repo Repository,
	datapoint *Aggregate,
) error {
	return repo.Save(ctx, datapoint.id, datapoint)
}
