package lemi025

type ReadTimeCommand struct{}

type SetTimeCommand struct {
	year   uint16
	month  uint8
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

type SetCoefficients1Command struct {
	mode uint8
}

type ReadCoefficients1Command struct{}

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

type ReadCoefficients2Command struct{}

type ReadGPSDataCommand struct{}

type StopCommand struct{}

type StartCommand struct{}

type CheckFLASHCommand struct{}

type SetDACxCommand struct {
	Value int16
}

type SetDACyCommand struct {
	Value int16
}

type SetDACzCommand struct {
	Value int16
}
