package lemi025

// Event TODO
type Event interface {
	Data() []byte
}

type DataFrameAcquiredEvent struct {
	stationNumber uint8
	year          uint8
	month         uint8
	day           uint8
	hour          uint8
	minute        uint8
	second        uint8
	tf            int16
	te            int16
	dacX          int16
	dacY          int16
	dacZ          int16
	bxDAC         int16
	byDAC         int16
	bzDAC         int16
	readings      [10]struct {
		dataBx float32
		dataBy float32
		dataBz float32
	}
	mode       uint8
	flashFree  uint8
	voltageUin uint8
	statusGPS  byte
	checksum   byte
}

// // Type TODO
// func (event *DataFrameAcquiredEvent) Type() string {
// 	return "DataFrameAcquired"
// }

type ConfigReadEvent struct {
	stationNumber uint8
}

type TimeReadEvent struct {
	year   uint16
	month  uint8
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

type TimeSetEvent struct {
	year   uint16
	month  uint8
	day    uint8
	hour   uint8
	minute uint8
	second uint8
}

type Coefficients1SetEvent struct {
	mode uint8
}

type Coefficients1ReadEvent struct {
	mode  uint8
	uin   uint8
	mode1 uint8
}

type Coefficients2SetEvent struct{}

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

type GPSDataReadEvent struct {
	Latitude  [5]byte
	Longitude [6]byte
	Altitude  [3]byte
}

type StartedEvent struct{}

type FLASHCheckedEvent struct {
	size uint16
	free uint8
}
