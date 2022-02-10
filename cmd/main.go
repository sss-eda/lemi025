package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/infrastructure/serial"

	tarm "github.com/tarm/serial"
)

func main() {
	nc, err := nats.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	readConfig.Handle(
		nats.ReadConfigHandler(nc),     // Where it gets the requests from
		serial.ReadConfigGateway(port), //
	)

	service := instrument.NewService()

	sub, err := nc.Subscribe(
		"hermanus.lemi025.1.readconfig",
		nats.ReadConfigMsgHandler(
			service.ReadConfig(
				serial.ReadConfig(port),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
}
