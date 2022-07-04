package jetstream

import (
	"encoding/json"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

type CommandHandler struct{}

type Command struct{}

func NewCommandHandler(
	station lemi025.Station,
) natsio.MsgHandler {
	return func(msg *natsio.Msg) {
		command := lemi025.Command{}
		json.Unmarshal(msg.Data, &command)

		switch command.Kind {
		case lemi025.ReadConfigCommandKind:
			instrument.ReadConfig(command.Payload)
		case lemi025.ReadTimeCommandKind:
			instrument.ReadTime(command.Payload)
		case lemi025.SetTimeCommandKind:
			instrument.SetTime(command.Payload)
		}
	}
}
