package nats

import (
	"testing"

	natsio "github.com/nats-io/nats.go"
)

func TestPresenter(t *testing.T) {
	nc, err := natsio.Connect("nats://sansa.dev:4222")
	if err != nil {
		t.Fatal(err)
	}

	presenter := Presenter[*someStruct](nc, "testing")
	err = presenter(&someStruct{"Hello World!"})
	if err != nil {
		t.Fatal(err)
	}
}

type someStruct struct {
	content string
}

func (ss *someStruct) MarshalBinary() ([]byte, error) {
	return []byte(ss.content), nil
}
