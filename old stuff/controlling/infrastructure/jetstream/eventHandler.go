package jetstream

import (
	"github.com/sss-eda/lemi025/controlling"
)

// EventMessage TODO
type EventMessage[Event controlling.Events] struct {
	Event Event `json:"event"`
}

// // EventDispatcher TODO
// func EventDispatcher[Event controlling.Events](
// 	js nats.JetStream,
// ) func(*Event) error {
// 	return func(event *Event) error {
// 		request := EventRequest[Command]{}

// 		err := json.NewDecoder(r.Body).Decode(&request)
// 		if err != nil {
// 			response, err2 := json.Marshal(CommandFailureResponse{
// 				Error: err,
// 			})
// 			if err2 != nil {
// 				log.Println(err2)
// 			}
// 			w.Write(response)
// 			w.WriteHeader(http.StatusBadRequest)
// 		}

// 		err = handle(&request.Command)
// 		if err != nil {
// 			response, err2 := json.Marshal(CommandFailureResponse{
// 				Error: err,
// 			})
// 			if err2 != nil {
// 				log.Println(err2)
// 			}
// 			w.Write(response)
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}

// 		response, err := json.Marshal(CommandSuccessResponse{})
// 		w.Write(response)
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// EventHandlerFunc TODO
// func EventHandlerFunc[Event controlling.Events](
// 	js nats.JetStream,
// ) func(string, func(*Event) error) {
// 	return func(subject string, raise func(*Event) error) {
// 		msg, err := json.Marshal(event)
// 		if err != nil {
// 			fmt.Printf("Error: %v\n", err)
// 		}

// 		ack, err := js.Publish(subject, msg)
// 		if err != nil {
// 			fmt.Printf("Error: %v, Ack: %v\n", err, ack)
// 		}
// 	}
// }
