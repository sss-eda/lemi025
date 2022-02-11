package main

import (
	"log"
	"net/http"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/usecases/getconfig"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
	"github.com/sss-eda/lemi025/internal/infrastructure/nats"
	"github.com/sss-eda/lemi025/internal/infrastructure/serial"

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

	// This is very nice, everything is so nicely decoupled and everything.
	// It might be nice to package together a bunch of them into applications.
	sub, err := nc.Subscribe(
		"lemi025.1.queries.readconfig",
		nats.Handle(
			readconfig.UseCase(
				serial.ConfigReader(port), // returns func() error
			), // returns func(request) response
			newencoding.Serializer(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	controller := serial.Run(port)
	controller(eventstore)

	http.HandleFunc(
		"/config",
		rest.Handle(
			getconfig.UseCase(
				postgres.ConfigGetter(psqlConn),
			),
		),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
