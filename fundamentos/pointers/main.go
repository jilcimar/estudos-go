package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var firstNumber int = 10
	var secondNumber int = firstNumber

	fmt.Println(firstNumber, secondNumber)

	firstNumber++

	fmt.Println(firstNumber, secondNumber)

	//Ponteiro é uma referência de memória
	var pointerOne int = 100
	var pointerTwo *int

	fmt.Println(pointerOne, pointerTwo)

	pointerTwo = &pointerOne

	fmt.Println(pointerOne, pointerTwo)

	//Desreferenciação
	fmt.Println(pointerOne, *pointerTwo)

	pointerOne = 150

	fmt.Println(pointerOne, *pointerTwo)

}
