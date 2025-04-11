package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for range make([]struct{}, 5) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
