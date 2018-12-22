package main

import (
	"fmt"
)

// Usinge the go standard library we often need the type 'byte slice' which is a computer friendly way of representing a string
// See this table: http://www.asciitable.com/ (Bytes are referred to here as decimals in the table)

func main() {
	// execute go run main.go deck.go
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	definitelyAString := "This is most definitely a string"
	fmt.Println([]byte(definitelyAString))
}
