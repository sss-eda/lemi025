package serial

import (
	"io"
)

// ReadTime TODO
func ReadTime(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x31 // Token "1" indicating command type: read time

		_, err := writer.Write(buffer)
		return err
	}
}
