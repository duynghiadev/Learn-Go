package main

import "fmt"

func main() {
	// fmt.Println("start")
	// defer fmt.Println("Deferred 1")
	// defer fmt.Println("Deferred 2")
	// defer fmt.Println("Deferred 3")
	// fmt.Println("end")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start")
	panic("Something went wrong!")
	fmt.Println("This line will not be executed")
}
