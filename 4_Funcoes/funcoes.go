package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}

func calculos(x, y int) (int, int) {
	soma := x + y
	sub := x - y
	return soma, sub
}

func main() {
	soma := somar(2, 10)
	print(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("funcao")
	print(resultado)

	resultadoSoma, _ := calculos(25, 10)
	print(resultadoSoma)
}
