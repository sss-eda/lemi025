package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/application/query"
)

// HandleQuery TODO
func HandleQuery(
	nc *nats.Conn,
	sub
) func(query.Gateway) (query.Handler, error) {
	return func(gateway query.Gateway) (query.Handler, error) {
		sub, err := nc.Subscribe(
			"",
			QueryMsgHandler(gateway)
		)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		return nil
	}
}

// QueryMsgHandler TODO
func QueryMsgHandler[Request query.Request](
	gateway func(Request) (query.Response, error),
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var request Request
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Println(err)
			return
		}
		response, err := gateway(request)
		if err != nil {
			log.Println(err)
			return
		}
		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			return
		}
		err = msg.Respond(data)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
