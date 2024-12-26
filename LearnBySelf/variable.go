package main

import (
	"fmt"
	"log"
)

func main() {
	var value interface{} = []byte("Hello, Go!")

	bytes, ok := value.([]byte)

	if ok {
		// fmt.Println("Type assertion successful:", string(bytes))     // log type 1
		log.Printf("Type assertion successful: %s\n", bytes) // log type 2
	} else {
		fmt.Println("Type assertion failed.", bytes)
	}
}
