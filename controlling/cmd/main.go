package main

import (
	"log"

	"github.com/gorilla/mux"
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/controlling"
	"github.com/sss-eda/lemi025/controlling/infrastructure/http/rest"
	"github.com/sss-eda/lemi025/controlling/infrastructure/jetstream"
	"github.com/sss-eda/lemi025/controlling/infrastructure/serial"
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

	// instrument := controlling.Instrument{}

	commandRouter := mux.NewRouter()

	commandRouter.HandleFunc(
		"/commands/read-config",
		rest.HandleCommand( // returns http.HandlerFunc
			controlling.ExecuteCommand( // This might be usefull in terms of directing
				serial.ReadConfig(port), // returns
			),
		),
	).Methods("POST")

	commandRouter.HandleFunc(
		"/commands/read-time",
		rest.CommandHandlerFunc( // returns http.HandlerFunc
			controlling.HandleCommands( // This might be usefull in terms of directing
				serial.ReadTime(port), // returns func(Command) error
			),
		),
	).Methods("POST")

	serial.Acquire(
		port,
		controlling.EventHandlerFunc(
			jetstream.EventDispatcher(js), // func(Event)
		),
	)
}
