package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		waitingTime := 100 * time.Millisecond
		fmt.Printf("waiting %s for %s\n", s, waitingTime)
		time.Sleep(waitingTime)
		fmt.Println(s)
	}
}

func main() {
	go say("World")
	say("Hello")
}
