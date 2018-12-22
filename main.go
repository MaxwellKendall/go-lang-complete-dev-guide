package main

// Usinge the go standard library we often need the type 'byte slice' which is a computer friendly way of representing a string
// See this table: http://www.asciitable.com/ (Bytes are referred to here as decimals in the table)

func main() {
	// execute go run main.go deck.go
	cards := newDeck()
	cards.saveToFile("mycards")
}
