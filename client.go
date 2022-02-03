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
	OnConfigRead()
	OnTimeRead()
	OnTimeSet()
}

// Drive TODO
// func (client *Client) Drive(
// 	strategy DriveClientStrategy,
// ) DriveClientStrategy {
// 	return
// }
