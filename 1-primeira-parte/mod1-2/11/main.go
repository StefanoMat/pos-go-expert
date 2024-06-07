package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func main() {
	stefano := Cliente{
		Nome:  "Stefano",
		Idade: 24,
		Ativo: true,
	}
	Desativacao(stefano)

	fmt.Println(stefano)
}
