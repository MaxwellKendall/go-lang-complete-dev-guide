package main

import (
	"fmt"
)

// Create a new type of deck which is a slice of string

type deck []string

func (d deck) print() {
	// d here is like 'this'
		for i, card := range d {
		// every iteration, we create a new i and card variable
		fmt.Println(i, card)
	}
}