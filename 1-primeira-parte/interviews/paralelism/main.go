package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

const (
	numberRange     = 1000000 // Range for finding prime numbers.
	goroutineNumber = 4       // Number of launched goroutines.
)

func main() {
	// Start tracing.
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	// Set the number of processors to be used.
	runtime.GOMAXPROCS(4)

	// Use WaitGroup to wait for the execution of all goroutines.
	wg := &sync.WaitGroup{}
	wg.Add(4)
	findInRange(wg) // Launch a CPU-bound task
	wg.Wait()
}

// Search for prime numbers in 4 ranges.
func findInRange(wg *sync.WaitGroup) {
	start := 0
	end := numberRange
	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Println(findPrimeNumbers(start, end))
			wg.Done()
		}(i)
		start += numberRange
		end += numberRange
	}
}

// Function that finds prime numbers in a specified range.
func findPrimeNumbers(start, end int) []int {
	var primes []int

	for num := start; num <= end; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}

	return primes
}

// Check if a number is prime.
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
