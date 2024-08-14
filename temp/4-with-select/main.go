package main

import (
	"fmt"
	"time"
)

type Event struct {
	Msg string
}

func main() {
	ch1 := make(chan Event)
	ch2 := make(chan Event)

	go func() {
		count := 0
		for {
			ch1 <- Event{Msg: "Event from Azure ServiceBus " + string(count)}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		count := 0
		for {
			ch2 <- Event{Msg: "Event from Kafka " + string(count)}
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case e := <-ch1:
			fmt.Println("Recevied from ASB: ", e.Msg)
		case e := <-ch2:
			fmt.Println("Recevied from Kafka: ", e.Msg)
		case <-time.After(3 * time.Second):
			println("timeout")
			return
		}
	}
}
