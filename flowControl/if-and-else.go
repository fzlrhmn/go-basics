package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	// variables v only available in if-else scope
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v below this line
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
