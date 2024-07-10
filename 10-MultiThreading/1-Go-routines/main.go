package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
	}
}

// Thread 1
func main() {
	//T2
	go task("A")
	//T3
	go task("B")

	time.Sleep(15 * time.Second)
}
