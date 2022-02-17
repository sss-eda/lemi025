package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
	"github.com/sss-eda/lemi025/pkg/listing/infrastructure/nats"
	"github.com/sss-eda/lemi025/pkg/listing/infrastructure/postgres"
)

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

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// The subject is part of the schema and really belongs in the infrastructure
	// layer rather than here. For now, I'll keep it like this. A container layer
	// can very easily be added later.
	// The projector does not need to go through a service/usecase layer.
	sub, err := js.Subscribe(
		"hermanus.lemi025.*.config.read",
		nats.Projector( // Returns func(*nats.Msg)
			postgres.OnConfigRead(db), // Returns func(instrument.ConfigReadEvent)
			func(data []byte, event *instrument.ConfigReadEvent) error {
				return json.Unmarshal(data, event)
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// jetstream.ProjectReadConfig(
	// 	js,
	// 	postgres.ReadConfig(db),
	// 	nats.Projector,
	// )
}
