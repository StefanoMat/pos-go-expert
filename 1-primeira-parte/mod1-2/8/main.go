package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(1, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	var result int = a + b
	if result >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return result, nil
}
