package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// create channel with datatype string
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Tantowi Putra Agung Setiawan"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestWaitGoRoutine(t *testing.T) {
	var name string

	go func() {
		name = "Tantowi"
	}()

	fmt.Println(name)
}

func TestChannelBetweenGoRoutines(t *testing.T) {
	channel := make(chan string)

	go func() {
		channel <- "Tantowi Putra Agung Setiawan"
		time.Sleep(10 * time.Second)
	}()

	fmt.Println(<-channel)

	close(channel)
}
func TestTriggerChannel(t *testing.T) {

}

func ChannelSender() {

}

func ChannelReceiver() {
	
}