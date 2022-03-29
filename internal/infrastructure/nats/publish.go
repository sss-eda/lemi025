package nats

import (
	"context"
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// Publish TODO
func Publish[Payload lemi025.CommandPayloads | lemi025.EventPayloads](
	nc *natsio.Conn,
	subject string,
	serialize func(*Payload) ([]byte, error),
) func(context.Context, *Payload) {
	return func(
		ctx context.Context,
		payload *Payload,
	) {
		data, err := serialize(payload)
		if err != nil {
			log.Println(err)
			return
		}

		err = nc.Publish(subject, data)
		if err != nil {
			log.Println(err)
		}
	}
}
