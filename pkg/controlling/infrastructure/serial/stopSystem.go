package serial

import (
	"io"
)

// StopSystem TODO
func StopSystem(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x38 // Token "8" indicating command type: stop system

		_, err := writer.Write(buffer)
		return err
	}
}
