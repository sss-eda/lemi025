package lemi025

const (
	// MenuMode1 TODO
	MenuMode1 byte = 0x00
	// RecordMode1 TODO
	RecordMode1 byte = 0x01
)

// Mode1 TODO
type Mode1 struct {
	mode1 byte
}

// MarshalText TODO
func (mode1 Mode1) MarshalText() ([]byte, error) {
	return []byte{}, nil
}

// UnmarshalText TODO
func (mode1 *Mode1) UnmarshalText(text []byte) error {
	return nil
}

// String TODO
func (mode1 Mode1) String() string {
	switch mode1.mode1 {
	case MenuMode1:
		return "Menu"
	case RecordMode1:
		return "Record"
	default:
		return "unknown"
	}
}
