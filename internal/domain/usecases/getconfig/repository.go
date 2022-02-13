package getconfig

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// Repository TODO
type Repository interface {
	GetConfigByInstrumentID(instrument.ID) (instrument.Config, error)
}
