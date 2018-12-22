package main

func main() {
	// execute go run main.go deck.go
	// returns an error: cards, _ := newDeckFromFile("yo")
	cards, _ := newDeckFromFile("mycards")
	cards.print()
}
