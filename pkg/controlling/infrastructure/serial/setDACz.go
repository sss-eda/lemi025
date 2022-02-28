package serial

import (
	"encoding/binary"
	"io"
)

// SetDACz TODO
func SetDACz(
	writer io.Writer,
) func(int16) error {
	return func(z int16) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x3F // Token "?" indicating command type: set DAC-Z

		// Encode and append x
		binary.LittleEndian.PutUint16(buffer[2:], uint16(z))

		_, err := writer.Write(buffer)
		return err
	}
}
