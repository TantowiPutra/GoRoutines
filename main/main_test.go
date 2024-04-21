package main

import (
	"fmt"
	"testing"
)

func TestChannelWaitWhenEmpty(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Waitin indefinitely...")
}
