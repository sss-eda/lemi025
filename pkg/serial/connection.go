package serial

import "io"

type Connection struct {
	Client
}

// Connect TODO
func Connect(
	r io.ReadWriteCloser,
) (*Connection, error) {
	connection := Connection{
		rwc,
	}

	return &connection, nil
}
