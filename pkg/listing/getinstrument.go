package listing

import (
	"context"

	"github.com/sss-eda/lemi025/pkg/instrument"
)

// GetInstrumentByID TODO
func GetInstrumentByID(
	repo Repository,
) func(context.Context, instrument.ID) (*Instrument, error) {
	return func(ctx context.Context, id instrument.ID) (*Instrument, error) {
		return repo.GetInstrumentByID(ctx, id)
	}
}
