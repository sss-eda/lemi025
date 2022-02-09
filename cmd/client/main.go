package main

import (
	"log"
	"os"
	"strings"

	"github.com/sss-eda/lemi025/internal/infrastructure/nats"
	"github.com/sss-eda/lemi025/internal/infrastructure/serial"
)

func main() {
	// Connect to NATS server
	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Convert NATS connection to Jetstream
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Open serial port
	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Read Instrument ID from the environment
	id := os.Getenv("INSTRUMENT_ID")
	siteName := os.Getenv("SITE_NAME")

	// Handle commands
	readConfigSub, err := js.Subscribe(
		strings.Join([]string{siteName, "lemi025", id, "commands"}, "."),
		nats.CommandMsgHandler(serial.Transmitter(port)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer readConfigSub.Unsubscribe()

	serial.Receiver(port)(instrument)

	// instruments := nats.NewInstrumentStore()
	// data := nats.NewDatumStore()

	// updater := updating.NewService(instruments)
	// acquirer := acquiring.NewService(data)

	// serial.Subscribe(port, updater, acquirer)
}
