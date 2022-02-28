package datapoint

// Aggregate TODO
type Aggregate struct {
	id            ID
	StationNumber uint8
	Time          Time
	DataTF        int16
	DataTE        int16
	DACx          int16
	DACy          int16
	DACz          int16
	BxDAC         int16
	ByDAC         int16
	BzDAC         int16
	Readings      [10]Reading
	Mode          uint8
	FlashFree     uint8
	VoltageUIN    uint8
	StatusGPS     uint8
	CheckSum      byte
}

// Add TODO
func (aggregate *Aggregate) Add() AddedEvent {
	return AddedEvent{
		DataPoint: aggregate,
	}
}
