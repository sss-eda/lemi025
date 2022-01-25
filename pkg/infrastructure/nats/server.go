package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/sss-eda/lemi025/pkg/application"
)

// Server TODO
type Server struct {
	nc      *nats.Conn
	service application.Server
}

// NewServer TODO
func NewServer(
	natsConn *nats.Conn,
) (*Server, error) {
	return &Server{
		nc: natsConn,
	}, nil
}

// Serve TODO
func (server *Server) Serve(
	subject string,
	id string,
) error {
	sub, err := server.nc.Subscribe("LEMI025."+id, server.readConfig)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	return nil
}

// func (adapter *Adapter) ReadConfig() error {
// 	request := lemi025.ReadConfigRequest{}
// 	data, err := json.Marshal(request)
// 	if err != nil {
// 		return err
// 	}

// 	adapter.nc.Publish("")
// }

func (server *Server) readConfig(msg *nats.Msg) {
	request := ReadConfigRequest{}
	err := json.Unmarshal(msg.Data, &request)
	if err != nil {
		err2 := msg.Respond([]byte(fmt.Sprintf("Error: %q", err.Error())))
		if err2 != nil {
			log.Println(err2)
			return
		}
		return
	}

	// 	lemi025.ReadConfig()
	// }
}
