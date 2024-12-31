# **Slice Operations and Indexing in Go**

## **1. Slice Indexing in Go**

In Go, the elements of a **slice** (or **array**) are indexed starting from `0`. This means:

- The **first element** of the slice is at **index 0**.
- The **second element** is at **index 1**.
- And so on.

### **Example**

Given the following slice:

```go
operations := []Operation{
	{"Hypotenuse", func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}},

	{"Power", math.Pow},

	{"Addition", func(x, y float64) float64 {
		return x + y
	}},

	{"Multiplication", func(x, y float64) float64 {
		return x * y
	}},
}
```

The elements are indexed as follows:

| **Index** | **Element**               | **Name**           |
| --------- | ------------------------- | ------------------ |
| `0`       | `{"Hypotenuse", ...}`     | `"Hypotenuse"`     |
| `1`       | `{"Power", math.Pow}`     | `"Power"`          |
| `2`       | `{"Addition", ...}`       | `"Addition"`       |
| `3`       | `{"Multiplication", ...}` | `"Multiplication"` |

### **Accessing Elements**

You can access elements of the slice using their index:

```go
fmt.Println(operations[0].Name) // Output: "Hypotenuse"
fmt.Println(operations[1].Name) // Output: "Power"
fmt.Println(operations[2].Name) // Output: "Addition"
fmt.Println(operations[3].Name) // Output: "Multiplication"
```

### **Handling User Input**

If a user selects an operation based on a 1-based input (e.g., `1` for the first operation), you need to adjust for the 0-based indexing by subtracting `1`:

```go
selectedOp := operations[choice-1]
```

For example:

- If the user enters `1`, `choice-1` becomes `0`, which corresponds to the first element in the slice.

---

## **2. Using `len()` in Go**

The built-in function `len()` in Go is used to determine the **length** of a slice, array, string, or map. The length is **1-based**, meaning it represents the total count of elements, not their indices.

### **Example**

For the `operations` slice:

```go
fmt.Println(len(operations)) // Output: 4
```

### **Key Difference Between `len()` and Indexing**

- **`len()`** returns the **total number of elements** in the slice (starts from `1`).
- **Indexing** starts from `0` and goes up to `len(slice) - 1`.

### **Example in Practice**

To ensure the user's input is valid:

```go
if choice < 1 || choice > len(operations) {
	fmt.Println("Invalid choice.")
	return
}
```

This checks if the user’s input is within the valid range of `1` to `len(operations)`.

---

## **3. Summary**

- **Indexing** in Go starts from `0`.
- **`len()`** provides the total number of elements in a slice, starting from `1`.
- Always adjust for 1-based user inputs when working with 0-based indexing.

### **Code Example**

Here’s a full example demonstrating the concepts:

```go
package main

import (
	"fmt"
	"math"
)

type Operation struct {
	Name string
	Func func(float64, float64) float64
}

func main() {
	operations := []Operation{
		{"Hypotenuse", func(x, y float64) float64 {
			return math.Sqrt(x*x + y*y)
		}},
		{"Power", math.Pow},
		{"Addition", func(x, y float64) float64 {
			return x + y
		}},
		{"Multiplication", func(x, y float64) float64 {
			return x * y
		}},
	}

	fmt.Println("Available Operations:")
	for i, op := range operations {
		fmt.Printf("%d. %s\n", i+1, op.Name)
	}

	var choice int
	fmt.Print("Select an operation (1-4): ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(operations) {
		fmt.Println("Invalid choice.")
		return
	}

	var x, y float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	selectedOp := operations[choice-1]
	result := selectedOp.Func(x, y)

	fmt.Printf("Result of %s(%f, %f): %f\n", selectedOp.Name, x, y, result)
}
```
