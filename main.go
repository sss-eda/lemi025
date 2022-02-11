package main

import (
	"github.com/nats-io/nats.go"
)

func main() {
	handler := nats.QueryMsgHandler()
}

// type Hello interface {
// 	ReadConfig | ReadTime
// }

// type ReadConfig struct{}

// type ReadTime struct{}

// func SayHello[H Hello](hello H) {
// 	fmt.Printf("Hello %v!\n", hello)

// 	return
// }
