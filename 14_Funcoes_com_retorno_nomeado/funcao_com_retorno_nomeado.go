package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	//Já retorna os valores direto
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, substracao := calculosMatematicos(10, 20)
	fmt.Println(soma, substracao)
}
