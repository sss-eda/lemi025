package nats

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/entities/instrument"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
)

// ReadConfig TODO
func ReadConfig(
	ctx context.Context,
	encode func(readconfig.Request) []byte,
	decode func([]byte) readconfig.Response,
	nc *nats.Conn,
	id string,
) error {
	request := readconfig.Request{}

	response, err := nc.RequestWithContext(
		ctx,
		"dev.lemi025."+id+".config.read",
		encode(request),
	)
	if err != nil {
		return err
	}

	return decode(response.Data).Error
}

// ConfigReaderFunc TODO
func ConfigReaderFunc(
	nc *nats.Conn,
	id instrument.ID,
) func(context.Context) error {
	return func(ctx context.Context) error {
		return ReadConfig(ctx, nc, id.String())
	}
}
