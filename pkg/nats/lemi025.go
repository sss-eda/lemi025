package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/sss-eda/lemi025"
)

// Lemi025 TODO
type Lemi025 struct {
	nc *nats.Conn // Just a plain subscription to get commands.
}

func (l *Lemi025) readConfig(msg *nats.Msg) {
	command := lemi025.ReadConfigCommand{}

	err := json.Unmarshal(msg.Data, &command)
	if err != nil {
		log.Println(err)
		return
	}

	command.Execute()
}
