package serial

import "time"

// Gateway TODO
type Gateway struct{}

// NewGateway TODO
func NewGateway() (*Gateway, error) {
	return &Gateway{}, nil
}

// ReadConfig TODO
func (gateway *Gateway) ReadConfig() error {
	return nil
}

// ReadTime TODO
func (gateway *Gateway) ReadTime() error {
	return nil
}

// SetTime TODO
func (gateway *Gateway) SetTime(time.Time) error {
	return nil
}
