package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool
	c int
	d string = "Y"
	f ID     = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray {
		fmt.Printf("O valor de %d é %d\n", i, v)
	}

	fmt.Println(meuArray[2])
}
