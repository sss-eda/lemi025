package main

import (
	"net/http"

	"github.com/tarm/serial"
)

func main() {
	port, _ := serial.OpenPort(&serial.Config{
		Name: "COM6",
		Baud: 115200,
	})
	// connection, _ := gpio.Connect(port)

	config := lemi025.NewConfig()
	readConfigCommand := gpio.NewReadConfigCommand(port, config) // gets invoked by a lemi025, the client is NATS or CLI or HTTP/REST or whatever

	time := lemi025.NewTime()
	readTimeCommand := gpio.NewReadTimeCommand(port, time)
	setTimeCommand := gpio.NewSetTimeCommand(port, time)

	driver := driving.New(
		readConfigCommand,
		ReadTimeCommand,
		SetTimeCommand,
	)

	driver.Serve(http.Server{})
}
