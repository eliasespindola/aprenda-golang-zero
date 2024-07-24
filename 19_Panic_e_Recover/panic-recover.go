package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	//Media nunca pode ser 6, se for 6 o programa vai entrar em panico
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//Panic vai interromper o fluxo normal do programa e parar tudo basicamente
	//Ela vai chamar todas as funcoes que tem defer
	//Se a gente nao recuperar com o recover o programa morre
	panic("A MEDIA É EXATAMENTE 6!")
}

func recuperarExecucao() {
	//A gente tenta recuperar a execucacao assim, e vamos conseguir recuperar com sucesso
	//Se a função nao estiver entrando em panico ele só vai ignorar
	if r := recover(); r != nil {
		fmt.Println("Execucacao recuperado com sucesso")
	}
}

func main() {
	// O retorno é falso pq ele vai retornar o valor 0 do que nós retornamos, no caso falso
	fmt.Println(alunoEstaAprovado(6, 6))
	//Se entrar no panic a linha abaixo nao vai ser executada
	fmt.Println("Pós execução")

}
