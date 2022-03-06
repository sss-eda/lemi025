package serial

import (
	"bytes"
	"fmt"
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/controlling"
)

// ReadTime TODO
func ReadTime(
	writer io.Writer,
) func() error {
	return func() error {
		buffer := make([]byte, 2)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x31 // Token "1" indicating command type: read time

		_, err := writer.Write(buffer)
		return err
	}
}

// OnTimeRead TODO
func OnTimeRead(
	handle func(*controlling.TimeReadEvent) error,
) func([]byte) error {
	return func(data []byte) error {
		if len(data) != 8 {
			return fmt.Errorf("invalid data length (expected slice of length 8, got: %v)", len(data))
		}
		if !bytes.Equal(data[:2], []byte{0x3F, 0x31}) {
			return fmt.Errorf("invalid data (expected slice to start with 0x3F31, got: %v)", data[:2])
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
		return handle(&controlling.TimeReadEvent{
			Year:   uint16(year) + 2000,
			Month:  month,
			Day:    day,
			Hour:   hour,
			Minute: minute,
			Second: second,
		})
	}
}
