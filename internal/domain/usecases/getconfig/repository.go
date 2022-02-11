package getconfig

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// Repository TODO
type Repository interface {
	GetInstrumentByID(instrument.ID) (instrument.Instrument, error)
}
