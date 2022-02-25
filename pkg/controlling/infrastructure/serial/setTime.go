package serial

import (
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// SetTime TODO
func SetTime(
	writer io.Writer,
	time lemi025.Time,
) error {
	year, err := bcd.Encode(uint8(time.Year - 2000))
	if err != nil {
		return err
	}
	month, err := bcd.Encode(time.Month)
	if err != nil {
		return err
	}
	day, err := bcd.Encode(time.Day)
	if err != nil {
		return err
	}
	hour, err := bcd.Encode(time.Hour)
	if err != nil {
		return err
	}
	minute, err := bcd.Encode(time.Minute)
	if err != nil {
		return err
	}
	second, err := bcd.Encode(time.Second)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte{0x3D, 0x32, year, month, day, hour, minute, second})
	return err
}

// SetTimeFunc TODO
func SetTimeFunc(
	writer io.Writer,
) func(instrument.Time) error {
	return func() error {
		return SetTime(writer, year, month, day, hour, minute, second)
	}
}
