package main

import (
	"fmt"
	"math"
)

func pow(m, n, lim float64) float64 {
	if v := math.Pow(m, n); v < lim {
		return v
	}
	return lim
}

// Square: bình phương của một số
func checkEvenSquare(num, threshold int) int {
	if square := num * num; square <= threshold {
		if square%2 == 0 {
			fmt.Printf("Square of %d is %d and is event.\n", num, square)
		} else {
			fmt.Printf("Square of %d is %d and is odd.\n", num, square)
		}
		return square
	}
	fmt.Printf("Square of %d exceeds the threshold (%d).\n", num, threshold)
	return -1
}

func main() {
	fmt.Println("------------start For loop---------------------")
	sum := 0
	cat := 0
	j := 0
	tom := 0
	black := 0
	pink := 0
	countInfinite := 0
	resultInfinite := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	for ; j < 5; j++ {
		cat += j
	}

	for k := 0; k < 5; {
		tom += k
		k++
	}

	for pink < 10 { // no init, no post -> the same while loop
		black += 1
		pink++
	}

	for { // no init, no condition, no post
		resultInfinite += countInfinite
		countInfinite++
		if countInfinite > 10 {
			break
		}
	}

	fmt.Println("sum:", sum)
	fmt.Println("cat:", cat)
	fmt.Println("tom:", tom)
	fmt.Println("black:", black)
	fmt.Println("resultInfinite:", resultInfinite)
	fmt.Println("------------end For loop---------------------")
	/*
		end For loop
	*/

	fmt.Println("result pow:", pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println("checkEvenSquare:", checkEvenSquare(2, 10))
	fmt.Println("checkEvenSquare:", checkEvenSquare(3, 10))
	fmt.Println("checkEvenSquare:", checkEvenSquare(5, 20))
}
