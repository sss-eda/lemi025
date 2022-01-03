package nats

import (
	"encoding/json"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

type Connection struct {
}

func Connect() *Connection {
	return &Connection{}
}

func (connection *Connection) ReadConfig(msg *natsio.Msg) {
	command := lemi025.ReadConfigCommand{}

	err := json.Unmarshal(msg.Data, &command)
	if err != nil {
		msg.Respond()
	}
}
