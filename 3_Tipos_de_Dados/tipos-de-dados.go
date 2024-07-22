package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero = 100
	numero2 := 100
	fmt.Println(numero)
	fmt.Println(numero2)
	//alias
	// INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)
	// BYTE = UNIT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//Numeros reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	//Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var erro error = errors.New("Erro")
	print(erro)
}
