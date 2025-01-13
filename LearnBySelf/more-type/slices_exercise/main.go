package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

// Pic generates a slice of slices of uint8
func Pic(dx, dy int) [][]uint8 {
	// Create the outer slice
	picture := make([][]uint8, dy)

	// Populate the outer slice
	for y := 0; y < dy; y++ {
		// Create the inner slice
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Generate the value using a formula
			row[x] = uint8((x + y) / 2) // Change this formula for different patterns
		}
		// Assign the row to the outer slice
		picture[y] = row
	}

	return picture
}

func main() {
	// Call the Pic function to generate and display the picture
	pic.Show(Pic)

	fmt.Printf("\n--------------\n")

	// Tạo một slice chứa các slice
	grid := make([][]uint8, 3) // Tạo 3 hàng

	for i := range grid {
		grid[i] = make([]uint8, 5) // Mỗi hàng có 5 cột
		for j := range grid[i] {
			grid[i][j] = uint8(i + j) // Giá trị trong bảng là i + j
		}
	}

	// In bảng
	for _, row := range grid {
		fmt.Println(row)
	}
}
