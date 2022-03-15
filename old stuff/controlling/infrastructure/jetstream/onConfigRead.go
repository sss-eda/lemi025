package jetstream

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/controlling"
)

// OnEvent TODO
func OnEvent[Event controlling.Events](
	js nats.JetStream,
	subject string,
) func(*Event) error {
	return func(event *Event) error {
		data, err := json.Marshal(event)
		if err != nil {
			return err
		}
		_, err = js.Publish(subject, data)
		return err
	}
}
