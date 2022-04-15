package main

import (
	"log"
	"net/http"

	natsio "github.com/nats-io/nats.go"
)

func main() {
	nc, err := natsio.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	eventsource, err := jetstream.NewEventSource(js)
	if err != nil {
		log.Fatal(err)
	}

	handler, err := http.NewRESTHandler(eventsource)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
