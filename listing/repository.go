package listing

import (
	"context"
)

// Repository TODO
type Repository interface {
	GetInstrumentByID(context.Context, InstrumentID) (*Instrument, error)
	GetAllInstruments(context.Context) ([]*Instrument, error)
}
