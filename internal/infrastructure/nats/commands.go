package nats

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/internal/application"
)

// // Command TODO
// type Command struct{}

// // // ReadCoefficients1Response TODO
// // type ReadCoefficients1Response struct {
// // 	Mode  lemi025.Mode  `json:"mode"`
// // 	Uin   uint16        `json:"uin"`
// // 	Mode1 lemi025.Mode1 `json:"mode1"`
// // }

// CommandMsgHandler TODO - Here we actually shouldn't be sending the strategy itself. It should be an application level wrapper around the strategy? Because the application is where the validation will happen. Now it can be bypassed and serial can be directly injected into nats.
func CommandMsgHandler(
	dispatch func(application.Command) error,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		command := application.Command{}
		err := json.Unmarshal(msg.Data, &command)
		if err != nil {
			log.Println(err)
			return
		}
		err = dispatch(command)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
