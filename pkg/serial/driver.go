package serial

import (
	"io"

	"github.com/sss-eda/lemi025/pkg/driving"
)

type Driver struct {
	io.Reader
}

// Drive TODO
func (connection *Connection) Drive(
	input driving.Drive,
) error {
	return nil
}
