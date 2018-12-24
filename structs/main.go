package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	max := person{
		firstName: "Maxwell",
		lastName: "Kendall",
		contactInfo: contactInfo{
			email:"checkyoself@checkyoself.com",
			zipcode: 22015,
		},
	}

	pointerToMax := &max
	pointerToMax.updateName("Maxie Boy")
	max.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)	
}