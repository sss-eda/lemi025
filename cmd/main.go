package main

import (
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readtime"
	"github.com/sss-eda/lemi025/pkg/jetstream"
	"github.com/sss-eda/lemi025/pkg/nats"
)

// The listing and controlling will be in two separate
func main() {
	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	sub, err := nc.Subscribe(
		"dev.lemi025.*.mutations.readconfig",
		nats.HandleMutation( // returns func(*natsio.Msg)
			readtime.Mutate( // returns func(lemi025.ResponseWriter, lemi025.ReadConfigRequest)
				jetstream.Load(js), // returns func() error
				jetstream.Save(js),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
}
