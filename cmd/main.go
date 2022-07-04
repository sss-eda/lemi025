package main

import (
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
	tarm "github.com/tarm/serial"
)

func main() {
	// This is one client of the instrumentation service.
	port, _ := tarm.OpenPort(&tarm.Config{
		Name: "COM1",
		Baud: 19200,
	})

	nc, _ := natsio.Connect("wss://jetstream.sansa.dev/instrumentation")
	js, _ := nc.JetStream()

	js.Subscribe("instrumentation.lemi025.1.commands.read-config", serial.NewReadConfigHandlerFunc(port))
	js.Subscribe("instrumentation.lemi025.1.commands.read-time", serial.NewReadTimeHandlerFunc(port))

	station := serial.NewStation(port)

	instrument.Subscribe(lemi025.ConfigReadEventKind, jetstream.NewDispatchEventFunc("instrumentation.lemi025.1.events.config-read"))
	instrument.Subscribe(lemi025.TimeReadEventKind, jetstream.NewDispatchEventFunc("instrumentation.lemi025.1.events.time-read"))
}
