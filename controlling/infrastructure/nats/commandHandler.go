package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/controlling"
)

// CommandHandlerFunc TODO
func CommandHandlerFunc[Command controlling.Commands](
	handle func(*Command) error,
) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var command Command
		err := json.Unmarshal(msg.Data, &command)
		if err != nil {
			err2 := msg.Respond([]byte(err.Error()))
			if err2 != nil {
				log.Println(err, err2)
			}
		}
		err = handle(&command)
		if err != nil {
			err2 := msg.Respond([]byte(err.Error()))
			if err2 != nil {
				log.Println(err, err2)
			}
		}
	}
}
