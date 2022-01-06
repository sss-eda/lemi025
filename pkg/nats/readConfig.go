package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

// ReadConfigCommand TODO
type ReadConfigCommand struct {
	subject string
	nc      *nats.Conn
}

// NewReadConfigCommand TODO
func NewReadConfigCommand(
	natsConn *nats.Conn,
) *ReadConfigCommand {
	return &ReadConfigCommand{
		nc: natsConn,
	}
}

// Execute TODO
func (command *ReadConfigCommand) Execute() error {
	data, err := json.Marshal(command)
	if err != nil {
		return err
	}

	err = command.nc.Publish(command.subject, data)
	if err != nil {
		return err
	}

	return nil
}
