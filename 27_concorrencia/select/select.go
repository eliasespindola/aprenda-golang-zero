package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			//Vai esperar 2 segundos
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//Se ele estiver pronto para receber a mensagem ele vai receber, ao inves de ter o delay
		select {
		case msg := <-canal1:
			fmt.Println(msg)
		case msg := <-canal2:
			fmt.Println(msg)
		}
	}
}
