package nats

// // ReadCoefficients1Request TODO
// type ReadCoefficients1Request struct{}

// // ReadCoefficients1Response TODO
// type ReadCoefficients1Response struct {
// 	Mode  lemi025.Mode  `json:"mode"`
// 	Uin   uint16        `json:"uin"`
// 	Mode1 lemi025.Mode1 `json:"mode1"`
// }

// // ReadCoefficients1CommandMsgHandler TODO - Here we actually shouldn't be sending the strategy itself. It should be an application level wrapper around the strategy? Because the application is where the validation will happen. Now it can be bypassed and serial can be directly injected into nats.
// func ReadCoefficients1CommandMsgHandler(
// 	strategy lemi025.ReadCoefficients1Strategy,
// ) func(*nats.Msg) {
// 	return func(msg *nats.Msg) {
// 		request := ReadTimeRequest{}
// 		err := json.Unmarshal(msg.Data, &request)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		err = strategy(lemi025.ReadCoefficients1Command{})
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
