package nats

import (
	"context"
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// OnEvent TODO
func OnEvent[Event lemi025.EventPayloads](
	nc *natsio.Conn,
	deserialize func([]byte) (*Event, error),
) func(context.Context, string, func(*Event)) error {
	return func(
		ctx context.Context,
		subject string,
		handle func(*Event),
	) error {
		sub, err := nc.Subscribe(subject, func(msg *natsio.Msg) {
			event, err := deserialize(msg.Data)
			if err != nil {
				log.Println(err)
			}
			handle(event)
		})
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		<-ctx.Done()

		return ctx.Err()
	}
}
