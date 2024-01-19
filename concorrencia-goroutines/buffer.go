package main

import "fmt"

func main() {
	//Concorrência != Paralelismo

	//Paralelismo = 2 ou mais tarefas executando ao mesmo tempo

	channel := make(chan string, 2) //Especificando a capacidade do canal
	channel <- "Olá mundo"
	channel <- "Olá mundo"

	mensagem := <-channel

	fmt.Println(mensagem)
}
