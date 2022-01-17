package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/pkg/driving"
)

// Client TODO
type Client struct {
	*nats.Conn
}

func NewClient(
	nc *nats.Conn,
) (*Client, error) {
	client := Client{
		nc,
	}

	return &client, nil
}

// Drive TODO
func (client *Client) Drive(
	ctx context.Context,
	input driving.Drive,
) error {
	sub, err := client.Subscribe("EVENTS.lemi025.1", func(msg *nats.Msg) {
		var event Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Println(err)
			return
		}

		switch event.Type {
		case driving.ConfigRead:
			var configReadEvent driving.ConfigReadEvent
			err = json.Unmarshal(event.Payload, &configReadEvent)
			if err != nil {
				log.Println(err)
				return
			}
			err = input.OnConfigRead(configReadEvent)
			if err != nil {
				log.Println(err)
				return
			}
		case driving.TimeRead:
			var timeReadEvent driving.TimeReadEvent
			err = json.Unmarshal(event.Payload, &timeReadEvent)
			if err != nil {
				log.Println(err)
				return
			}
			err = input.OnTimeRead(timeReadEvent)
			if err != nil {
				log.Println(err)
				return
			}
		case driving.TimeSet:
			var timeSetEvent driving.TimeSetEvent
			err = json.Unmarshal(event.Payload, &timeSetEvent)
			if err != nil {
				log.Println(err)
				return
			}
			err = input.OnTimeSet(timeSetEvent)
			if err != nil {
				log.Println(err)
				return
			}
		}
	})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	<-ctx.Done()

	return ctx.Err()
}

// ReadConfig TODO
func (client *Client) ReadConfig(
	input driving.ReadConfig,
) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = client.Publish("COMMANDS.lemi025.1.config.read", data)
	if err != nil {
		return err
	}

	return nil
}
