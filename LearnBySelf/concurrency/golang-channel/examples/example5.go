package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func Example5() {
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

	select {
	case v1 := <-ch1:
		fmt.Println("Ch1 comes first with value:", v1)
		fmt.Println("Then ch2 with value:", <-ch2)
	case v2 := <-ch2:
		fmt.Println("Ch2 comes first with value:", v2)
		fmt.Println("Then ch1 with value:", <-ch1)
	}
}
