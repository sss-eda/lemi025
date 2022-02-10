package command

import "context"

// UpdateConfigHandler TODO
type UpdateConfigHandler struct {
	instruments instrument.Repository
}

// NewUpdateConfigHandler TODO
func NewUpdateConfigHandler(
	instrumentRepo instrument.Repository,
) UpdateConfigHandler {
	return UpdateConfigHandler{
		instruments: instrumentRepo,
	}
}

// Handle TODO
func (handler UpdateConfigHandler) Handle(
	ctx context.Context,
	newConfig instrument.Config,
) error {
	instrument := handler.instruments.Load(newConfig)
}
