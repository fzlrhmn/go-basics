package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func timesFour(v Vertex) float64 {
	return ((v.X * 4) + (v.Y * 4))
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(), timesFour(v))
}
