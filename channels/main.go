// ONLY PRINTS ONE WEBSITE GO TO NEXT BRANCH TO SEE SOLUTION

package main

import (
	"time"
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
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
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
