package main

import "fmt"

//Sao parecidos com funcoes, mas o metodo obrigatoriamente está associado a alguma coisa

type usuario struct {
	nome  string
	idade uint8
}

// Atribuindo o metodo ao usuario
func (u usuario) salvar() {
	fmt.Println(u.nome)
}

// Mudando atributos do usuario
func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario1.aniversario()
	fmt.Println(usuario1.idade)
}
