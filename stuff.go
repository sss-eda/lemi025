package lemi025

// ReadConfigRequest TODO
type ReadConfigRequest struct{}

type ReadConfigResponse struct {
	Error error
}

type ReadConfigResponseWriter interface {
	Wrrite(ReadConfigResponse) error
}

type ReadTimeRequest struct{}

type ReadTimeResponse struct {
	Error error
}

// ReadTimeResponseWriter TODO
type ReadTimeResponseWriter interface {
	Write(ReadTimeResponse) error
}

type SetTimeRequest struct {
	Year   uint16
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

type SetTimeResponse struct {
	Error error
}

type SetTimeResponseWriter interface {
	Write(SetTimeResponse) error
}

type AcquireRequest struct{}

type AcquireResponse struct {
	Error   error
	Payload interface{}
}

type AcquireReponseWriter interface {
	Write(AcquireResponse) error
}
