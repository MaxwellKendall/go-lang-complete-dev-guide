package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	io.Copy(os.Stdout, file)	
}