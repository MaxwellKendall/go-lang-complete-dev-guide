package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // only strings are legit!
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	// actually creates a new slice!
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}