package serial

import (
	"encoding/binary"
	"io"
)

// SetDACy TODO
func SetDACy(
	writer io.Writer,
) func(int16) error {
	return func(y int16) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x3E // Token ">" indicating command type: set DAC-Y

		// Encode and append x
		binary.LittleEndian.PutUint16(buffer[2:], uint16(y))

		_, err := writer.Write(buffer)
		return err
	}
}
