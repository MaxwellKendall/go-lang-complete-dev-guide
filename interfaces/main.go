// Interfaces will help us reuse the printGreeting accross multiple types

package main

import (
	"fmt"
)

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

type spanishBot struct{}
type englishBot struct{}

func (englishBot) getGreeting() string {
	// very custom logic for english greeting
	return "Hello there"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }