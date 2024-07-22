package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//Ponteiro é um referencia de memoria, eu jogo o endereço de memoria onde a variavel está salva
	var variavel3 int = 100
	//Guarda o endereço de memoria de um inteiro e com o & pega a referencia da memoria
	//Eu jogo o endereço da memoria
	var ponteiro *int = &variavel3

	//Neste memoria só vai mostrar o endereço da memoria
	fmt.Println(variavel3, ponteiro)
	//Neste momento ele vai fazer a desferenciação, assim vamos obter o valor que está na memoria
	fmt.Println(variavel3, *ponteiro)
}
