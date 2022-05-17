package main

import (
	"context"
	"log"

	tarm "github.com/tarm/serial"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// nc, err := natsio.Connect("nats://sansa.dev:4222")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	gateway := websocket.NewInstrumentationGateway("ws://localhost:8080/ws")
	err := instrumentation.Register(gateway)
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM5",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}

	// go nats.Serve(ctx, nc,
	// 	nats.WithEndpoint("lemi025.1.commands.readConfig", lemi025.ReadConfig(port)),
	// 	nats.WithEndpoint("lemi025.1.commands.readCoefficients", lemi025.ReadCoefficients(port)),
	// )

	// go serial.Serve(ctx, port,
	// 	serial.WithConfigReadEventHandler(lemi025.OnConfigRead(nats.Publish())),
	// )

	// nats.Subscribe(ctx, nc, "lemi025.1.commands.readConfig", serial.ReadConfig(port))
	// // websocket.Subscribe(ctx, ws, "/lemi025/1/commands/readConfig", serial.ReadConfig(port))
	// nats.Subscribe(ctx, nc, "lemi025.1.commands.readTime", serial.ReadTime(port))
	// nats.Subscribe(ctx, nc, "lemi025.1.commands.setTime", serial.SetTime(port))

	// serial.Subscribe(ctx, port,
	// 	nats.Publish(nc, "lemi025.1.events.datumAcquired", json.Serialize[lemi025.DatumAcquiredEventPayload]),
	// 	nats.Publish(nc, "lemi025.1.events.configRead", json.Serialize[lemi025.ConfigReadEventPayload]),
	// 	nats.Publish(nc, "lemi025.1.events.timeRead", json.Serialize[lemi025.TimeReadEventPayload]),
	// 	nats.Publish(nc, "lemi025.1.events.timeSet", json.Serialize[lemi025.TimeSetEventPayload]),
	// )
}
