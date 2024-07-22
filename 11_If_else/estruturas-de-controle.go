package main

import "fmt"

func main() {
	fmt.Println("Hola")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("numero é maior que zero")
	} else if outronumero < -10 {
		fmt.Println("menor ou igual a zero")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
