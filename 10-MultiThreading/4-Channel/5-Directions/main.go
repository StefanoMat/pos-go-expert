package main

import "fmt"

// T1
func main() {
	hello := make(chan string)
	go recebe("Stefano", hello)
	ler(hello)
}

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}
