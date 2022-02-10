package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

// ReadConfigMsgHandler TODO
func ReadConfigMsgHandler(
	ctx context.Context,
	repo domain.ReadRepository,
	gateway func(),
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		instrument := repo.GetLemi025(request.InstrumentID)
		instrument.Config.Read()
	}
}
