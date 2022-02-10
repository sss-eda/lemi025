package getting

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
	String() string
}

// Instrument TODO
type Instrument struct{}
