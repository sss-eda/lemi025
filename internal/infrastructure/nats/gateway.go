package nats

import (
	"time"
)

// Gateway TODO
type Gateway struct{}

// ReadConfig TODO
func (gateway *Gateway) ReadConfig() error {
	return nil
}

// ReadTime TODO
func (gateway *Gateway) ReadTime() error {
	return nil
}

func (gateway *Gateway) SetTime(t time.Time) error {
	return nil
}
