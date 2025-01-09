package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Abs implements Abser.
func (v *Vertex) Abs() float64 { // Use pointer receiver
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat implements Abser
	a = &v // *Vertex implements Abser

	// Uncommenting the following line will now cause a compilation error:
	// a = v // Vertex (not *Vertex) does NOT implement Abser

	fmt.Println(a.Abs()) // This will call the Abs() method of *Vertex
}
