package commanding

import "io"

type Command interface {
	Execute() error
}

type ReadConfigCommand struct {
	writer io.Writer
}

func (command ReadConfigCommand) Execute() (err error) {
	_, err = command.writer.Write([]byte{0x3D, 0x30})

	return
}

type Connection struct {
	io.ReadWriteCloser
}

func (connection *Connection) () Command {
	return ReadConfigCommand{
		writer: connection,
	}
}

func (connection *Connection) ReadConfig() Command {
	return ReadConfigCommand{
		writer: connection,
	}
}
