package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Stefano Kaefer"
	fmt.Printf("O cliente %v andou\n", c.nome)
}
func main() {
	stefano := Cliente{
		nome: "Stefano",
	}
	stefano.andou()
	fmt.Printf("O valor da struct com nome %v", stefano.nome)
}
