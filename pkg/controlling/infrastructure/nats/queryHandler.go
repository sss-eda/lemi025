package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

// QueryMsgHandler TODO
func QueryMsgHandler(
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
