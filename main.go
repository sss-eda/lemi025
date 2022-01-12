package main

import (
	"context"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/storage/local"
)

func main() {
	// We can also have the args serve as device repo...?
	devices, err := local.NewDeviceRepository("/home/engineer/lemi025/devices")
	if err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect("nats://10.160.11.3:4222")
	if err != nil {
		log.Fatal(err)
	}

	device, err := devices.Load(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	connection, err := gpio.Connect(device)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	go device.Drive(ctx, connection)
	go device.Observe(ctx, nc)

	<-ctx.Done()
}
