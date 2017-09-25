package main

import (
	"fmt"
)

// same as 'while' in C
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
