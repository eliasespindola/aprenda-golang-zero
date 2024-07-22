package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	idade      uint8
}

func main() {
	fmt.Println("Arquivo structs")

	endereco := endereco{"Rua dos bobos", 0}

	var u usuario
	u.nome = "Elias"
	u.idade = 21
	fmt.Println(u)

	usuario2 := usuario{"Elias", 21, endereco}
	fmt.Println(usuario2)
	//Maneira para preencher só um campo
	usuario3 := usuario{nome: "Elias"}
	fmt.Println(usuario3)
}
