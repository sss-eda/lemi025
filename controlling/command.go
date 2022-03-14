package controlling

// ReadConfigCommand TODO
type ReadConfigCommand struct{}

// ReadTimeCommand TODO
type ReadTimeCommand struct{}

// SetTimeCommand TODO
type SetTimeCommand struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// SetCoefficients1Command TODO
type SetCoefficients1Command struct {
	Mode Mode // Yes, this would make sense outside of this module. But
	// we know that we're in the lemi025 module/context. So this is redundant.
}

// ReadCoefficients1Command TODO
type ReadCoefficients1Command struct{}

// SetCoefficients2Command TODO
type SetCoefficients2Command struct {
}

// Commands TODO
type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand | SetCoefficients1Command
}
