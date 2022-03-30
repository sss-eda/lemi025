package serial

import (
	"fmt"

	"github.com/sss-eda/encoding/bcd"
)

type configReadEvent struct {
	stationType   string
	stationNumber uint8
}

func (event *configReadEvent) UnmarshalBinary(data []byte) error {
	if len(data) != 5 {
		return fmt.Errorf("unexpected data length: %v", len(data))
	}

	event.stationType = string(data[:4])
	if event.stationType != "025 " {
		event = nil
		return fmt.Errorf("unexpected station type: %v", event.stationType)
	}

	event.stationNumber, err = bcd.Decode(data[4])
	if err != nil {
		event = nil
		return fmt.Errorf("failed to decode station number: %w", err)
	}

	return nil
}
