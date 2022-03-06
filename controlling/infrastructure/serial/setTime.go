package serial

import (
	"bytes"
	"fmt"
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/controlling"
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

// OnTimeSet TODO
func OnTimeSet(
	handle func(*controlling.TimeSetEvent) error,
) func([]byte) error {
	return func(data []byte) error {
		if len(data) != 8 {
			return fmt.Errorf("invalid data length (expected slice of length 8, got: %v)", len(data))
		}
		if !bytes.Equal(data[:2], []byte{0x3F, 0x32}) {
			return fmt.Errorf("invalid data (expected slice to start with 0x3F32, got: %v)", data[:2])
		}
		year, err := bcd.Decode(data[2])
		if err != nil {
			return fmt.Errorf("invalid year (bcd.Decode() returned with error: %v)", err)
		}
		month, err := bcd.Decode(data[3])
		if err != nil {
			return fmt.Errorf("invalid month (bcd.Decode() returned with error: %v)", err)
		}
		day, err := bcd.Decode(data[4])
		if err != nil {
			return fmt.Errorf("invalid day (bcd.Decode() returned with error: %v)", err)
		}
		hour, err := bcd.Decode(data[5])
		if err != nil {
			return fmt.Errorf("invalid hour (bcd.Decode() returned with error: %v)", err)
		}
		minute, err := bcd.Decode(data[6])
		if err != nil {
			return fmt.Errorf("invalid minute (bcd.Decode() returned with error: %v)", err)
		}
		second, err := bcd.Decode(data[7])
		if err != nil {
			return fmt.Errorf("invalid second (bcd.Decode() returned with error: %v)", err)
		}
		return handle(&controlling.TimeSetEvent{
			Year:   uint16(year) + 2000,
			Month:  month,
			Day:    day,
			Hour:   hour,
			Minute: minute,
			Second: second,
		})
	}
}
