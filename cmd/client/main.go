package main

import (
	"log"
	"os"

	"github.com/tarm/serial"
)

func main() {
	// This is the client, in other words the code running on the instrument
	// itself.

	port, err := serial.OpenPort(&serial.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	instrumentID := os.Getenv("LEMI025_INSTRUMENT_ID")

	instruments := instrumentation.NewHTTPClient()

	instrument, err := instruments.GetInstrument(instrumentID)
	if err != nil {
		log.Fatal(err)
	}

}
