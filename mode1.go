package lemi025

import "fmt"

const (
	// MenuMode1 TODO
	MenuMode1 byte = 0x00
	// RecordMode1 TODO
	RecordMode1 byte = 0x01
)

// Mode1 TODO
type Mode1 struct {
	value byte
}

// MarshalText TODO
func (mode1 Mode1) MarshalText() ([]byte, error) {
	switch mode1.value {
	case MenuMode1:
		return []byte("0"), nil
	case RecordMode1:
		return []byte("1"), nil
	default:
		// This should never happen. Mode1 should be always valid due to proper
		// encapsulation.
		return nil, fmt.Errorf("invalid value: %d", mode1.value)
	}
}

// UnmarshalText TODO
func (mode1 *Mode1) UnmarshalText(text []byte) error {
	if len(text) != 1 {
		return fmt.Errorf("invalid length: Expected 1, got %d", len(text))
	}

	switch string(text[0]) {
	case "0":
		mode1.value = MenuMode1
	case "1":
		mode1.value = RecordMode1
	default:
		return fmt.Errorf("invalid value: Expected \"0\" or \"1\", got %s", string(text[0]))
	}

	return nil
}

// String TODO
func (mode1 Mode1) String() string {
	switch mode1.value {
	case MenuMode1:
		return "Menu"
	case RecordMode1:
		return "Record"
	default:
		return "unknown"
	}
}
