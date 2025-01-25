package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func Example4() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch2 <- 2
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
