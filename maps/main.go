package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red": "#ff000",
		"green": "#745",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for color", color, "is", hex)
	}
}