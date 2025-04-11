package main

import (
	"fmt"
	"math"
)

// Operation struct chứa tên và logic thực hiện của phép toán
type Operation struct {
	Name string
	Func func(float64, float64) float64
}

// compute áp dụng hàm toán học với 2 tham số x và y
func compute(fn func(float64, float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func main() {
	// Danh sách các phép toán
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

	// Hiển thị danh sách phép toán
	fmt.Println("Available Operations:")
	for i, op := range operations {
		fmt.Printf("%d. %s\n", i+1, op.Name)
	}

	// Người dùng chọn phép toán
	var choice int
	fmt.Print("Select an operation (1-4): ")
	fmt.Scan(&choice)

	// Kiểm tra lựa chọn hợp lệ
	if choice < 1 || choice > len(operations) {
		fmt.Println("Invalid choice.")
		return
	}

	// Nhập các giá trị x và y
	var x, y float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	// Thực hiện phép toán
	selectedOp := operations[choice-1]
	result := compute(selectedOp.Func, x, y)

	// Hiển thị kết quả
	fmt.Printf("Result of %s(%f, %f): %f\n", selectedOp.Name, x, y, result)
}
