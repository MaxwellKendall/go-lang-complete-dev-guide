package main

import (
	"fmt"
)

// different ways to declare a struct

type person struct {
	firstName string
	lsatName string
}

func main() {
	// relying on the order of definition
	max := person{"Max", "Kendall"}
	// explicitly specifying the properties and giving them values
	claire := person{firstName: "Claire", lsatName: "Kendall"}

	fmt.Println(max, claire)
}