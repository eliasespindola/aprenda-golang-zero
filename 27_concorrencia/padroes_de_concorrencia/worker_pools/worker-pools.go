package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	//Antes mesmo de popular o channel de tarefas
	//Ele vai chamar varias vezes os workes, pq vamos ter dois processos que vao estar puxando numero dessa fila e fazendo o processamento

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// Podemos falar se o canal so envia ou só recebe
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}
