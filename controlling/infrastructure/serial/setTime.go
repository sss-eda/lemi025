package serial

import (
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/instrument"
)

// SetTime TODO
func SetTime(
	writer io.Writer,
) func(instrument.Time) error {
	return func(time instrument.Time) error {
		var err error
		buffer := make([]byte, 8)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x32 // Token "2" indicating command type: set time

		// Encode and append year
		// Assumption: Year is in the 21st century
		buffer[2], err = bcd.Encode(uint8(time.Year - 2000))
		if err != nil {
			return err
		}

		// Encode and append month
		buffer[3], err = bcd.Encode(time.Month)
		if err != nil {
			return err
		}

		// Encode and append day
		buffer[4], err = bcd.Encode(time.Day)
		if err != nil {
			return err
		}

		// Encode and append hour
		buffer[5], err = bcd.Encode(time.Hour)
		if err != nil {
			return err
		}

		// Encode and append minute
		buffer[6], err = bcd.Encode(time.Minute)
		if err != nil {
			return err
		}

		// Encode and append second
		buffer[7], err = bcd.Encode(time.Second)
		if err != nil {
			return err
		}

		_, err = writer.Write(buffer)
		return err
	}
}
