package main

import "fmt"

func main() {
	//Função anônima
	func() {
		fmt.Println("Olá mundo")
	}()

	retorno := func(value string) string {
		return fmt.Sprintf("Recebido -> %s + %d", value, 10)
	}("Passando o parametro")

	fmt.Println(retorno)
}
