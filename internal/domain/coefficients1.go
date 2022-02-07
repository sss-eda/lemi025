package domain

// Mode TODO
type Mode int

const (
	// FLASH TODO
	FLASH Mode = iota
	// PC TODO
	PC
	// FLASHandPC TODO
	FLASHandPC
)

// Mode1 TODO
type Mode1 int

const (
	// MENU TODO
	MENU Mode1 = iota
	// RECORDING TODO
	RECORDING
)

// Coefficients1 TODO
type Coefficients1 struct {
	Mode  Mode
	Uin   uint16
	Mode1 Mode1
}

// SetCoefficients1Command TODO
type SetCoefficients1Command struct {
	Mode Mode
}

// SetCoefficients1Strategy TODO
type SetCoefficients1Strategy func(SetCoefficients1Command) error

// Set TODO
func (coefficients1 *Coefficients1) Set(
	strategy SetCoefficients1Strategy,
) func(SetCoefficients1Command) error {
	return func(command SetCoefficients1Command) error {
		return strategy(command)
	}
}

// Coefficients1SetEvent TODO
type Coefficients1SetEvent struct {
	Mode Mode
}

// OnSet TODO
func (coefficients1 *Coefficients1) OnSet(
	event Coefficients1SetEvent,
) {
	coefficients1.Mode = event.Mode
}

// ReadCoefficients1Command TODO
type ReadCoefficients1Command struct {
}

// ReadCoefficients1Strategy TODO
type ReadCoefficients1Strategy func(ReadCoefficients1Command) error

// Read TODO
func (coefficients1 *Coefficients1) Read(
	strategy ReadCoefficients1Strategy,
) func(ReadCoefficients1Command) error {
	return func(command ReadCoefficients1Command) error {
		return strategy(command)
	}
}

// Coefficients1ReadEvent TODO
type Coefficients1ReadEvent struct {
	Mode  Mode
	Uin   uint16
	Mode1 Mode1
}

// OnRead TODO
func (coefficients1 *Coefficients1) OnRead(
	event Coefficients1ReadEvent,
) {
	coefficients1.Mode = event.Mode
	coefficients1.Uin = event.Uin
	coefficients1.Mode1 = event.Mode1
}
