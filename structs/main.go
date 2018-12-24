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
		firstName: "Max",
		lastName: "Kendall",
		contactInfo: contactInfo{
			email:"checkyoself@checkyoself.com",
			zipcode: 22015,
		},
	}
	fmt.Printf("%+v", max)
}