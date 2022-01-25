package nats

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

// ReadConfigRequest TODO
type ReadConfigRequest struct {
	ID string `json:"id"`
}

// ReadConfigResponse TODO
type ReadConfigResponse struct{}

// ReadConfig TODO
func readConfig(
	nc *nats.Conn,
	id string,
) error {
	request := ReadConfigRequest{
		ID: id,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	err = nc.Publish("LEMI025."+id, data)
	if err != nil {
		return err
	}
	return nil
}
