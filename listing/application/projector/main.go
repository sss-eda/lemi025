package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/listing"
	"github.com/sss-eda/lemi025/listing/infrastructure/nats"
	"github.com/sss-eda/lemi025/listing/infrastructure/postgres"
)

const (
	envStorageType string = "STORAGE_TO_USE"
	envServerType  string = "SERVER_TO_USE"
)

func main() {
	storageType, ok := os.LookupEnv(envStorageType)
	if !ok {
		storageType = "postgres"
	}
	storageType := os.Getenv("STORAGE_TO_USE")
	serverType := os.Getenv("SERVER_TO_USE")

	var storage listing.Repository
	var server listing.Server

	switch storageType {
	case "", "postgres":
		postgresHost := os.Getenv("POSTGRES_HOST")
		postgresPort := os.Getenv("POSTGRES_PORT")
		postgresUser := os.Getenv("POSTGRES_USER")
		postgresPassword := os.Getenv("POSTGRES_PASSWORD")
		postgresDatabase := os.Getenv("POSTGRES_DATABASE")

		connString := "postgres://" + postgresUser + ":" + postgresPassword +
			"@" + postgresHost + ":" + postgresPort + "/" + postgresDatabase

		db, err := pgx.Connect(context.Background(), connString)
		if err != nil {
			log.Fatal(err)
		}

		storage := postgres.NewListingRepository(db)
	case "mongodb":
	default:
		log.Fatal
	}

	nc, err := natsio.Connect("nats://sansa.dev:4222, nats://sansa.dev:4223, nats://sansa.dev:4224")
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
