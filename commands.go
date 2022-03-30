package lemi025

import "fmt"

// CommandPayloads TODO
type CommandPayloads interface {
	ReadConfigCommandPayload | ReadTimeCommandPayload | SetTimeCommandPayload
}

// ReadConfigCommandPayload TODO
type ReadConfigCommandPayload struct{}

// ReadTimeCommandPayload TODO
type ReadTimeCommandPayload struct{}

// SetTimeCommandPayload TODO
type SetTimeCommandPayload struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// SetCoefficients1CommandPayload TODO
type SetCoefficients1CommandPayload struct {
	Mode  Mode
	Uin   Uin
	Mode1 Mode1
}

// Mode TODO
type Mode byte

const (
	// FLASHMode TODO
	FLASHMode Mode = 0x01
	// PCMode TODO
	PCMode Mode = 0x02
	// FLASHandPCMode TODO
	FLASHandPCMode Mode = 0x03
)

// Mode1 TODO
type Mode1 byte

const (
	// MenuMode1 TODO
	MenuMode1 Mode1 = 0x00
	// RecordMode1 TODO
	RecordMode1 Mode1 = 0x01
)

// Uin TODO
type Uin uint8

// String TODO
func (uin Uin) String() string {
	return fmt.Sprintf("%d", uin/10)
}
