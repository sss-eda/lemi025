package serial

import (
	"fmt"

	"github.com/sss-eda/encoding/bcd"
	"github.com/sss-eda/lemi025"
)

// Config TODO
type Config struct {
	StationType   string
	StationNumber uint8
}

// UnmarshalBinary TODO
func (config *Config) UnmarshalBinary(data []byte) error {
	config = &Config{}
	var err error

	if len(data) != 5 {
		return fmt.Errorf("unexpected data length: %v", len(data))
	}

	config.StationType = string(data[:4])
	if config.StationType != "025 " {
		config = nil
		return fmt.Errorf("unexpected station type: %v", config.StationType)
	}

	config.StationNumber, err = bcd.Decode(data[4])
	if err != nil {
		config = nil
		return fmt.Errorf("failed to decode station number: %w", err)
	}

	return nil
}

type readConfigCommand lemi025.ReadConfigCommandPayload

// MarshalBinary TODO
func (command readConfigCommand) MarshalBinary() ([]byte, error) {
	return []byte{sendToken, readConfigToken}, nil
}
