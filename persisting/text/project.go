package text

import (
	"context"

	lemi025 "github.com/sss-eda/lemi025/internal"
)

// ONLY ACK MSG ONCE IT HAS SUCCESSFULLY BEEN WRITTEN! So the exactly-once
// semantics can be handled by the primary adapter.

// Project TODO
func Project(
	ctx context.Context,
	configRead <-chan lemi025.ConfigReadEvent,
	timeRead <-chan lemi025.TimeReadEvent,
) {
	go directory.OnConfigRead(ctx, configRead)
	go directory.OnTimeRead(ctx, timeRead)
}

text.Project(
	context.Background(),
	nats.ConfigReadEvents(),
	nats.TimeReadEvents(),
)