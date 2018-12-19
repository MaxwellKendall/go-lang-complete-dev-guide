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

// no receiver, because caller presumably doesn't have a deck yet
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
// returning multiple values!
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}