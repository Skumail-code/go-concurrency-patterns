package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup, sem chan struct{}) {
	for job := range jobs {
		// Acquire semaphore (limit critical section)
		sem <- struct{}{}

		fmt.Printf("Worker %d started job %d\n", id, job)

		// Simulate work
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Worker %d finished job %d\n", id, job)

		// Release semaphore
		<-sem

		wg.Done()
	}
}

func main() {
	numJobs := 20
	numWorkers := 5

	jobs := make(chan int, numJobs)
	sem := make(chan struct{}, 2) // only 2 concurrent critical ops

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, &wg, sem)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}
	close(jobs)

	// Wait for all jobs
	wg.Wait()

	fmt.Println("All jobs done")
}
