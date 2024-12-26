package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0 // Initial guess
	for i := 0; i < 10; i++ {
		change := (z*z - x) / (2 * z)
		newZ := z - change
		fmt.Printf("Iteration %d: z = %f, change = %f, formula: %f - (%f*%f - %f) / (2*%f) = %f\n",
			i+1, z, change, z, z, z, x, z, newZ)
		z = newZ
	}
	return z
}

func SqrtWithConvergence(x float64) float64 {
	z := 1.0 // Initial guess
	threshold := 1e-9
	for {
		change := (z*z - x) / (2 * z)
		newZ := z - change
		fmt.Printf("Current z = %f, change = %f, formula: %f - (%f*%f - %f) / (2*%f) = %f\n",
			z, change, z, z, z, x, z, newZ)
		if math.Abs(newZ-z) < threshold {
			break
		}
		z = newZ
	}
	return z
}

func main() {
	x := 2.0

	// Fixed iterations
	fmt.Println("Sqrt with fixed iterations:")
	fmt.Printf("Approximation: %f\n\n", Sqrt(x))

	// With convergence
	fmt.Println("Sqrt with convergence:")
	fmt.Printf("Approximation: %f\n\n", SqrtWithConvergence(x))

	// Compare with math.Sqrt
	fmt.Printf("math.Sqrt result: %f\n", math.Sqrt(x))
}
