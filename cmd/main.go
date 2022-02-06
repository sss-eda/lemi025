package main

import (
	"log"
	"os"
	"strings"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
	"github.com/sss-eda/lemi025/pkg/infrastructure/nats"
	"github.com/sss-eda/lemi025/pkg/infrastructure/serial"
	tarm "github.com/tarm/serial"
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

	// Maybe send the strategies in here already?
	// client, err := lemi025.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// serial.Drive(
	// 	port,
	// 	client.Drive(
	// 		nats.DriveClient(js),
	// 	),
	// )

	instrument, err := lemi025.New()
	if err != nil {
		log.Fatal(err)
	}

	// Handle usecase: ReadConfig
	readConfigSub, err := js.Subscribe(
		strings.Join([]string{"lemi025", id, "config", "read"}, "."),
		nats.ReadConfigCommandMsgHandler(
			instrument.Config.Read(
				serial.ReadConfig(port),
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer readConfigSub.Unsubscribe()

	// Handle usecase: ReadTime
	readTimeSub, err := js.Subscribe(
		strings.Join([]string{"lemi025", id, "time", "read"}, "."),
		nats.ReadTimeCommandMsgHandler(instrument.Time.Read(serial.ReadTime(port))),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer readTimeSub.Unsubscribe()
}
