package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/listing/getinstrumentbyid"
)

func main() {
	nc, err := nats.Connect("nats://sansa.dev:4222")
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

	postgres.GetInstrumentByID(db) // Returns what? func(InstrumentID) (Instrument, error)
	postgres.Project(db)           // Returns what? func(event) error

	// Projector
	sub, err := js.Subscribe(
		"hermanus.lemi025.*.events.>",
		func(msg *nats.Msg) {
			var event listing.Event
			err := json.Unmarshal(msg.Data, &event)
			if err != nil {
				log.Println(err)
				return
			}
			event.Project(postgres.OnEvent(db))
			// postgres.Project(db)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	http.HandleFunc( // Get method
		"/lemi025/instrument/1",
		rest.GetInstrumentByID( // returns http.HandlerFunc
			getinstrumentbyid.UseCase( // returns listing.HandleFunc
				postgres.GetInstrumentByID(db), // returns func(InstrumentID) Instrument
			),
		),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
