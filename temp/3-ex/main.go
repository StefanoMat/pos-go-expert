package main

import "fmt"

func main() {
	ch := make(chan int)
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i
	}
	go sum(numbers, ch)
	fmt.Println(<-ch)
}

func sum(numbers []int, ch chan int) {
	totalSum := 0
	for _, number := range numbers {
		totalSum += number
	}
	ch <- totalSum
}
