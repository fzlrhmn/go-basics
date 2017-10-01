package main

import (
	"fmt"
)

// Square is struct for define sides width of a box
type Square struct {
	side float64
}

type Shape interface {
	area() float64
}

func (z Square) area() float64 {
	side := z.side
	return side * side
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{
		side: 10,
	}

	info(s)
}
