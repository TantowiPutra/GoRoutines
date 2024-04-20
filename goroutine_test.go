package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type TestStruct struct {
	name int
}

func RunHelloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go RunHelloWorld(&wg)
	wg.Wait()

	fmt.Println("Oops")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

// Goroutine is super lightweight compared to Thread, consumes only 2kb of memory as a comparison, thread consumes 1000kb
// In advance, go routine also runs concurrently
func TestManyWithoutGoRoutine(t *testing.T) {
	for i := 0; i < 10000 ; i++ {
		DisplayNumber(i)
	}
}

func TestManyGoRoutines(t *testing.T) {
	for i := 0; i < 10000 ; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}







