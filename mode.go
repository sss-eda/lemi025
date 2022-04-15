package lemi025

import "fmt"

const (
	// FLASHMode TODO
	FLASHMode byte = 0x01
	// PCMode TODO
	PCMode byte = 0x02
	// FLASHandPCMode TODO
	FLASHandPCMode byte = 0x03
)

// Mode TODO
type Mode struct {
	mode byte
}

// NewMode TODO
func NewMode(
	mode byte,
) (*Mode, error) {
	if mode != FLASHMode && mode != PCMode && mode != FLASHandPCMode {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}

	return &Mode{mode}, nil
}

// MarshalText TODO
func (mode Mode) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("Mode: %d", mode.mode)), nil
}

func (mode Mode) UnmarshalText(text []byte) error {
	return nil
}

// String TODO
func (mode Mode) String() string {
	switch mode.mode {
	case FLASHMode:
		return "FLASH"
	case PCMode:
		return "PC"
	case FLASHandPCMode:
		return "FLASH and PC"
	default:
		return "unknown"
	}
}
