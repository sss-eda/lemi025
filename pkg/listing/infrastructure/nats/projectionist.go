package nats

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

// Projector - Projects any event type using a projector func
// This is where the decoder/unmarshaller can be as well?
// This should move out to the universal infrastructure layer.
func Projector[Event any](
	project func(context.Context, Event) error,
	decode func([]byte, *Event) error,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var event Event
		err := decode(msg.Data, &event)
		if err != nil {
			log.Println(err)
		}
		err = project(context.Background(), event)
		if err != nil {
			log.Println(err)
		}
	}
}
