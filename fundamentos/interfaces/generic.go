package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Tipo genérico")

	generic("Jilll")
	generic(1)
	generic(true)
}
