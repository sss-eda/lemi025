package lemi025

import "fmt"

// Mode TODO
type Mode uint8

const (
	// FLASHMode TODO
	FLASHMode Mode = iota
	// PCMode TODO
	PCMode
	// FLASHandPCMode TODO
	FLASHandPCMode
)

// String TODO
func (mode Mode) String() string {
	switch mode {
	case FLASHMode:
		return "FLASH"
	case PCMode:
		return "PC"
	case FLASHandPCMode:
		return "FLASH and PC"
	default:
		return "Invalid Mode"
	}
}

// // NewMode TODO
// func NewMode(
// 	mode byte,
// ) (*Mode, error) {
// 	if mode != FLASHMode && mode != PCMode && mode != FLASHandPCMode {
// 		return nil, fmt.Errorf("invalid mode: %d", mode)
// 	}

// 	return &Mode{mode}, nil
// }

// MarshalBinary TODO
func (mode Mode) MarshalBinary() ([]byte, error) {
	return []byte{byte(mode)}, nil
}

// UnmarshalBinary TODO
func (mode *Mode) UnmarshalBinary(data []byte) error {
	if len(data) != 1 {
		return fmt.Errorf("invalid data length: Expected 1; got %d", len(data))
	}

	m := Mode(data[0])

	if m != FLASHMode && m != PCMode && m != FLASHandPCMode {
		return fmt.Errorf("invalid mode: %d", m)
	}

	*mode = m

	return nil
}

// MarshalText TODO
func (mode Mode) MarshalText() ([]byte, error) {
	return []byte(mode.String()), nil
}

// UnmarshalText TODOs
func (mode Mode) UnmarshalText(text []byte) error {
	switch string(text) {
	case "FLASH":
		mode = FLASHMode
	case "PC":
		mode = PCMode
	case "FLASH and PC":
		mode = FLASHandPCMode
	default:
		return fmt.Errorf("invalid mode: Expected \"FLASH\", \"PC\" or \"FLASH and PC\"; got %s", text)
	}

	return nil
}
