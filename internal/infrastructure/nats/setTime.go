package nats

// // SetTimeRequest TODO
// type SetTimeRequest struct {
// 	Year   uint16 `json:"year"`
// 	Month  uint8  `json:"month"`
// 	Day    uint8  `json:"day"`
// 	Hour   uint8  `json:"hour"`
// 	Minute uint8  `json:"minute"`
// 	Second uint8  `json:"second"`
// }

// // SetTimeResponse TODO
// type SetTimeResponse struct{}

// // SetTimeCommandMsgHandler TODO - Here we actually shouldn't be sending the strategy itself. It should be an application level wrapper around the strategy? Because the application is where the validation will happen. Now it can be bypassed and serial can be directly injected into nats.
// func SetTimeCommandMsgHandler(
// 	strategy domain.SetTimeStrategy,
// ) func(*nats.Msg) {
// 	return func(msg *nats.Msg) {
// 		request := ReadTimeRequest{}
// 		err := json.Unmarshal(msg.Data, &request)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		err = strategy(lemi025.SetTimeCommand{})
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }
