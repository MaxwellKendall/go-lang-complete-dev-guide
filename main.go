// Go is not Object Oriented, no "Deck Class" with defined classes that have certain methods and properties
// but this is vaugely like the OO approach, we create a new type, similar to strings, ints, slices, and arrays
// execute go run main.go deck.go
package main

func main() {
	cards := newDeck()

	cards.print()
}
