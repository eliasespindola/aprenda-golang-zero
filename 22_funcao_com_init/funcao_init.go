package main

import "fmt"

var n int

// Funcao init que é executada antes da funcao main
// Posso ter uma funcao POR ARQUIVO, podemos fazer setup e essas coisas
func init() {
	fmt.Println("init called")
	n = 10
}

func main() {
	fmt.Println("My fazer com init")
	fmt.Println(n)
}
