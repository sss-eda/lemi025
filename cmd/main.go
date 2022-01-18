package main

import (
	"log"

	tarm "github.com/tarm/serial"
)

func main() {
	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()
}
