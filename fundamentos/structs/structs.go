package main

import "fmt"

type user struct {
	name string
	age  int
	address
}

type address struct {
	number  int
	zipcode string
}

func main() {
	fmt.Println("Arquivo de structs")

	var customer user
	customer.name = "Jilcimar"
	customer.age = 26
	fmt.Println(customer)

	professional := user{"Jilcimar Fernandes", 27, address{511, "59945000"}}
	fmt.Println(professional)
	fmt.Println(professional.zipcode)

	person := user{name: "João Fernandes"}
	fmt.Println(person)
}
