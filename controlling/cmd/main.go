package main

import (
	"log"

	"github.com/gorilla/mux"
	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/controlling"
	"github.com/sss-eda/lemi025/controlling/infrastructure/serial"
	tarm "github.com/tarm/serial"
)

func main() {
	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	port, err := tarm.OpenPort(&tarm.Config{
		Name: "COM6",
		Baud: 115200,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	router := mux.NewRouter()

	// Let's say that the resource is the lemi itself. We can send commands to
	// it via POST or PUT
	router.HandleFunc(
		"/",
		rest.HandleQuery( // returns http.HandlerFunc
			controlling.Query( // This might be usefull in terms of directing
				serial.ReadConfig(port), // returns
				serial.ReadTime(port),
				serial.SetTime(port),
				serial.SetCoefficients1(port),
				serial.ReadCoefficients1(port),
				serial.SetCoefficients2(port),
				serial.ReadCoefficients2(port),
			),
		),
	).Methods("PATCH")

	// This is if we want to use NATS instead of http.
	// sub1, err := js.Subscribe(
	// 	"dev.lemi025.1.controlling.queries",
	// 	nats.QueryMsgHandler( // returns func(*natsio.Msg)
	// 		controlling.Query( // returns func(lemi025.ResponseWriter, lemi025.ReadConfigRequest)
	// 			serial.ConfigReaderFunc(port), // returns func() error
	// 		),
	// 		func(response *readconfig.Response) []byte {
	// 			return []byte{}
	// 		},
	// 	),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer sub1.Unsubscribe()
}
