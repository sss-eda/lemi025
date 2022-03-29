package main

import (
	"context"
	"log"

	"github.com/sss-eda/lemi025/internal/infrastructure/nats"
	"github.com/sss-eda/lemi025/internal/infrastructure/serial"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}

	nats.Control(ctx, nc, "lemi025.1.read-config", driver.ReadConfig(serial.ReadConfigAdapter(port)))
	nats.Control(ctx, nc, "lemi025.1.read-time", driver.ReadTime(serial.ReadTimeAdapter(port)))
	nats.Control(ctx, nc, "lemi025.1.set-time", driver.SetTime(serial.SetTimeAdapter(port)))

	serial.NewDriverAdapter(
		driver.OnConfigRead(nats.OnConfigReadAdapter(nc)),
		driver.OnTimeRead(nats.OnTimeReadAdapter(nc)),
		driver.OnTimeSet(nats.OnTimeSetAdapter(nc)),
	)
}
