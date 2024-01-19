package main

import "fmt"

func main() {
	fmt.Println("Switchs")

	fmt.Println(daysWeek(1))
}

func daysWeek(day int) string {
	switch day {
	case 1:
		fallthrough // joga seu código para a próxima condição
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terça feira"
	case 4:
		return "Quarta feira"
	case 5:
		return "Quinta feira"
	case 6:
		return "Sexta feira"
	default:
		return "Sábado"
	}
}
