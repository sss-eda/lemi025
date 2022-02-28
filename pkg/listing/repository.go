package listing

import (
	"context"

	"github.com/sss-eda/lemi025/pkg/instrument"
)

// Repository TODO
type Repository interface {
	GetInstrumentByID(context.Context, instrument.ID) (*Instrument, error)
	GetAllInstruments(context.Context) ([]*Instrument, error)
}
