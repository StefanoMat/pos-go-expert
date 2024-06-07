package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("O resultado é:", s)
	fmt.Println(uuid.New())
}
