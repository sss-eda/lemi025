package lemi025

// Message TODO
type Message interface {
	AggregateID() uint64
	Payload() []byte
}

// BaseMessage TODO
type BaseMessage struct {
	aggregateID uint64
	payload     []byte
}

// AggregateID TODO
func (message *BaseMessage) AggregateID() uint64 {
	return message.aggregateID
}

// Payload TODO
func (message *BaseMessage) Payload() []byte {
	return message.payload
}
