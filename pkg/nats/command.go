package nats

import "github.com/sss-eda/lemi025/pkg/eventsourcing"

// command TODO
type command struct {
	Type    driving.CommandType `json:"type"`
	Payload []byte              `json:"payload"`
}

// Type TODO
func (c *command) Type() eventsourcing.CommandType {
	return c.Type
}
