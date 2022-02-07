package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// ReadConfigRequest TODO
type ReadConfigRequest struct{}

// ReadConfigResponse TODO
type ReadConfigResponse struct {
	StationNumber uint8 `json:"stationNumber"`
}

// ReadConfigCommandHandler TODO
func ReadConfigCommandMsgHandler(
	strategy lemi025.ReadConfigStrategy,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		request := ReadConfigRequest{}
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Println(err)
			return
		}
		err = strategy(lemi025.ReadConfigCommand{})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
