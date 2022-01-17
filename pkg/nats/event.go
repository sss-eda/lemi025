package nats

import "github.com/sss-eda/lemi025/pkg/driving"

type Event struct {
	Type    driving.EventType `json:"type"`
	Payload []byte            `json:"payload"`
}
