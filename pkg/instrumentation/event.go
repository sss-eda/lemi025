package instrumentation

import (
	"context"

	"github.com/google/uuid"
)

type Subscription interface {
	Cancel() error
}

type subscription struct {
	events chan<- []byte
	cancel func()
}

// Cancel TODO
func (sub *subscription) Cancel() {
	sub.cancel()
}

// ConfigReadEvent TODO
type ConfigReadEvent struct {
	subscriptions map[uuid.UUID]*subscription
	stream        chan []byte
}

func (event *ConfigReadEvent) Publish(
	ctx context.Context,
	data []byte,
) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case event.stream <- data:
		return nil
	}
}

// Subscribe TODO
func (event *ConfigReadEvent) Subscribe(
	ctx context.Context,
	data chan<- []byte,
) (Subscription, error) {
	subID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	event.subscriptions[subID] = &subscription{
		events: data,
		cancel: func() { delete(event.subscriptions, subID) },
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event.subscriptions[subID].events <- event.stream:
			}
		}

	}()

	return event.subscriptions[subID]
}
