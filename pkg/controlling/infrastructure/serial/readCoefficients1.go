package serial

import (
	"io"
)

// ReadCoefficients1 TODO
func ReadCoefficients1(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x34 // Token "4" indicating command type: read coefficients1

		_, err := writer.Write(buffer)
		return err
	}
}
