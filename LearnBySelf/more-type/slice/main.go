package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// --- Structs Field Example ---
	fmt.Println("\n--- Structs Field Example ---")
	v := Vertex{1, 2}
	fmt.Println("v has X before change = ", v)
	v.X = 4
	fmt.Println("v has X after change = ", v)
	fmt.Println("v.X = ", v.X)

	// --- Structs Field Pointer Example ---
	fmt.Println("\n--- Structs Field Pointer Example ---")
	p := &v
	p.X = 1e9
	fmt.Println("v has X after change via pointer = ", v)

	// --- Structs Field Pointer Example 2 ---
	fmt.Println("\n--- Structs Field Pointer Example 2 ---")
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p1 := &Vertex{3, 4}
	fmt.Println("v1 = ", v1)
	fmt.Println("v2 = ", v2)
	fmt.Println("v3 = ", v3)
	fmt.Println("p1 = ", p1)

	// --- Structs and Slice Example ---
	fmt.Println("\n--- Structs and Slice Example ---")
	v4 := []Vertex{
		{1, 2},
		{3, 4},
	}
	fmt.Println("v4 = ", v4)

	// --- Structs, Slice, Array, Map, Function Example ---
	fmt.Println("\n--- Structs, Slice, Array, Map, Function Example ---")
	var a [2]Vertex
	a[0] = Vertex{1, 2}
	a[1] = Vertex{3, 4}
	fmt.Println("a = ", a)

	// --- Structs, Slice, Array, Map, Function more advanced example ---

}
