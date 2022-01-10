package gpio

import (
	"io"

	"github.com/sss-eda/lemi025"
)

// ReadConfigCommand TODO
type ReadConfigCommand struct {
	io.Writer
	*lemi025.Config
}

// Execute TODO
func (command *ReadConfigCommand) Execute() error {
	_, err := command.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}
