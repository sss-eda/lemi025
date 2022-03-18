package openapi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/sss-eda/lemi025/internal/receiving"
)

// DecoderFunc TODO
type DecoderFunc func([]byte) (*Message, error)

// Receive TODO
func Receive(
	ctx context.Context,
	station *receiving.Station,
	ws *websocket.Conn,
	decode DecoderFunc,
) {
	defer ws.Close()
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		message, err := decode(data)
		if err != nil {
			log.Println(err)
			continue
		}

		switch message.Type {
		case receiving.ConfigReadEventMessage:
			handleEvent()
			event := receiving.ConfigReadEvent{}
			err = json.Unmarshal(message.Payload, &event)
			if err != nil {
				log.Println(err)
				continue
			}
		case receiving.TimeReadEventMessage:
			event := receiving.TimeReadEvent{}
			err = json.Unmarshal(message.Payload, &event)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
