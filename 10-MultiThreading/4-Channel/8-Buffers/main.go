package main

type Message struct {
	Id  int64
	Msg string
}

// T1
func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
