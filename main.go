package main

func main() {
	// execute go run main.go deck.go
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
