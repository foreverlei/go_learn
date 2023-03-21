package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"learn/common"
)

func main() {
	nc, _ := nats.Connect("nats://10.11.1.120:4222")
	defer nc.Drain()

	//// Simple Publisher
	//nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo/test", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	common.CaptureStopSingal("Receive")

}
