package main

import "fmt"

func modString(a *string) {
	fmt.Println("Inside:", a)
	*a = "modified"
}

func main() {
	a := "hello"
	fmt.Println("variable a in memory:", &a)

	fmt.Println("Before:", a)
	modString(&a)
	fmt.Println("After:", a)
}
