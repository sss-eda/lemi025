package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/sss-eda/instruments"
	"github.com/sss-eda/lemi025"
)

func main() {
	instruments.Connect()

	idEnv, idEnvDefined := os.LookupEnv("LEMI025_ID")
	if !idEnvDefined {
		idEnv = "1"
	}
	if idEnv == "" {
		log.Fatal("ID may not be empty")
	}

	// Connecting to this endpoint will:
	//   -> Tell the server that the instrument is connected
	//   -> Open two-way comms between the server and the instrument
	//   -> We will listen for incoming commands
	//   -> We will transmit events
	ws, _, err := websocket.DefaultDialer.Dial("wss://api.sansa.dev/v0/instruments/"+idEnv, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(string(data))

		message := &Message{}
		err = json.Unmarshal(data, message)
		if err != nil {
			log.Println(err)
			continue
		}

		switch message.Type {
		case "command":
			var commandMessage CommandMessage
			json.Unmarshal(message.Payload, &commandMessage)
			switch commandMessage.Type {
			case "readconfig":
				command := lemi025.ReadConfigCommand{}
				json.Unmarshal(commandMessage.Payload, &command)
			case "readtime":
			case "settime":
			}
		case "event":
		}
	}
}

// Message TODO
type Message struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type CommandMessage struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}
