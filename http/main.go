package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// io.Copy accepts 2 params: (1) implements Writer if, (2) Implements Reader if
	// os.Stdout is of type file, which implements Writer, which therefore satisfies (1)
	// resp.Body implements the Read() fn which satisfies (2)
	io.Copy(os.Stdout, resp.Body)
}