package driver

import (
	"fmt"
	"time"
)

// Gateway TODO
type Gateway interface {
	ReadConfig() error
	ReadTime() error
	SetTime(time.Time) error
	Acquire(chan<- Event)
}

type GatewayType int

const (
	SerialGateway GatewayType = iota
	NATSGateway
)

func (gatewayType *GatewayType) MarshalText() ([]byte, error) {
	switch *gatewayType {
	case SerialGateway:
		return []byte("serial"), nil
	case NATSGateway:
		return []byte("nats"), nil
	default:
		return nil, fmt.Errorf("unknown gateway type: %d", gatewayType)
	}
}

func (gatewayType *GatewayType) UnmarshalText(text []byte) error {
	var err error = nil

	switch string(text) {
	case "serial":
		*gatewayType = SerialGateway
	case "nats":
		*gatewayType = NATSGateway
	default:
		err = fmt.Errorf("unknown gateway type: %s", text)
	}

	return err
}

func (gatewayType GatewayType) String() string {
	text, err := (&gatewayType).MarshalText()
	if err != nil {
		return fmt.Sprintf("unknown gateway type: %d", gatewayType)
	}

	return string(text)
}
