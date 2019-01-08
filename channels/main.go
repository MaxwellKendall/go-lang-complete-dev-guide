package main

import (
	"fmt"
	"net/http"
)

func main() {
	// this is the main go routine
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// go creates a new routine
		go checkLink(link, c)
	}

	fmt.Println(<- c) // receive the value in the channel and print it
}

func checkLink(website string, c chan string) {
	// this is a blocking call
	_, err := http.Get(website)
	if err != nil {
		fmt.Println("Error:", website, "might be down")
		c <- "Might be down...?"
		return
	}
	fmt.Println(website, "is good!")
	c <- "Appears to be good!"
}
