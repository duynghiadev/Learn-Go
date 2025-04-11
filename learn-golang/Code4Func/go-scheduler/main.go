package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID      int
	Execute func() error
}

// Worker represents a goroutine that executes tasks
type Worker struct {
	ID      int
	taskCh  chan Task
	quit    chan bool
	wg      *sync.WaitGroup
	errChan chan error
}

// NewWorker creates a new worker
func NewWorker(id int, taskCh chan Task, wg *sync.WaitGroup, errChan chan error) *Worker {
	return &Worker{
		ID:      id,
		taskCh:  taskCh,
		quit:    make(chan bool),
		wg:      wg,
		errChan: errChan,
	}
}

// Start begins the worker's processing loop
func (w *Worker) Start(ctx context.Context) {
	go func() {
		defer w.wg.Done()
		for {
			select {
			case task, ok := <-w.taskCh:
				if !ok {
					return
				}
				log.Printf("Worker %d processing task %d\n", w.ID, task.ID)
				if err := task.Execute(); err != nil {
					w.errChan <- fmt.Errorf("worker %d failed task %d: %v", w.ID, task.ID, err)
				}
			case <-ctx.Done():
				log.Printf("Worker %d shutting down...\n", w.ID)
				return
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop signals the worker to stop processing
func (w *Worker) Stop() {
	w.quit <- true
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
	workers  []*Worker
	taskCh   chan Task
	wg       sync.WaitGroup
	errChan  chan error
	numCPU   int
	maxProcs int
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int) *WorkerPool {
	if numWorkers <= 0 {
		numWorkers = runtime.NumCPU()
	}
	return &WorkerPool{
		taskCh:   make(chan Task, numWorkers),
		errChan:  make(chan error, numWorkers),
		numCPU:   runtime.NumCPU(),
		maxProcs: numWorkers,
	}
}

// Start initializes and starts the worker pool
func (wp *WorkerPool) Start(ctx context.Context) {
	runtime.GOMAXPROCS(wp.maxProcs)
	log.Printf("Starting worker pool with %d workers (System has %d CPUs)\n", wp.maxProcs, wp.numCPU)

	// Initialize workers
	for i := 0; i < wp.maxProcs; i++ {
		wp.wg.Add(1)
		worker := NewWorker(i+1, wp.taskCh, &wp.wg, wp.errChan)
		wp.workers = append(wp.workers, worker)
		worker.Start(ctx)
	}
}

// Submit adds a new task to the pool
func (wp *WorkerPool) Submit(task Task) {
	wp.taskCh <- task
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.taskCh)
	wp.wg.Wait()
}

// Example task functions
func task1() error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing task 1")
	return nil
}

func task2() error {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Executing task 2")
	return nil
}

func task3() error {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Executing task 3")
	return nil
}

func task4() error {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Executing task 4")
	return nil
}

func main() {
	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		log.Println("Received interrupt signal. Shutting down...")
		cancel()
	}()

	// Create and start worker pool
	pool := NewWorkerPool(4)
	pool.Start(ctx)

	// Create error handling goroutine
	go func() {
		for err := range pool.errChan {
			log.Printf("Error: %v\n", err)
		}
	}()

	// Submit tasks
	tasks := []Task{
		{ID: 1, Execute: task1},
		{ID: 2, Execute: task2},
		{ID: 3, Execute: task3},
		{ID: 4, Execute: task4},
	}

	for _, task := range tasks {
		pool.Submit(task)
	}

	// Wait for completion and cleanup
	pool.Stop()
	close(pool.errChan)
}
