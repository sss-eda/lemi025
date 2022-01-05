package main

import (
	"github.com/sss-eda/lemi025"
	"github.com/sss-eda/lemi025/pkg/gpio"

	"github.com/tarm/serial"
)

func main() {
	port, _ := serial.OpenPort(&serial.Config{
		Name: "COM6",
		Baud: 115200,
	})
	connection, _ := gpio.Connect(port)

	readConfigCommand := lemi025.ReadConfig(connection)
	readTimeCommand := lemi025.ReadTime(connection)

	api.Serve(readConfigCommand)

	// lemi025.Drive(connection, driver)     // the serial driver
	// lemi025.Observe(connection, observer) // the nats publisher

	// commander, _ := gpio.NewCommander(connection)

	// observer := &nats.Observer{}

	// driver, _ := gpio.NewDriver(connection)
	// driver.Drive(observer)

}
