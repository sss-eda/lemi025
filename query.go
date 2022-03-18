package lemi025

// Queries can only be sent to readmodels. Each readmodel will have its own
// set of queries. Just like each write model has its own set of commands.

// Query TODO
func Query[Q Queries](
	query Q,
) {
	return nil
}

// Queries TODO
type Queries interface {
	GetConfigQuery
	Respond()
}

// GetConfigQuery TODO
type GetConfigQuery struct {
	ReponseWriter func(GetConfigResponse) error
	Request       GetConfigRequest
}

// GetConfigRequest TODO
type GetConfigRequest struct {
	StationNumber uint8
}

// GetConfigResponse TODO
type GetConfigResponse struct {
	Config *Config
}

// Config TODO
type Config struct{}
