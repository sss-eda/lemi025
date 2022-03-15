package datapoints

import "context"

// EventSource TODO
type EventSource interface {
	Subscribe(InstrumentID)
}

// Subscribe TODO
func Subscribe(
	ctx context.Context,
	eventsource EventSource,
	stream chan<- *DataPoint,
) error {

	return nil
}
