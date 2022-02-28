package serial

import "io"

// ReadGPSData TODO
func ReadGPSData(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x37 // Token "7" indicating command type: read GPS data
		buffer[2] = 0xFF // XX - Don't care
		buffer[3] = 0xFF // XX - Don't care

		_, err := writer.Write(buffer)
		return err
	}
}
