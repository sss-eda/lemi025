package lemi025

type Mode1 byte

const (
	MENU   Mode1 = 0x00
	RECORD       = 0x01
)

func (mode1 Mode1) String() string {
	switch mode1 {
	case MENU:
		return "MENU"
	case RECORD:
		return "RECORD"
	}

	return "UNDEFINED"
}
