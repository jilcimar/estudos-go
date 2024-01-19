package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name": "Jilcimar",
		"age":  "26",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	delete(user, "age")

	fmt.Println(user)

	age := 26

	if age < 15 {
		fmt.Println("Menor que 15")
	} else {
		fmt.Println("Maior que 15")
	}
}
