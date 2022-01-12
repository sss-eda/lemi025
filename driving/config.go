package driving

// Config TODO
type Config struct {
	ConfigUpdatedEvent
	stationNumber uint8
}

// ConfigUpdatedEvent TODO
type ConfigUpdatedEvent struct {
	handlers []interface {
		Handle(ConfigUpdatedEventPayload)
	}
}

// ConfigUpdatedEventPayload TODO
type ConfigUpdatedEventPayload struct {
	StationNumber uint8
}

// Register TODO
func (event *ConfigUpdatedEvent) Register(
	handler interface {
		Handle(ConfigUpdatedEventPayload)
	},
) {
	event.handlers = append(event.handlers, handler)
}

// Trigger TODO
func (event *ConfigUpdatedEvent) Trigger(
	payload ConfigUpdatedEventPayload,
) {
	for _, handler := range event.handlers {
		handler.Handle(payload)
	}
}
