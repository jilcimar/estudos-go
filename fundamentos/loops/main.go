package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//"While"
	for i < 10 {
		fmt.Println("Incrementendo")
		fmt.Println(i)
		//time.Sleep(time.Second)
		i++
	}

	//For
	for j := 0; j < 10; j++ {
		fmt.Println("For: ")
		fmt.Println(j)
	}

	names := [3]string{"jil", "cimar", "fernandes"}

	for indice, name := range names {
		fmt.Println("Indice ", indice)
		fmt.Println("Nome ", name)
	}

	//Loop infinito
	for {
		fmt.Println("Infinito")
		time.Sleep(time.Second)
	}
}
