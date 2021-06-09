package main

import (
	"fmt"
	"applied-go/structs_practise"
)

func main() {
	person := Person{
		Name: "Harrison",
		Age: 28,
		Address: Address{
			Name: "Orrong Road",
			Suburb: "Toorak",
		},
	}
	fmt.Println(person)
}


