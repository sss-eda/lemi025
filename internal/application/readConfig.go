package application

// ReadConfigRequest TODO
type ReadConfigRequest struct{}

// ReadConfigResponse TODO
type ReadConfigResponse struct{}

// ReadConfig TODO
type ReadConfig func(ReadConfigRequest) (ReadConfigResponse, error)
