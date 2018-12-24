package main

import (
	"fmt"
)

// different ways to declare a struct

type person struct {
	firstName string
	lastName string
}

func main() {
	// relying on the order of definition
	max := person{"Max", "Kendall"}
	// explicitly specifying the properties and giving them values
	claire := person{firstName: "Claire", lastName: "Kendall"}
	// init empty struct with 0 values
	var ross person
	// updating new person using . syntax
	claire.lastName = "SQRL"

	fmt.Println(max, claire, ross)
	fmt.Printf("%+v", ross) // shows all the fields!
}