package main

import (
	"fmt"
	"time"
)

// T1
func main() {
	data := make(chan int)
	go worker(1, data)
	go worker(2, data)
	go worker(3, data)

	for i := 0; i < 5; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received: %d\n", workerId, x)
		time.Sleep(time.Second * 1)
	}
}
