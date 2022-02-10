package getting

// Repository TODO
type Repository interface {
	GetInstrument(InstrumentID) (Instrument, error)
}
