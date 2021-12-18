package lemi025

// Coefficients1 Value Object
type Coefficients1 struct {
	mode  Mode
	uin   Uin
	mode1 Mode1
}

type Mode uint8

const (
	FLASH      Mode = 1
	PC              = 2
	FLASHandPC      = 3
)

func (mode Mode) String() string {
	switch mode {
	case FLASH:
		return "FLASH"
	case PC:
		return "PC"
	case FLASHandPC:
		return "FLASH and PC"
	}

	return "Undefined"
}

type Uin uint8

func (uin Uin) String() string {
	return string(uin / 10)
}

type Mode1 uint8

const (
	MENU   Mode1 = 0
	RECORD       = 1
)

func (mode1 Mode1) String() string {
	switch mode1 {
	case MENU:
		return "MENU"
	case RECORD:
		return "RECORD"
	}

	return "Undefined"
}
