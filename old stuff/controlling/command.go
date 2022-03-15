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
	Mode Mode
}

// ReadCoefficients1Command TODO
type ReadCoefficients1Command struct{}

// SetCoefficients2Command TODO
type SetCoefficients2Command struct {
	Ax1 float32
	Ay1 float32
	Az1 float32
}

// Commands TODO
type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand | SetCoefficients1Command
}
