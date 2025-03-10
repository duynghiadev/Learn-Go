package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	email     string
}

func main() {
	s := Student{
		firstName: "nghia",
		lastName:  "Le",
		email:     "abc@gmail.com",
	}

	fmt.Println(s)
}
