package main

import (
	"fmt"
)

type shape interface{
	getArea() float64
}
type triangle struct{
	length float64
	width float64
}
type square struct{
	length float64
	width float64
}

func main() {
	triangle := triangle{width: float64(10), length: float64(6)}
	square := square{width: float64(5), length: float64(30)}
	
	printArea(triangle)
	printArea(square)
}

func (t triangle) getArea() float64 {
	return t.length * t.width / 2
}
func (s square) getArea() float64 {
	return s.length * s.width
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}