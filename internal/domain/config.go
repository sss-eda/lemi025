package domain

// Config TODO
type Config struct {
	StationNumber uint8
}

type ReadConfigUseCase func(ReadConfigPresenter, *ReadConfigRequest) error

type ReadConfigPresenter func(*ReadConfigResponse)

type ReadConfigRequest struct{}

type ReadConfigResponse struct{}
