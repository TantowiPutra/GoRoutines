package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker Began Working...", id)
	fmt.Println("Job Assigned..", len(jobs))
	for job := range jobs {
		fmt.Println("Job Assigned Inner..", len(jobs))
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Simulate doing some work
		result := job * 2
		// Send the result back to the results channel
		results <- result
		fmt.Printf("Worker %d finished job %d with result %d\n", id, job, result)
	}
}

func main() {
	// Create channels for jobs and results
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
		fmt.Println("Assigned Worker...", i)
	}

	// Send some jobs to the workers
	for i := 1; i <= 10; i++ {
		jobs <- i
		fmt.Println("Assigned Jobs...", i)
		time.Sleep(2 * time.Second)
	}

	close(jobs) // no more jobs send, close channel

	// Wait for all workers to finish
	wg.Wait()

	// Close the results channel after all workers have finished
	close(results)

	// Collect results from the results channel
	for result := range results {
		fmt.Println("Result:", result)
	}
}
