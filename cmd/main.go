package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
