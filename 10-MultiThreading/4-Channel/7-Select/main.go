package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

// T1
func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)
	var i int64 = 0

	//RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{i, "Hello from RabbitMQ"}
			ch1 <- msg
		}

	}()

	//Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(1 * time.Second)
			msg := Message{i, "Hello from Kafka"}
			ch1 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from RabbitMQ:%v\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from Kafka:%v\n", msg2)

		case <-time.After(3 * time.Second):
			println("timeout")

		}

	}
}
