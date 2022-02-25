package main

import (
	"log"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
	"github.com/sss-eda/lemi025/pkg/jetstream"
	"github.com/sss-eda/lemi025/pkg/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	sub1, err := js.Subscribe(
		"dev.lemi025.*.readconfig",
		jetstream.HandleQuery( // returns func(*natsio.Msg)
			readconfig.Query( // returns func(lemi025.ResponseWriter, lemi025.ReadConfigRequest)
				serial.ConfigReaderFunc(port), // returns func() error
			),
			func(response *readconfig.Response) []byte {
				return []byte{}
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub1.Unsubscribe()

	sub2, err := js.Subscribe(
		"dev.lemi025.*.readtime",
		jetstream.HandleQuery( // returns func(*natsio.Msg)
			readtime.Query( // returns func(lemi025.ResponseWriter, lemi025.ReadConfigRequest)
				serial.TimeReaderFunc(port), // returns func() error
			),
			func(response *readtime.Response) []byte {
				return []byte{}
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer sub2.Unsubscribe()

	acquirer := serial.Acquire(port)
	acquirer.Subscribe(func(msg readconfig.CommandMessage) {

	})

	HandleCommand(
		acquire.Command(
			serial.Acquire(),
		)
	)

	// serial.Acquire(
	// 	port,
	// ) // returns func()

	// acquirer := serial.Acquire(port)

	// acquirer(
	// 	acquire.Command(js, js),
	// )
}
