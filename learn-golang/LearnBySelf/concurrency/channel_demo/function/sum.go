package function

import (
	"fmt"
)

// sum calculates the sum of elements in a slice and sends the result to a channel
func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("value:", v)
		sum += v
	}
	c <- sum // Send the result to the channel
}

// AddProduct demonstrates the use of channels with the sum function
func AddProduct() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// Debug: Print array parts
	fmt.Println("Nửa đầu mảng:", s[:len(s)/2])
	fmt.Println("Nửa sau mảng:", s[len(s)/2:])

	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)

	x, y := <-c, <-c // Receive values from the channel

	fmt.Println("Tổng phần 1:", x)
	fmt.Println("Tổng phần 2:", y)
	fmt.Println("Tổng cuối cùng:", x+y)
}
