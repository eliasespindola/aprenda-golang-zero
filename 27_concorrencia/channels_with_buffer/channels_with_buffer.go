package main

import "fmt"

func main() {
	//Falando que o canal tem a capacidade de 2
	// O canal com o buffer só vai bloquear quando ele atingir a capacidade maxima dele
	// Ele é ATÉ 2
	canal := make(chan string, 2)
	canal <- "Olá"
	mensagem := <-canal
	fmt.Println(mensagem)
}
