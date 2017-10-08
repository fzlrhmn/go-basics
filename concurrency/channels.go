package main

import (
	"time"
)

func tick(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func mains() {
	for i := 0; i < 24; i++ {
		c := tick(100 * time.Millisecond)
		<-c
	}
}
