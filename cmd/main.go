package main

import (
	"log"
	"net/http"

	"github.com/sss-eda/lemi025"
)

func main() {
	tx := serial.NewTransmitter()
	rx := serial.NewReceiver()

	connection := lemi025.Connect(&lemi025.Config{
		Transceiver: tx,
		Recevier:    rx,
	})

	api := rest.NewAPI(connection)

	log.Fatal(http.ListenAndServe(":8080", api))
}
