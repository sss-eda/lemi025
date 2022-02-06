package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// SetCoefficients1Request TODO
type SetCoefficients1Request struct {
	Mode lemi025.Mode
}

// SetCoefficients1Response TODO
type SetCoefficients1Response struct{}

// SetCoefficients1CommandMsgHandler TODO - Here we actually shouldn't be sending the strategy itself. It should be an application level wrapper around the strategy? Because the application is where the validation will happen. Now it can be bypassed and serial can be directly injected into nats.
func SetCoefficients1CommandMsgHandler(
	strategy lemi025.SetCoefficients1Strategy,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		request := ReadTimeRequest{}
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Println(err)
			return
		}
		err = strategy(lemi025.SetCoefficients1Command{})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
