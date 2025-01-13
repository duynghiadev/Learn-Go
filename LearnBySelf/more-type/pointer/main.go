package main

import "fmt"

func addElement(slice *[]int, elements ...int) {
	*slice = append(*slice, elements...)
}

func main() {
	mySlice := []int{}

	addElement(&mySlice, 1, 2, 3, 4, 5)
	fmt.Println("My Slice:", mySlice)
}
