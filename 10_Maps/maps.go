package main

import "fmt"

func main() {
	fmt.Println("Maps")
	//[Chave], valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedeo",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	//Removendo chave
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	//Adicionando chave
	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}

	fmt.Println(usuario2["signo"]["nome"])
}
