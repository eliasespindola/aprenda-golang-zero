package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	//Temos as tags em json que se a gente deixar o `json:"-"` assim nao vai deserializar o campo
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c cachorro
	err := json.Unmarshal([]byte(cachorroEmJSON), &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

}
