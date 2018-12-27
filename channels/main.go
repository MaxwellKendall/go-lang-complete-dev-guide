package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(website string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println("Error:", website, "might be down")
		return
	}
	fmt.Println(website, "is good!")
}