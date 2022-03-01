package serial

import "io"

// CheckFLASH TODO
func CheckFLASH(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x3A // Token ":" indicating command type: check FLASH

		_, err := writer.Write(buffer)
		return err
	}
}
