package domain

// ReadConfig TODO
type ReadConfig func(ReadConfigInput) error

// ReadConfigInput TODO
type ReadConfigInput struct {
	InstrumentID
}
