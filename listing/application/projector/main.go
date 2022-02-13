package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/listing/infrastructure/nats"
	"github.com/sss-eda/lemi025/listing/infrastructure/postgres"
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
	// layer rather than here.
	sub, err := js.Subscribe(
		"hermanus.lemi025.*.config.read", // This is part of the schema and
		nats.Projector(postgres.ReadConfig(db)),
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
