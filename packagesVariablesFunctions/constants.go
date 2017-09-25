package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Printf("Pi is %T\n", Pi)

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
