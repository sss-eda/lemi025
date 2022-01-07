package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/pkg/instrumentation"
)

// Driver TODO
type Driver struct {
	nc      *nats.Conn
	service instrumentation.Service
}

// Drive TODO
func (driver *Driver) Drive(
	ctx context.Context,
) {
	sub, _ := driver.nc.Subscribe(
		"hermanus.lemi025.1.commands.config.read",
		func(msg *nats.Msg) {
			driver.ReadConfig()
		},
	)
	defer sub.Unsubscribe()

	<-ctx.Done()
}
