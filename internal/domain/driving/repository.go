package driving

import "context"

// InstrumentRepository TODO
type InstrumentRepository interface {
	GetInstrumentByStationNumber(context.Context, StationNumber) (Instrument, error)
}
