// accessing slices:

// slices are indexed from 0
// acessing a slice like so slice[0]
// selecting a range of items in a slice: `slice[index1:index2]`
//  - index1 is inclusive, index2 is exclusive

// `slice[:index2]` items up to and excluding index2
// `slice[index2:]` items begining at and including index2, to the end 

package main

func main() {
	// execute go run main.go deck.go
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
