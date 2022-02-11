package updateconfig

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// Repository TODO
type Repository interface {
	Load(instrument.ID) (instrument.Instrument, error)
	Save(instrument.ID, instrument.Instrument) error
}
