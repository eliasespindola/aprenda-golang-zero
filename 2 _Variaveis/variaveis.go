package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lala"
		variavel4 string = "lele"
	)
	fmt.Println(variavel3, variavel4)
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	println(constante1)

	//Invertendo valores
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
