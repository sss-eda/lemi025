package main

import (
	"context"
	"log"

	"github.com/sss-eda/lemi025/internal/domain"
	"github.com/sss-eda/lemi025/internal/infrastructure/nats"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}

	presenter := nats.Presenter[domain.ReadConfigResponse](nc, "lemi025.config.events")
	nats.Control(ctx, nc, "lemi025.readconfig", func(request *domain.ReadConfigRequest) error {
		serial.ReadConfig(port)
		presenter(request)
	})
}
