package serial

import (
	"io"
	"log"

	"github.com/sss-eda/lemi025"
)

type Connection struct {
	station *lemi025.Station
	buffer  []byte
	writer  io.Writer
}

func Connect(port io.ReadWriteCloser) (*Connection, error) {
	connection := Connection{
		station: &lemi025.Station{},
		buffer:  []byte{},
		writer:  port,
	}

	go connection.read(port)

	return &connection, nil
}

func (connection *Connection) read(reader io.Reader) {
	buffer := make([]byte, 128)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			log.Println(err)
			continue
		}
		for i := 0; i < n; i++ {

		}
	}
}

func (connection *Connection) ReadConfig() error {
	_, err := connection.writer.Write([]byte{0x3D, 0x30})
	if err != nil {
		return err
	}

	return nil
}

func (connection *Connection) ReadTime() error {
	_, err := connection.writer.Write([]byte{0x3D, 0x31})
	if err != nil {
		return err
	}

	return nil
}
