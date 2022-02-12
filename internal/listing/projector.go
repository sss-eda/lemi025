package listing

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// Projector TODO
type Projector struct {
	repo Repository
}

// OnConfigRead TODO
func (projector *Projector) OnConfigRead(
	event instrument.ConfigReadEvent,
) {
	repo.Get
}

// Project - What will the caller look like?
func Project(
	js nats.JetStream,
) {
	// This is where it would be nice to only subscribe to a specific site's instruments and then only project to a local read model!!!
	sub, err := js.Subscribe(
		"marion.lemi025.*.events.>",
		func(msg *nats.Msg) {
			subject := strings.Split(msg.Subject, ".")
			id := subject[2]

			var event instrument.Event
			err := json.Unmarshal(msg.Data, &event)
			if err != nil {
				log.Println(err)
				return
			}
			instr, err := repo.GetInstrumentByID(id)
			if err != nil {
				log.Println(err)
				return
			}
			instr.On(event)
			repo.UpdateInstrument(id, instr)
		},
	)
	if err != nil {
		log.Println(err)
	}
	defer sub.Unsubscribe()
}
