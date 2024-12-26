package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
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

	/*
		start Switch case
	*/
	fmt.Println("------------start Switch case---------------------")
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Print("----------------------------------------\n")

	fmt.Print("When is Saturday? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Print("----------------------------------------\n")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 20:
		fmt.Println("Good evening.")
	}
	/*
		end Switch case
	*/
}
