package nats

import (
	"encoding"

	natsio "github.com/nats-io/nats.go"
)

// Presenter TODO
func Presenter[Response encoding.BinaryMarshaler](
	nc *natsio.Conn,
	subject string,
) func(Response) error {
	return func(response Response) error {
		data, err := response.MarshalBinary()
		if err != nil {
			return err
		}

		return nc.Publish(subject, data)
	}
}
