package lemi025

type Mode byte

const (
	FLASH      Mode = 0x01
	PC              = 0x02
	FLASHandPC      = 0x03
)

func (mode Mode) String() string {
	switch mode {
	case FLASH:
		return "FLASH"
	case PC:
		return "PC"
	case FLASHandPC:
		return "FLASH AND PC"
	}

	return "UNDEFINED"
}
