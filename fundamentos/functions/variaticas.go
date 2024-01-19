package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	fmt.Println(numbers)

	return total
}

func main() {
	fmt.Println("Funções variáticas")

	total := sum(1, 2, 5)

	fmt.Println(total)

}
