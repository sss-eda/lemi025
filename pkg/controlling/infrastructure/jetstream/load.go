package jetstream

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// Load TODO
func Load(
	js nats.JetStream,
	aggregateType string,
	id string,
) (*instrument.Instrument, error) {
	aggregate, err := instrument.New()
	if err != nil {
		return nil, err
	}
	subject := strings.Join(
		[]string{"dev", aggregateType, id},
		".",
	)
	sub, err := js.Subscribe(subject, func(msg *nats.Msg) {
		var event instrument.Event
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Println(err)
			return
		}
		aggregate.Raise(event)
	})
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	return &aggregate, nil
}
