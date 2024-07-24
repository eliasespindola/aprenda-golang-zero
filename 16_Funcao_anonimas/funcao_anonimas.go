package main

import "fmt"

func main() {
	//Funcao anonima que é uma funcao que nao tem nome para invocar
	func(texto string) {
		fmt.Println("Olá mundo!", texto)
	}("Texto da funcao anonima")
	//Como fazer para executar colocamos os paranteses, ai vai ter a declaracao e vai executar ela

	//Podemos fazer que ela retorne um valor
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
