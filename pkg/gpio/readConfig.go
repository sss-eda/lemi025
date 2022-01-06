package gpio

import (
	"io"

	"github.com/sss-eda/lemi025/pkg/driving"
)

// ReadConfigCommand TODO
type ReadConfigCommand struct {
	writer io.Writer
	config *driving.Config
}

// Execute TODO
func (command *ReadConfigCommand) Execute() error {
	_, err := command.writer.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
