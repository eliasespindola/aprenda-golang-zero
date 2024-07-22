package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Print("Escrevendo do pacote auxiliar")
	EscreverAuxiliar2()
	escreverPrivado()
}
