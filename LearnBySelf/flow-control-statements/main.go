package main

import "fmt"

func main() {
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
}
