package jetstream

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
)

// Store TODO
type Store struct {
	js        nats.JetStream
	snapshots SnapshotRepository
}

// SnapshotRepository TODO
type SnapshotRepository interface {
	Load(ID) (Aggregate, error)
	Save(ID, Aggregate) error
}

// ID TODO
type ID interface {
	Equals(ID) bool
	String() string
}

// Aggregate TODO
type Aggregate interface {
}

// Load TODO
func (store *Store) Load(id ID) (interface{}, error) {
	aggregate, err := instrument.New()
	if err != nil {
		return nil, err
	}

	sub, err := store.js.Subscribe(
		// How are we going to maintain the versioning of the events if they
		// are being published to different subjects? Does it then go by
		// stream?
		strings.Join([]string{"lemi025", id.String(), "events", ">"}, "."),
		func(msg *nats.Msg) {
			var event instrument.Event
			err := json.Unmarshal(msg.Data, &event)
			if err != nil {
				log.Println(err)
				return
			}
			aggregate.Raise(event)
		},
	)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	return nil, nil
}
