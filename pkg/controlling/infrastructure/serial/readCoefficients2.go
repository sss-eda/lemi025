package serial

import (
	"io"
)

// ReadCoefficients2 TODO
func ReadCoefficients2(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x36 // Token "6" indicating command type: read coefficients2

		_, err := writer.Write(buffer)
		return err
	}
}
