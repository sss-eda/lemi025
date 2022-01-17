package driving

import "github.com/sss-eda/lemi025/pkg/eventsourcing"

type Drive struct {
	OnConfigRead func(ConfigReadEvent) error
	OnTimeRead   func(TimeReadEvent) error
	OnTimeSet    func(TimeSetEvent) error
}

type CommandType int

const (
	ReadConfig CommandType = iota
	ReadTime
	SetTime
)

func (ct CommandType) String() string {
	var s string

	switch ct {
	case ReadConfig:
		s = "ReadConfig"
	case ReadTime:
		s = "ReadTime"
	case SetTime:
		s = "SetTime"
	default:
		s = ""
	}

	return s
}

func (this CommandType) Equals(that eventsourcing.CommandType) bool {
	if this != that {
		return false
	}

	return true
}

type ReadConfigCommand struct{}

// Type TODO
func (command ReadConfigCommand) Type() eventsourcing.CommandType {
	return ReadConfig
}

// Payload TODO
func (command ReadConfigCommand) Payload() []byte {
	return []byte{}
}

type ReadTimeCommand struct{}

type SetTimeCommand struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}
