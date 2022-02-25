package nats

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
	readconfig.Request | readtime.Request
}

// HandleQuery TODO
func HandleQuery[ResponseType Response, RequestType Request](
	handle func(func(*ResponseType) error, *RequestType),
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
				data, err := json.Marshal(response)
				if err != nil {
					return err
				}
				return msg.Respond(data)
			},
			&request,
		)
	}
}
