package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	//Ordem é
	//Entrando na funcao
	//Funcao com o defer que é a ultima antes do retorno
	// retorno
	defer fmt.Println("Media calculada.Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false

}

func main() {
	//Adia a execucao de um determinado pedaco de codigo, pode ser funcao minha ou de package

	//Defer = Adiar
	//Ele vai adiar a execucao da funcao1 até o ultimo momento
	//O ultimo momento da funcao main seria antes dela terminar pq ela nao tem retorno

	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(1.0, 2.0))
}
