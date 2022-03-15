package main

import (
	"context"
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/controlling"
	"github.com/sss-eda/lemi025/controlling/infrastructure/nats"
	"github.com/sss-eda/lemi025/controlling/infrastructure/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	nc, err := natsio.Connect("nats://sansa.dev:4222, nats://sansa.dev:4223, nats://sansa.dev:4224")
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

	// commandRouter := mux.NewRouter()

	// commandRouter.HandleFunc(
	// 	"/commands/read-config",
	// 	rest.HandleCommand( // returns http.HandlerFunc
	// 		controlling.ExecuteCommand( // This might be usefull in terms of directing
	// 			serial.ReadConfig(port), // returns
	// 		),
	// 	),
	// ).Methods("POST")

	// commandRouter.HandleFunc(
	// 	"/commands/read-time",
	// 	rest.CommandHandlerFunc( // returns http.HandlerFunc
	// 		controlling.HandleCommands( // This might be usefull in terms of directing
	// 			serial.ReadTime(port), // returns func(Command) error
	// 		),
	// 	),
	// ).Methods("POST")

	instrument := controlling.Instrument{}

	sub1, err := nc.Subscribe(
		"dev.lemi025.1.commands.readconfig",
		nats.CommandHandlerFunc(
			instrument.ReadConfig(
				serial.ConfigReader(port),
			),
		),
	)
	if err != nil {
		log.Println(err)
		cancel()
	}
	defer sub1.Unsubscribe()

	sub2, err := nc.Subscribe(
		"dev.lemi025.1.commands.readtime",
		nats.CommandHandlerFunc(
			instrument.ReadTime(
				serial.TimeReader(port),
			),
		),
	)
	if err != nil {
		log.Println(err)
		cancel()
	}
	defer sub2.Unsubscribe()

	func(event *controlling.ConfigReadEvent) {
		js.Publish(
			"dev.lemi025.1.events.configread",
			instrument.OnTimeRead(), // Returns []byte
		),
	}

	nc.Pu

	// err = serial.Acquire(
	// 	port,
	// 	serial.OnConfigRead(
	// 		jetstream.EventDispatcher[controlling.ConfigReadEvent](
	// 			js, "dev.lemi025.1.events.configread",
	// 		)),
	// 	serial.OnTimeRead(
	// 		jetstream.EventDispatcher[controlling.TimeReadEvent](
	// 			js, "dev.lemi025.1.events.timeread",
	// 		)),
	// 	serial.OnTimeSet(
	// 		jetstream.EventDispatcher[controlling.TimeSetEvent](
	// 			js, "dev.lemi025.1.events.timeset",
	// 		)),
	// 	serial.DataFrameAcquiredEventHandler(
	// 		jetstream.EventDispatcher[controlling.DataFrameAcquiredEvent](
	// 			js, "dev.lemi025.1.events.dataframeacquired",
	// 		)),
	// )

	<-ctx.Done()
}
