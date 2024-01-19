package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculate(first, second int) (int, string) {
	sum := first + second
	description := "soma"

	return sum, description
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var imprimir = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := imprimir("imprimindo texto")
	fmt.Println(resultado)

	fmt.Println("==== dois retornos =======")
	fmt.Println(calculate(2, 4))
}
