package nats

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

type driver struct {
	js nats.JetStream
	id string
}

// Drive TODO
func Drive(
	js nats.JetStream,
	id string,
) lemi025.DriveClientStrategy {
	return &driver{js, id}
}

// OnConfigRead TODo
func (d *driver) OnConfigRead(
	input lemi025.OnConfigReadInput,
) {
	_, err := d.js.Publish(
		strings.Join([]string{"lemi025", d.id, "events", "configRead"}, "."),
		input,
	)
	if err != nil {
		log.Println(err)
	}
}

// OnTimeRead TODo
func (d *driver) OnTimeRead(
	input lemi025.OnTimeReadInput,
) {
	_, err := d.js.Publish(
		strings.Join([]string{"lemi025", d.id, "events", "timeRead"}, "."),
		input,
	)
	if err != nil {
		log.Println(err)
	}
}
