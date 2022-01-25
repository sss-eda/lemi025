package nats

import (
	"github.com/nats-io/nats.go"
)

// Client TODO
type Client struct {
	nc *nats.Conn
}

// NewClient TODO
func NewClient(
	natsConn *nats.Conn,
) (*Client, error) {
	return &Client{
		nc: natsConn,
	}, nil
}

// ReadConfig TODO
func (client *Client) ReadConfig(
	id string,
) error {
	return readConfig(client.nc, id)
}
