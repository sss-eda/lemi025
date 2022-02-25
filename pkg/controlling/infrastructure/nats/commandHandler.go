package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readtime"
)

type Command interface {
	readconfig.Command | readtime.Command
}

// HandleMutation TODO
func HandleMutation[CommandType Command](
	handle func(CommandType) error,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var command CommandType
		err := json.Unmarshal(msg.Data, &command)
		if err != nil {
			log.Println(err)
			return
		}
		err = handle(command)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
