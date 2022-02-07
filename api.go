package lemi025

import "github.com/nats-io/nats.go"

// Connection TODO
type Connection interface {
	ReadConfig(ReadConfigCommand) error
	ReadTime(ReadTimeCommand) error
}

// Connect TODO
func Connect(
	storeType string,
	apiType string,
) (Connection, error) {
	var store Store

	switch storageType {
	case "jetstream":
		nc, err := nats.Connect("nats://sansa.dev:4222")
		if err != nil {
			return nil, err
		}
		defer nc.Close()

		js, err := nc.JetStream()
		if err != nil {
			return nil, err
		}

		store, err = jetstream.NewStore(js)
		if err != nil {
			return nil, err
		}
	}

	switch apiType {
	case "serial":

	case "openapi":
	case "graphql":
	}
}
