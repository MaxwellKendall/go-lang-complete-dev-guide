// ONLY PRINTS ONE WEBSITE GO TO NEXT BRANCH TO SEE SOLUTION

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
	for l := range c {
		// when this channel gets a value, assign it to l and execute for loop body
		go checkLink(l, c)
	}
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
	c <- website
}
