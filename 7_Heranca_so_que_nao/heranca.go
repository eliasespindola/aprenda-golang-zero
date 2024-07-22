package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	pessoa := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(pessoa)

	estudante := estudante{pessoa, "Computacao", "USP"}
	fmt.Println(estudante.nome)
}
