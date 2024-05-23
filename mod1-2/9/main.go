package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 4, 51))
}

func sum(nums ...int) int {
	total := 0
	for _, numero := range nums {
		total += numero
	}
	return total
}
