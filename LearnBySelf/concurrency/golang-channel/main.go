package main

import (
	"fmt"
	"golang-channel-example/examples"
)

func main() {
	fmt.Println("Choose an example to run:")
	fmt.Println("1: Example1")
	fmt.Println("2: Example2")
	fmt.Println("3: Example3")
	fmt.Println("4: Example4")
	fmt.Println("5: Example5")
	fmt.Println("6: Example6")
	fmt.Println("7: Example7")
	fmt.Println("8: Example8")
	fmt.Println("9: Example9")
	fmt.Println("10: Example10")

	var choice int
	fmt.Printf("Your choice: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		examples.Example1()
	case 2:
		examples.Example2()
	case 3:
		examples.Example3()
	case 4:
		examples.Example4()
	case 5:
		examples.Example5()
	case 6:
		examples.Example6()
	case 7:
		examples.Example7()
	case 8:
		examples.Example8()
	case 9:
		examples.Example9()
	case 10:
		examples.Example10()
	default:
		fmt.Println("Invalid choice.")
	}
}
