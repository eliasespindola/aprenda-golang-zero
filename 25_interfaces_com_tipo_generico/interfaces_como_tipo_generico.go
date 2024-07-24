package main

import "fmt"

// Interface como tipo generico
// !!CUIDADO PARA NAO VIRAR UMA GAMBIARRA
// Isso está generico, pq tudo atende
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
	//Println recebe varias interfaces para funcionar por exemplo
}
