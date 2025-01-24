package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		fmt.Println("value:", v)
		sum += v
	}
	c <- sum // gửi sum vào channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	// In ra các phần của mảng để debug
	fmt.Println("Nửa đầu mảng:", s[:len(s)/2])
	fmt.Println("Nửa sau mảng:", s[len(s)/2:])

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // nhận giá trị từ channel

	fmt.Println("Tổng phần 1:", x)
	fmt.Println("Tổng phần 2:", y)
	fmt.Println("Tổng cuối cùng:", x+y)
}
