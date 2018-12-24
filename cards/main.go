package main

/*
problem with existing project from section 3, not easy to tell suits etc bc its just strings
*/

func main() {
	// execute go run main.go deck.go
	cards, _ := newDeckFromFile("mycards")
	cards.print()
}
