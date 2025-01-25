package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func Example7() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 1 done.")
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		fmt.Println("Goroutine 2 done.")
	}()

	time.Sleep(time.Second * 6)
}
