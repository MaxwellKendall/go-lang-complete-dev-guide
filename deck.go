package main

import (
	"os"
	"io/ioutil" // ioutil is a subpackage of io pkg!
	"strings"
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

// coercing deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	bytes := []byte(d.toString())
	return ioutil.WriteFile(filename, bytes, 0666) //anyone can read/write this file
}

func newDeckFromFile(filename string) (deck, error) {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Error Handling Options:
			// #1. Log error, return new deck via call to newDeck()
			// #2. Log error, entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
		return deck{}, err
	}
	
	s := strings.Split(string(byteSlice), ",")
	return deck(s), nil
}