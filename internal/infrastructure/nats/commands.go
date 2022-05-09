package nats

import (
	"context"
	"errors"
	"io"

	natsio "github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025"
)

// MessageBus TODO
type MessageBus struct {
	connection    *natsio.EncodedConn
	subscriptions []*natio.Subscription
	streams       map[lemi025.MessageType]string
}

// NewMessageBus TODO
func NewMessageBus(ec *natsio.EncodedConn, subj string) (*MessageBus, error) {
	return &MessageBus{
		connection: ec,
		subject:    subj,
	}, nil
}

// Publish TODO
func (bus *MessageBus) Publish(ctx context.Context, message lemi025.Message) error {
	if !bus.connection.Conn.IsConnected() {
		return errors.New("not connected")
	}

	return bus.connection.Publish(bus.subject, message)
}

// Subscribe TODO
func (bus *MessageBus) Subscribe(
	ctx context.Context,
	kind lemi025.MessageKind,
	messages chan<- lemi025.Message,
) error {
	sub, err := bus.connection.Conn.ChanSubscribe(kind.String(), func(message lemi025.Message) {
		select {
		case messages <- message:
		case <-ctx.Done():
		}
	})
	if err != nil {
		return err
	}

	bus.subscriptions = append(bus.subscriptions, sub)

	return nil
}

// ReadConfigCommandHandler - Does not belong here. NATS is just for MESSAGES.
// func CommandHandler[](nc *natsio.Conn, subject string) func(context.Context, *lemi025.ReadConfigCommand) error {
// 	return func(ctx context.Context, command *lemi025.ReadConfigCommand) error {
// 		message, err := command.MarshalBinary()
// 		if err != nil {
// 			return err
// 		}

// 		return nc.Publish(subject, message)
// 	}
// }

// ReadTime TODO
func ReadTime(port io.Writer) func(context.Context, *lemi025.ReadTimeCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.ReadTimeCommandPayload) error {
		_, err := port.Write([]byte{0x3D, 0x31})
		return err
	}
}

// SetTime TODO
func SetTime(port io.Writer) func(context.Context, *lemi025.SetTimeCommandPayload) error {
	return func(ctx context.Context, payload *lemi025.SetTimeCommandPayload) error {
		_, err := port.Write([]byte{0x3D, 0x31})
		return err
	}
}
