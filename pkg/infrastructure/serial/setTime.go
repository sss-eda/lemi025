package serial

import (
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/internal/domain"
)

// SetTime TODO
func SetTime(
	w io.Writer,
) func(domain.SetTimeCommand) error {
	return func(command domain.SetTimeCommand) error {
		year, err := bcd.Encode(uint8(command.Year - 2000))
		if err != nil {
			return err
		}
		month, err := bcd.Encode(command.Month)
		if err != nil {
			return err
		}
		day, err := bcd.Encode(command.Day)
		if err != nil {
			return err
		}
		hour, err := bcd.Encode(command.Hour)
		if err != nil {
			return err
		}
		minute, err := bcd.Encode(command.Minute)
		if err != nil {
			return err
		}
		second, err := bcd.Encode(command.Second)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte{0x3D, 0x32, year, month, day, hour, minute, second})
		if err != nil {
			return err
		}

		return nil
	}
}
