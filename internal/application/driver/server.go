package driver

import "fmt"

// Server TODO
type Server interface {
	Serve() error
}

type ServerType int

const (
	NATSServer ServerType = iota
)

func (serverType *ServerType) MarshalText() ([]byte, error) {
	switch *serverType {
	case NATSServer:
		return []byte("nats"), nil
	default:
		return nil, fmt.Errorf("unknown server type: %d", serverType)
	}
}

func (serverType *ServerType) UnmarshalText(text []byte) error {
	var err error = nil

	switch string(text) {
	case "nats":
		*serverType = NATSServer
	default:
		err = fmt.Errorf("unknown server type: %s", text)
	}

	return err
}

func (serverType ServerType) String() string {
	text, err := (&serverType).MarshalText()
	if err != nil {
		return fmt.Sprintf("unknown server type: %d", serverType)
	}

	return string(text)
}
