package main

import (
	"log"
	"os"
	"strconv"

	tarm "github.com/tarm/serial"
)

const (
	defaultSerialName string = "COM5"
	defaultSerialBaud string = "115200"
)

func main() {
	name, ok := os.LookupEnv("SERIAL_NAME")
	if !ok {
		name = defaultSerialName
	}

	baudEnv, ok := os.LookupEnv("SERIAL_BAUD")
	if !ok {
		baudEnv = defaultSerialBaud
	}
	baud, err := strconv.Atoi(baudEnv)
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

	serverType, ok := os.LookupEnv("SERVER_TYPE")
	if !ok {
		serverType = "nats"
	}
	switch serverType {
	case "nats":
		// Get the API spec from the api dir
		nats.Subscribe("lemi025.1.commands.readconfig", nats.CommandMsgHandleFunc(serial.ReadConfig(port)))
		nats.Subscribe("lemi025.1.commands.readtime", nats.CommandMsgHandleFunc(serial.ReadTime(port)))
		nats.Subscribe("lemi025.1.commands.settime", nats.CommandMsgHandleFunc(serial.SetTime(port)))
	case "openapi":

	}
}
