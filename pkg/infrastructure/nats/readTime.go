package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// ReadTimeRequest TODO
type ReadTimeRequest struct{}

// ReadTimeResponse TODO
type ReadTimeResponse struct{}

// ReadTimeAdapter TODO - Here we actually shouldn't be sending the strategy itself. It should be an application level wrapper around the strategy? Because the application is where the validation will happen. Now it can be bypassed and serial can be directly injected into nats.
func ReadTimeAdapter(strategy lemi025.ReadTimeStrategy) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		request := ReadTimeRequest{}
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Println(err)
			return
		}
		err = strategy(lemi025.ReadTimeInput{})
		if err != nil {
			log.Println(err)
			return
		}
	}
}
