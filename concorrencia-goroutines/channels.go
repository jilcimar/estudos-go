package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência != Paralelismo

	//Paralelismo = 2 ou mais tarefas executando ao mesmo tempo

	channel := make(chan string)

	go record("Olá mundo", channel)

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa")
}

func record(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
}
