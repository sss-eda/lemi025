package eventsourcing

type CommandType interface {
	String() string
	Equals(CommandType) bool
}

type Command interface {
	Type() CommandType
	Payload() []byte
}
