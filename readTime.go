package lemi025

// TimeReader TODO
type TimeReader interface {
	ReadTime() error
}

// ReadTimeCommand TODO
type ReadTimeCommand struct {
	TimeReader `json:"-"`
}

// Execute TODO
func (command *ReadTimeCommand) Execute() error {
	return command.ReadTime()
}
