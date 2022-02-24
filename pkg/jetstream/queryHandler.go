package jetstream

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readconfig"
	"github.com/sss-eda/lemi025/internal/domain/usecases/readtime"
)

type Response interface {
	readconfig.Response
}

type Request interface {
	readconfig.Request |
		readtime.Request
}

// HandleQuery TODO
func HandleQuery[ResponseType Response, RequestType Request](
	handle func(func(*ResponseType) error, *RequestType),
	serialise func(*ResponseType) []byte,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var request RequestType
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Println(err)
			return
		}
		handle(
			func(response *ResponseType) error {
				return msg.Respond(serialise(response))
			},
			&request,
		)
	}
}
