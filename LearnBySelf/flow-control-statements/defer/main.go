package main

import "fmt"

func main() {
	/*
		start type 1
	*/
	// fmt.Println("start")
	// defer fmt.Println("Deferred 1")
	// defer fmt.Println("Deferred 2")
	// defer fmt.Println("Deferred 3")
	// fmt.Println("end")
	/*
		end type 1
	*/

	// ---------------------------------------

	/*
		start type 2
	*/
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	// fmt.Println("Start")
	// panic("Something went wrong!")
	// fmt.Println("This line will not be executed")
	/*
		end type 2
	*/

	// ---------------------------------------
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
