package main

import (
	"fmt"
)

func main() {

	re := func() int {
		return sum(50, 5, 10) * 2
	}()
	fmt.Println(re)
}

func sum(nums ...int) int {
	total := 0
	for _, numero := range nums {
		total += numero
	}
	return total
}
