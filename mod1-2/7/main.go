package main

import "fmt"

func main() {
	salarios := map[string]int{"Stefano": 25000, "Wesley": 180000}
	fmt.Println(salarios["Stefano"])

	// sal := make(map[string]int)
	// sal1 = map[string]int{}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
