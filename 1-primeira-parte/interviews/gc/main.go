package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	var arrayBefore10Mb [1310720]int
	arrayBefore10Mb[0] = 1

	var arrayAfter10Mb [1310721]int
	arrayAfter10Mb[0] = 1

	sliceBefore64 := make([]int, 8192)
	sliceOver64 := make([]int, 8193)
	sliceOver64[0] = sliceBefore64[0]
}
