package serial

import (
	"fmt"

	"github.com/sss-eda/lemi025"
)

// ReadConfigCommand TODO
type ReadConfigCommand struct {
	payload *lemi025.ReadConfigCommandPayload
}

// Payload TODO
func (command *ReadConfigCommand) Payload() (payload *lemi025.ReadConfigCommandPayload, err error) {
	payload = command.payload
	if payload == nil {
		err = fmt.Errorf("invalid command: payload not set")
	}
	return
}
