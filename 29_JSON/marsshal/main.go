package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	// sTRUCT OU MAP PARA JSON
	cachorroEmJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))
	//Json para struct ou map
	//json.Unmarshal()
}
