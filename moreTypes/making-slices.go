package main

import (
	"fmt"
)

func printSliceAgain(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d val=%v\n", s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	printSliceAgain("a", a)

	b := make([]int, 0, 5)
	printSliceAgain("b", b)

	c := b[:2]
	printSliceAgain("c", c)

	d := c[2:5]
	printSliceAgain("d", d)
}
