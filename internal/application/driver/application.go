package driver

import (
	"fmt"

	"github.com/sss-eda/lemi025/internal/infrastructure/serial"
)

// Application TODO
type Application struct {
	Gateway GatewayType
	Server  ServerType
}

// Run TODO
func (application *Application) Run() error {
	var err error

	var gateway Gateway
	switch application.Gateway {
	case SerialGateway:
		gateway, err = serial.NewGateway()
	case NATSGateway:
		gateway, err = nats.NewGateway()
	default:
		err = fmt.Errorf("unknown gateway type: %d", application.Gateway)
	}
	if err != nil {
		return err
	}

	var server Server
	switch application.Server {
	case NATSServer:
		server, err = nats.NewServer()
	default:
		err = fmt.Errorf("unknown server type: %d", application.Server)
	}
	if err != nil {
		return err
	}

	return nil
}
