package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sss-eda/lemi025/pkg/driving"
	"github.com/sss-eda/lemi025/pkg/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	// stationName := os.Getenv("LEMI025_STATION_NAME")
	// stationNumber := os.Getenv("LEMI025_STATION_NUMBER")
	name := os.Getenv("LEMI025_SERIAL_NAME")
	baud, err := strconv.Atoi(os.Getenv("LEMI025_SERIAL_BAUD"))
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: name,
		Baud: baud,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	lemi025, err := serial.Connect(port)
	if err != nil {
		log.Fatal(err)
	}

	lemi025.Drive(driving.Drive{})

}
