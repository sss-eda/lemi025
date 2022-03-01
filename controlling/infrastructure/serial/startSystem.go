package serial

import (
	"io"
)

// StartSystem TODO
func StartSystem(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x39 // Token "9" indicating command type: start system

		_, err := writer.Write(buffer)
		return err

	}
}
