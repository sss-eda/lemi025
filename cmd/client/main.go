package main

import (
	"context"
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
	"github.com/sss-eda/lemi025/internal/infrastructure/json"
	"github.com/sss-eda/lemi025/internal/infrastructure/nats"
	"github.com/sss-eda/lemi025/internal/infrastructure/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nc, err := natsio.Connect("nats://sansa.dev:4222")
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

	nats.Subscribe(ctx, nc, "lemi025.1.commands.readConfig", json.Deserialize[lemi025.ReadConfigCommandPayload], serial.ReadConfig(port))
	nats.Subscribe(ctx, nc, "lemi025.1.commands.readTime", json.Deserialize[lemi025.ReadTimeCommandPayload], serial.ReadTime(port))
	nats.Subscribe(ctx, nc, "lemi025.1.commands.setTime", json.Deserialize[lemi025.SetTimeCommandPayload], serial.SetTime(port))

	serial.Subscribe(ctx, port,
		nats.Publish(nc, "lemi025.1.events.datumAcquired", json.Serialize[lemi025.DatumAcquiredEventPayload]),
		nats.Publish(nc, "lemi025.1.events.configRead", json.Serialize[lemi025.ConfigReadEventPayload]),
		nats.Publish(nc, "lemi025.1.events.timeRead", json.Serialize[lemi025.TimeReadEventPayload]),
		nats.Publish(nc, "lemi025.1.events.timeSet", json.Serialize[lemi025.TimeSetEventPayload]),
	)
}
