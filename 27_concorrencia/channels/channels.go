package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)

	//recebe os dados
	//Neste ponto ele vai esperar um valor que vai salvar na variavel mensagem
	//E neste ponto que vai ter a sincronizacao com a goroutine do escrever
	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	fmt.Println(mensagem)
	//}

	//Enquanto o canal estiver aberto ele vai receber e fecha sozinho
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for {
		//envia
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
