package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var numbers [5]int

	numbers[1] = 2

	fmt.Println(numbers)

	names := [3]string{"Jil", "Jilcimar", "Cimar"}

	fmt.Println(names)

	slice := []int{1, 3, 5, 66, 8}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(names))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := names[0:2]

	fmt.Println(slice2)
}
