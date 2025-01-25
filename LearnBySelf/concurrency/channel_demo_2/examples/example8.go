package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Example8() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	wc := new(sync.WaitGroup)
	wc.Add(2)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 1 done.")
		wc.Done()
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 2 done.")
		wc.Done()
	}()

	wc.Wait()
	fmt.Println("All Goroutines done")
}
