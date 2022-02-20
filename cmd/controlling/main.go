package main

import (
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
	tarm "github.com/tarm/serial"
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

	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	serial.Listen(port)(
		nats.controller,
		&lemi025.AcquireRequest{},
	) // returns func(responsewriter, request)

}
