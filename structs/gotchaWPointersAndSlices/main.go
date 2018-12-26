package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(slice []string) {
	slice[0] = "Bye"
}