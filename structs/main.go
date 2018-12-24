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
	contactInfo // does the same thing
	// contact contactInfo
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
	// demonstrates that GO is pass by value, not pass by reference!
	max.updateName("Max")
	max.print()
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)	
}