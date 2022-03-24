package openapi

import "github.com/sss-eda/lemi025/internal/receiving"

// Message TODO
type Message struct {
	Type    receiving.MessageTypes
	Payload []byte
}
