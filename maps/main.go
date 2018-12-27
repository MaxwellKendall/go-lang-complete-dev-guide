package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)

	// var colors map[string]string
	// colors := map[string]string{
	// 	"red": "#ff000",
	// 	"green": "#745",
	// }

	colors["white"] = "#ffff"
	// dot syntax (colors.white = "XYZ") does not work because key has to be of a certain type
	delete(colors, "white")

	fmt.Println(colors)
}