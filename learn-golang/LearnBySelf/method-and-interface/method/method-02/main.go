package main

import (
	"fmt"
	"math"
)

// Struct Vertex với hai field X, Y
type Vertex struct {
	X, Y float64
}

// Method Abs tính độ dài vector (norm) sử dụng pointer receiver
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Method Scale để thay đổi giá trị của X và Y theo một hệ số
func (v *Vertex) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}

// Method Angle tính góc giữa vector và trục X (theo radian)
func (v *Vertex) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func main() {
	// Khởi tạo giá trị cho Vertex
	v := &Vertex{3, 4}

	// In độ dài vector ban đầu
	fmt.Printf("Initial Vertex: X=%.2f, Y=%.2f\n", v.X, v.Y)
	fmt.Printf("Initial Norm (Abs): %.2f\n", v.Abs())
	fmt.Printf("Initial Angle (Radians): %.2f\n", v.Angle())

	// Thay đổi giá trị X và Y bằng Scale
	scaleFactor := 2.0
	v.Scale(scaleFactor)
	fmt.Printf("\nAfter Scaling by %.2f:\n", scaleFactor)
	fmt.Printf("Updated Vertex: X=%.2f, Y=%.2f\n", v.X, v.Y)
	fmt.Printf("Updated Norm (Abs): %.2f\n", v.Abs())
	fmt.Printf("Updated Angle (Radians): %.2f\n", v.Angle())
}
