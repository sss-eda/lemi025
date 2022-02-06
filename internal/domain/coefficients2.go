package domain

// Coefficients2 TODO
type Coefficients2 struct {
	Ax1   float32
	Ay1   float32
	Az1   float32
	Beta  float32
	Gamma float32
	Xi    float32
	Exy   float32
	Eyz   float32
	Exz   float32
	K1x   float32
	K1y   float32
	K1z   float32
	K2x   float32
	K2y   float32
	K2z   float32
	KTF   float32
	KTE   float32
	KTF0  float32
	KTE0  float32
	KVBAT float32
}

// SetCoefficients2Command TODO
type SetCoefficients2Command struct {
	Ax1   float32
	Ay1   float32
	Az1   float32
	Beta  float32
	Gamma float32
	Xi    float32
	Exy   float32
	Eyz   float32
	Exz   float32
	K1x   float32
	K1y   float32
	K1z   float32
	K2x   float32
	K2y   float32
	K2z   float32
	KTF   float32
	KTE   float32
	KTF0  float32
	KTE0  float32
	KVBAT float32
}

// SetCoefficients2Strategy TODO
type SetCoefficients2Strategy func(SetCoefficients2Command) error

// Set TODO
func (coefficients2 *Coefficients2) Set(
	strategy SetCoefficients2Strategy,
) func(SetCoefficients2Command) error {
	return func(command SetCoefficients2Command) error {
		return strategy(command)
	}
}

// Coefficients2SetEvent TODO
type Coefficients2SetEvent struct{}

// OnSet TODO
func (coefficients2 *Coefficients2) OnSet(
	event Coefficients2SetEvent,
) {
}

// ReadCoefficients2Command TODO
type ReadCoefficients2Command struct {
}

// ReadCoefficients2Strategy TODO
type ReadCoefficients2Strategy func(ReadCoefficients2Command) error

// Read TODO
func (coefficients2 *Coefficients2) Read(
	strategy ReadCoefficients2Strategy,
) func(ReadCoefficients2Command) error {
	return func(command ReadCoefficients2Command) error {
		return strategy(command)
	}
}

// Coefficients2ReadEvent TODO
type Coefficients2ReadEvent struct {
	Ax1   float32
	Ay1   float32
	Az1   float32
	Beta  float32
	Gamma float32
	Xi    float32
	Exy   float32
	Eyz   float32
	Exz   float32
	K1x   float32
	K1y   float32
	K1z   float32
	K2x   float32
	K2y   float32
	K2z   float32
	KTF   float32
	KTE   float32
	KTF0  float32
	KTE0  float32
	KVBAT float32
}

// OnRead TODO
func (coefficients2 *Coefficients2) OnRead(
	event Coefficients2ReadEvent,
) {
	coefficients2.Ax1 = event.Ax1
	coefficients2.Ay1 = event.Ay1
	coefficients2.Az1 = event.Az1
	coefficients2.Beta = event.Beta
	coefficients2.Gamma = event.Gamma
	coefficients2.Xi = event.Xi
	coefficients2.Exy = event.Exy
	coefficients2.Eyz = event.Eyz
	coefficients2.Exz = event.Exz
	coefficients2.K1x = event.K1x
	coefficients2.K1y = event.K1y
	coefficients2.K1z = event.K1z
	coefficients2.K2x = event.K2x
	coefficients2.K2y = event.K2y
	coefficients2.K2z = event.K2z
	coefficients2.KTF = event.KTF
	coefficients2.KTE = event.KTE
	coefficients2.KTF0 = event.KTF0
	coefficients2.KTE0 = event.KTE0
	coefficients2.KVBAT = event.KVBAT
}
