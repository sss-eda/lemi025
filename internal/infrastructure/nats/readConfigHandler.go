package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/application/readconfig"
)

// ReadConfigHandler TODO
func ReadConfigHandler(
	nc *nats.Conn,
) func(readconfig.Gateway) readconfig.Handler {
	return func(gateway readconfig.Gateway) (readconfig.Handler, error) {
		sub, err := nc.Subscribe(
			"lemi025.1.config.read",
			func(msg *nats.Msg) {
				request := readconfig.Request{}
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
				err = msg.Respond(response.String())
				if err != nil {
					log.Println(err)
					return
				}
			},
		)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		return nil
	}
}
