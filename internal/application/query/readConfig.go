package query

import "context"

// ReadConfigHandler TODO
type ReadConfigHandler struct {
	gateway ReadConfigGateway
}

// ReadConfigGateway TODO
type ReadConfigGateway interface {
	ReadConfig(ctx context.Context) error
}

// NewReadConfigHandler TODO
func NewReadConfigHandler(
	configReader ReadConfigGateway,
) ReadConfigHandler {
	return ReadConfigHandler{
		gateway: configReader,
	}
}

// Handle TODO
func (handler ReadConfigHandler) Handle(
	ctx context.Context,
) error {
	return handler.gateway.ReadConfig(ctx)
}
