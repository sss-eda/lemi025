package listing

// InstrumentID TODO
type InstrumentID interface {
	Equals(InstrumentID) bool
	String() string
}

// Instrument TODO
type Instrument struct {
	Config        Config
	Time          Time
	Coefficients1 Coefficients1
	Coefficients2 Coefficients2
	GPS           GPSData
	System        System
	FLASH         FLASHData
	DAC           DAC
}

// Config TODO
type Config struct {
	StationNumber uint8
}

// Time - We could change this to a time.Time thing, since this is read model
// and not something we have to validate.
type Time struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// Coefficients1 TODO
type Coefficients1 struct {
	Mode  uint8
	Uin   uint8
	Mode1 uint8
}

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

// GPSData TODO
type GPSData struct {
	Latitude  [5]byte
	Longitude [6]byte
	Altitude  [3]byte
}

// System TODO
type System bool

// FLASHData TODO
type FLASHData struct {
	Size uint16
	Free uint8
}

// DAC TODO
type DAC struct {
	X int16
	Y int16
	Z int16
}
