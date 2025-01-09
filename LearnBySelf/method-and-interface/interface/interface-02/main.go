package main

import (
	"fmt"
	"math"
)

// Abser interface
type Abser interface {
	Abs() float64
}

// MyFloat type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex struct
type Vertex struct {
	X, Y float64
}

// Abs implements Abser for *Vertex
func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// I interface
type I interface {
	M()
}

// T struct
type T struct {
	S string
}

// M method implements I for *T
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// Example 1: Demonstrating Abser interface
func example1() {
	fmt.Println("Example 1: Abser Interface")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat implements Abser
	fmt.Printf("MyFloat: %v\n", a.Abs())

	a = &v // *Vertex implements Abser
	fmt.Printf("*Vertex: %v\n", a.Abs())

	// Uncommenting the following line will cause a compilation error:
	// a = v // Vertex (not *Vertex) does NOT implement Abser
}

// Example 2: Demonstrating I interface with nil values
func example2() {
	fmt.Println("\nExample 2: I Interface with Nil Values")
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

// describe function
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// main function
func main() {
	example1()
	example2()
}
