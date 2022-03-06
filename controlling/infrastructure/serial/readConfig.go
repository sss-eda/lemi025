package serial

import (
	"bytes"
	"fmt"
	"io"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025/controlling"
)

// ConfigReader TODO
func ConfigReader(
	writer io.Writer,
) func(*controlling.ReadConfigCommand) error {
	return func(command *controlling.ReadConfigCommand) error {
		buffer := make([]byte, 4)

		buffer[0] = 0x3D // Token "=" indicating intention to send command
		buffer[1] = 0x30 // Token "0" indicating command type: read config
		buffer[2] = 0xFF // XX - Don't care
		buffer[3] = 0xFF // XX - Don't care

		_, err := writer.Write(buffer)
		return err
	}
}

// OnConfigRead TODO
func OnConfigRead(
	handle func(*controlling.ConfigReadEvent) error,
) func([]byte) error {
	return func(data []byte) error {
		if len(data) != 5 {
			return fmt.Errorf("invalid data length (expected slice of length 5, got: %v)", len(data))
		}
		if !bytes.Equal(data[:2], []byte{0x3F, 0x30}) {
			return fmt.Errorf("invalid data (expected slice to start with 0x3F30, got: %v)", data[:2])
		}
		stationType := string(data[:4])
		if stationType != "025 " {
			return fmt.Errorf("invalid station type (expected \"025 \", got: %v)", stationType)
		}
		stationNumber, err := bcd.Decode(data[4])
		if err != nil {
			return fmt.Errorf("invalid station number (bcd.Decode() returned with error: %v)", err)
		}
		return handle(&controlling.ConfigReadEvent{
			StationNumber: stationNumber,
		})
	}
}
