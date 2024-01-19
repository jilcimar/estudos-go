package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência != Paralelismo

	//Paralelismo = 2 ou mais tarefas executando ao mesmo tempo
	//Concorrência = Não necessariamente executando ao mesmo tempo..

	fmt.Println("Concorrência")

	go write("Olá mundo!") //goroutine
	write("Programando em GoGoGoooo!")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
