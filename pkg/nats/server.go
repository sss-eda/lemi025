package nats

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/pkg/eventsourcing"
)

// Server TODO
type Server struct {
	*nats.Conn
}

// NewServer TODO
func NewServer(
	nc *nats.Conn,
) (*Server, error) {
	server := Server{
		nc,
	}

	return &server, nil
}

// Serve TODO
func (server *Server) Serve(
	ctx context.Context,
	aggregate eventsourcing.Aggregate,
) error {
	sub, err := server.Subscribe(
		"COMMANDS."+aggregate.ID().String(),
		server.handleCommand(aggregate.CommandHandler()),
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	return nil
}

func (server *Server) handleCommand(
	f func(eventsourcing.Command) error,
) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var command command
		err := json.Unmarshal(msg.Data, &command)
		if err != nil {
			log.Println(err)
			return
		}
		err = f(command)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// func (server *Server) handle(
// 	msg *nats.Msg,
// ) {
// 	data := command
// 	err := json.Unmarshal(msg.Data, &input)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	switch input := input.(type) {
// 	case driving.ReadConfig:
// 		controller.driver.SetTime(command)
// 	case driving.ReadTime:
// 	case driving.SetTime:
// 	}
// }
