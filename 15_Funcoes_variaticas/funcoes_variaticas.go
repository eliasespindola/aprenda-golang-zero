package main

import "fmt"

// Vai receber de 0 a N numeros
func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Nao posso ter mais de um parametro variatico por funcao
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	//Isso vai retornar um slice
	total := soma(1, 2, 3, 4, 5)
	fmt.Println(total)

	escrever("Olá Mundo", 1, 2, 3, 4, 5)
}
