package listing

// Repository TODO
type Repository interface {
	GetInstrumentByID(InstrumentID) (Instrument, error)
	GetAllInstruments() ([]Instrument, error)
	UpdateInstrument(InstrumentID, Instrument) error
}
