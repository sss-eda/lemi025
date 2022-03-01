package serial

import (
	"encoding/binary"
	"io"
)

// SetDACx TODO
func SetDACx(
	writer io.Writer,
) func(int16) error {
	return func(x int16) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x3D // Token "=" indicating command type: set DAC-X

		// Encode and append x
		binary.LittleEndian.PutUint16(buffer[2:], uint16(x))

		_, err := writer.Write(buffer)
		return err
	}
}
