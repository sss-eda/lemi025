package gpio

import (
	"io"
	"sync"
)

// Connection TODO
type Connection struct {
	*sync.RWMutex
	io.ReadWriteCloser
}

// Connector TODO
// type Connector func(...interface{}) (io.ReadWriteCloser, error)

// Connect TODO
func Connect(
	rwc io.ReadWriteCloser,
) (*Connection, error) {
	// rwc, err := connector(args)
	// if err != nil {
	// 	return nil, err
	// }

	connection := &Connection{
		&sync.RWMutex{},
		rwc,
	}

	// We should start reading here already

	return connection, nil
}
