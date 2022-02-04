package lemi025

// Client TODO
type Client struct {
	Config *Config
	Time   *Time
}

// New TODO
func New() (*Client, error) {
	return &Client{}, nil
}

// ListeningStrategy TODO
type ListeningStrategy interface {
	OnConfigRead(OnConfigReadInput)
	OnTimeRead()
	OnTimeSet()
}

// OnConfigReadInput TODo
type OnConfigReadInput struct {
	StationType   string
	StationNumber uint8
}

// OnTimeReadInput TODO
type OnTimeReadInput struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// Drive TODO
// func (client *Client) Drive(
// 	strategy DriveClientStrategy,
// ) DriveClientStrategy {
// 	return
// }
