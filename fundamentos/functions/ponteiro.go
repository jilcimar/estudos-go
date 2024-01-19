package main

import "fmt"

func changeNumber(number int) int {
	return number * -1
}

func changeNumberPointer(number *int) {
	*number = *number * -1
}

func main() {
	fmt.Println("Função com ponteiro")
	result := changeNumber(20)
	fmt.Println(result)

	number := 10
	fmt.Println(number)
	changeNumberPointer(&number)

	fmt.Println("Novo numero", number)
}
