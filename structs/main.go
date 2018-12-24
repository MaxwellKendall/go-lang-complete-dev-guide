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
	contact contactInfo
}

func main() {
	max := person{
		firstName: "Max",
		lastName: "Kendall",
		contact: contactInfo{
			email:"checkyoself@checkyoself.com",
			zipcode: 22015,
		},
	}
	fmt.Printf("%+v", max)
}