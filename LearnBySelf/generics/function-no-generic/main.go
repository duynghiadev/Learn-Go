package main

import "fmt"

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	rs := mapInt(arr, func(v int) int {
		fmt.Println("v:", v)
		return v * 2
	})
	fmt.Println("result:", rs)
}
