package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Salvando o usuário %s no banco de dados\n", u.name)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	fmt.Println("Métodos")

	customer := user{"Jilcimar", 26}

	fmt.Println(customer)

	customer.save()

	age := customer.ofLegalAge()

	println(age)

	customer.birthday()

	fmt.Println(customer.age)
}
