package updateconfig

import "github.com/sss-eda/lemi025/internal/domain/entities/instrument"

// Request TODO
type Request struct {
	InstrumentID instrument.ID
	NewConfig    instrument.Config
}
