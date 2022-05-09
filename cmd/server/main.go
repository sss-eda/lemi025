package main

import (
	"log"
	"net/http"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
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

	// Listen at endpoint for new request. Publish to NATS.
	http.Subscribe("lemi025/1/commands/readConfig", jetstream.Publish(js, lemi025.ReadConfig()))

	log.Fatal(http.ListenAndServe(":8080", handler))
}
