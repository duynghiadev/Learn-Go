package main

import (
	"fmt"
	"math"
)

// --- Lesson 1: Abser Interface ---
type Abser interface {
	Abs() float64
}

// MyFloat type implementing Abser
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex struct implementing Abser
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// --- Lesson 2: Interface with Methods (Nil Handling) ---
type I interface {
	M()
}

// T struct implementing I
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// --- Lesson 3: The Empty Interface ---
func emptyInterfaceExample() {
	fmt.Println("\nLesson 3: The Empty Interface")
	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)

	// Demonstrating type assertion
	if val, ok := i.(string); ok {
		fmt.Printf("Type assertion successful: %q is a string\n", val)
	} else {
		fmt.Println("Type assertion failed")
	}
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// --- Helper Functions ---
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// --- Lesson 1 Example ---
func abserExample() {
	fmt.Println("Lesson 1: Abser Interface")
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // MyFloat implements Abser
	fmt.Printf("MyFloat: %v\n", a.Abs())

	a = &v // *Vertex implements Abser
	fmt.Printf("*Vertex: %v\n", a.Abs())
}

// --- Lesson 2 Example ---
func interfaceWithMethodsExample() {
	fmt.Println("\nLesson 2: Interface with Methods (Nil Handling)")
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

// --- Main Function ---
func main() {
	// Run each lesson example
	abserExample()
	interfaceWithMethodsExample()
	emptyInterfaceExample()
}
