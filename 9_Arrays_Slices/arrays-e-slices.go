package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Array é uma lista de valores
	//Todos os dados tem que ser do mesmo tipo
	var array1 [5]string
	array1[0] = "Posicao 1"
	array1[1] = "Posicao 2"
	fmt.Println(array1)

	array2 := [5]string{"Posicao1", "Posicao2"}
	fmt.Println(array2)
	//Array é inflexivel, para ser o maximo de ser flexivel
	// Os 3 pontos vai fixar o tamanho do array de acordo com a quantidade de items passados na inicializacao
	//Mas o array nao fica com o tamanho dinamico,se eu tentar adicionar o valor posteriormento
	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	//Usamos o slice, pq ele é mais flexivel
	//Slice nao é um array, mas ele aponta para o array
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	//Diferença de tipos
	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))
	//Vai adicionar um item no slice e retornar o slice
	slice = append(slice, 18)
	fmt.Println(slice)
	//pegando do indice 1 até o indice 2
	slice2 := array2[1:3]
	fmt.Println(slice2)
	//Se a gente alterar o array altera o slice tb, pq ele é um ponteiro que aponta para as posicoes do array
	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	//Arrays Internos
	//a FUNCAO MAKE QUE VAI LOCKAR UMA ESPACO NA MEMORIA
	//TIPO, TAMANHO, CAPACIDADE MAXIMA DE ITENS
	//Make vai ficar um array de 15 posicoes e vai nos retornar um slice de 10 posicoes
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Tamanho
	fmt.Println(cap(slice3)) //Capacidade
	//Quando ele ver que vai estourar o tamanho ele duplicar, ele vai pegar 11 + 1 e vai dar 12 e dps duplicar a capacidade
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
