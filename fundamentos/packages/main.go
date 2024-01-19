package main

import (
	"estudos/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Aula de Pacotes")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("jilcimaremail.com.br")
	fmt.Println(erro)
}
