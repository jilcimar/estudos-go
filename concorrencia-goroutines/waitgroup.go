package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Concorrência != Paralelismo

	//Paralelismo = 2 ou mais tarefas executando ao mesmo tempo

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //Passando que tem 2 GoRoutines no grupo de espera

	go func() {
		escrever("Olá mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
