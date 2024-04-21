package main

import (
	"fmt"
	"testing"
)

func TestChannelWaitWhenEmpty(t *testing.T) {
	// channels by default are pass by reference, even when passing to a function
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	// for range loop will wait until channel is closed, so it'll be blocked here
	for value := range ch {
		fmt.Println("Received:", value)
	}

	// this will wait until channel is closed
	fmt.Println("Waitin indefinitely...")
}
