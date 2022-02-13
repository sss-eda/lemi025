package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// Projector TODO
func Projector[Event any](
	project func(Event) error,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var event Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Println(err)
		}
		err = project(event)
		if err != nil {
			log.Println(err)
		}
	}
}
