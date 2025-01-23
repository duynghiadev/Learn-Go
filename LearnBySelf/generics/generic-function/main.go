package main

import (
	"fmt"
	"strings"
)

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	rs := mapAny(arr, func(v int) int {
		fmt.Println("v:", v)
		return v * 2
	})

	strArr := []string{"apple", "orange"}
	rsStr := mapAny(strArr, func(v string) string {
		return strings.ToUpper(v)
	})

	fmt.Println("result:", rs)
	fmt.Println("result string:", rsStr)
}
