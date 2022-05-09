package nats

import (
	"context"
	"encoding/json"
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// This is the code that can be called by the client that should be available
// as a library later in root of repo.

// Subscribe TODO
func Subscribe[Payload lemi025.CommandPayloads](
	ctx context.Context,
	nc *natsio.Conn,
	subject string,
	handle func(context.Context, *Payload) error,
) error {
	sub, err := nc.Subscribe(subject, func(msg *natsio.Msg) {
		var payload Payload
		err := json.Unmarshal(msg.Data, &payload)
		if err != nil {
			log.Println(err)
			return
		}
		err = handle(ctx, &payload)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	<-ctx.Done()

	return ctx.Err()
}

// // Send TODO
// func Send[Payload lemi025.CommandPayloads | lemi025.EventPayloads](
// 	nc *natsio.Conn,
// 	subject string,
// 	serialize func(*Payload) ([]byte, error),
// ) func(context.Context, *Payload) error {
// 	return func(
// 		ctx context.Context,
// 		payload *Payload,
// 	) error {
// 		data, err := serialize(payload)
// 		if err != nil {
// 			return err
// 		}

// 		return nc.Publish(subject, data)
// 	}
// }
