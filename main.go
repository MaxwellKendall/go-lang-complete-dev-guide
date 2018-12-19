// Go is not Object Oriented, no "Deck Class" with defined classes that have certain methods and properties
// instead, we create a new type, similar to strings, ints, slices, and arrays
// execute go run main.go deck.go
package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}

	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}