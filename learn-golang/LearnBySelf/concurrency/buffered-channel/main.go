package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numberOfURLs    = 100
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {
	queue := make(chan int, 100)

	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	return queue
}

func main() {
	queue := startQueue()
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i), &wg)
	}

	wg.Wait()
	fmt.Printf("\nHad crawl done %d URLs\n", numberOfURLs)
}
