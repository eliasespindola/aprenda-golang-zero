package main

import "fmt"

func fibonacci(posicao uint) uint {
	//Temos que falar qnd ele quer parar, para nao ter o estouro de pilha(famoso stack overflow)
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	//Funcao que vai retornar um determinado numero da funcao fibonati
	//posicao := fibonacci(12)
	//fmt.Println(posicao)

	posicao := uint(12)
	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
