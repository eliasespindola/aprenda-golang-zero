package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Ola mundo")
		//Vai tirar um do contador ficando 1 do waitgroup
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done()
	}()
	//Vai esperar o contador ficar em 0
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
