package listing

import (
	"context"
)

// GetInstrumentByID TODO
func GetInstrumentByID(
	repo Repository,
) func(context.Context, InstrumentID) (*Instrument, error) {
	return func(ctx context.Context, id InstrumentID) (*Instrument, error) {
		return repo.GetInstrumentByID(ctx, id)
	}
}
