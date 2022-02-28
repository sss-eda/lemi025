package controlling

import "io"

// Instrument TODO
type Instrument struct {
	rwc io.ReadWriteCloser
}

// RequestType TODO
type RequestType string

const (
	// ReadConfig TODO
	ReadConfig RequestType = "Read Config"
	// ReadTime TODO
	ReadTime RequestType = "Read Time"
	// SetTime TODO
	SetTime RequestType = "Set Time"
)

// RequestPayload TODO
type RequestPayload interface {
	ReadConfigRequest | ReadTimeRequest | SetTimeRequest
}

// Request TODO
type Request[T RequestPayload] struct {
	Type    RequestType
	Payload T
}

// Response TODO
type Response struct {
	Error error
}

// Query TODO
func Query[T RequestPayload]() func(func(*Response), *Request[T]) {
	return func(respond func(*Response), request *Request[T]) {
		request.Payload
	}
}

// ReadConfigRequest TODO
type ReadConfigRequest struct {
}

// ReadTimeRequest TODO
type ReadTimeRequest struct {
}

// SetTimeRequest TODO
type SetTimeRequest struct {
}

// nc.Subscribe(
// 	"dev.lemi025.1.queries",
// 	nats.QueryMsgHandler(
// 		controlling.Query(
// 			serial.ReadConfig(port),
// 		),
// 	),
// )
