package serial

import (
	"io"
)

// ReadConfig TODO
func ReadConfig(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x30 // Token "0" indicating command type: read config
		buffer[2] = 0xFF // XX - Don't care
		buffer[3] = 0xFF // XX - Don't care

		_, err := writer.Write(buffer)
		return err
	}
}
