package driving

// Time TODO
type Time struct {
	TimeUpdatedEvent
	Year  uint16
	Month uint8
	Day   uint8
}

// TimeUpdatedEvent TODO
type TimeUpdatedEvent struct {
	entity       *Time
	handlerFuncs []func(TimeUpdatedEventPayload)
}

// TimeUpdatedEventPayload TODO
type TimeUpdatedEventPayload struct {
	Year  uint16
	Month uint8
	Day   uint8
}

// Register TODO
func (event *TimeUpdatedEvent) Register(
	handler func(TimeUpdatedEventPayload),
) {
	event.handlerFuncs = append(event.handlerFuncs, handler)
}

// Trigger TODO
func (event *TimeUpdatedEvent) Trigger(
	payload TimeUpdatedEventPayload,
) {
	for _, handle := range event.handlerFuncs {
		handle(payload)
	}
}
