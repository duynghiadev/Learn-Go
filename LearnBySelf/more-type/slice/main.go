package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("v has X before change = ", v)
	v.X = 4
	fmt.Println("v has X after change = ", v)
	fmt.Println("v.X = ", v.X)
}
