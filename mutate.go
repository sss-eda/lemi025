package lemi025

// Command TODO
func Command[C Commands](
	command C,
) error {
	return nil
}

// Commands TODO
type Commands interface {
	ReadConfigCommand | ReadTimeCommand | SetTimeCommand
}

type ReadConfigCommand struct{}

type ReadTimeCommand struct{}

type SetTimeCommand struct{}
