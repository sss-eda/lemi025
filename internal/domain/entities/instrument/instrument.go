package instrument

// Instrument TODO
type Instrument struct {
	ID            ID
	Config        Config
	Time          Time
	Coefficients1 Coefficients1
	Coefficients2 Coefficients2
	// GPSData       *GPSData
	// FLASHData     *FLASHData
	// DACData       *DACData
	changes []Event
}

// Repository TODO
type Repository interface {
	Load(ID) (*Instrument, error)
	Save(ID, Event) error
}

// New TODO
func New() (Instrument, error) {
	return Instrument{}, nil
}

// Raise TODO
func (instrument Instrument) Raise(event Event) {
	// event.apply(instrument)
	instrument.changes = append(instrument.changes, event)
}

// // ListeningStrategy TODO
// type ListeningStrategy interface {
// 	OnConfigRead(OnConfigReadInput)
// 	OnTimeRead()
// 	OnTimeSet()
// }

// // OnConfigReadInput TODo
// type OnConfigReadInput struct {
// 	StationType   string
// 	StationNumber uint8
// }

// // OnTimeReadInput TODO
// type OnTimeReadInput struct {
// 	Year   uint16
// 	Month  uint8
// 	Day    uint8
// 	Hour   uint8
// 	Minute uint8
// 	Second uint8
// }

// Drive TODO
// func (client *Client) Drive(
// 	strategy DriveClientStrategy,
// ) DriveClientStrategy {
// 	return
// }
