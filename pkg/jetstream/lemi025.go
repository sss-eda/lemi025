package jetstream

import (
	"context"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// Lemi025 TODO - The Client?
type Lemi025 struct {
	commander lemi025.Commander
	nc        *nats.Conn
	site      string
	id        uint64
}

// Run TODO
func (l *Lemi025) Run(
	ctx context.Context,
) {
	subject := l.site + ".lemi025." + strconv.FormatUint(l.id, 10)

	sub, _ := l.nc.Subscribe(
		subject+".commands",
		l.readConfig,
	)
	defer sub.Unsubscribe()

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

// ReadConfig TODO
func (l *Lemi025) readConfig(msg *nats.Msg) {
	l.commander.ReadConfig()
}

// On TODO
func (l *Lemi025) On(
	event lemi025.Event,
) {
	subject := l.site + ".lemi025." + strconv.FormatUint(l.id, 10)

	l.nc.Publish(subject+".events", event.Data())
}
