package lemi025

// InstrumentRepository TODO
type InstrumentRepository interface {
	Load(InstrumentID) (Instrument, error)
	Save(InstrumentID, Instrument) error
}
