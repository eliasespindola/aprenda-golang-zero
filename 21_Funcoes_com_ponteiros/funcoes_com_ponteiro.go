package main

import "fmt"

// Qnd eu passo o valor sem o *, eu passo apenas uma copia
// Passando um parametro como copia
func inverterSinal(numero int) int {
	return numero * -1
}

// Passando a referencia do numero, ou seja vai alterar o que está sendo passado
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	//Usando ponteiros, podendo ser variaveis ou funcoes
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
