package lemi025

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
}

// Instrument TODO
type Instrument interface {
	ReadConfig() error
	ReadTime() error
	SetTime(Time) error
}
