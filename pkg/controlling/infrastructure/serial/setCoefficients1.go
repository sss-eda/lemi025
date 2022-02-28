package serial

import (
	"fmt"
	"io"

	"github.com/sss-eda/lemi025/pkg/instrument"
)

// SetCoefficients1 TODO
func SetCoefficients1(
	writer io.Writer,
) func(instrument.Coefficients1) error {
	return func(coefficients1 instrument.Coefficients1) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x33 // Token "3" indicating command type: set coefficients1
		buffer[2] = 0xFF // XX - Don't care

		switch coefficients1.Mode {
		case instrument.FLASH:
			buffer[3] = 0x01
		case instrument.PC:
			buffer[3] = 0x02
		case instrument.FLASHandPC:
			buffer[3] = 0x03
		default:
			return fmt.Errorf("invalid mode: %v", coefficients1.Mode)
		}

		_, err := writer.Write(buffer)
		return err

	}
}
