package main

import (
	"fmt"
	grid "learn-by-self/more-type/slices_of_slices/common"
)

func main() {
	// Initialize the grid
	myGrid := grid.CreateGrid()

	fmt.Println("Original Grid:")
	grid.PrintGrid(myGrid)

	// Modify the grid
	myGrid = grid.ModifyGrid(myGrid)

	// Add a new row
	newRow := []int{10, 11, 12}
	myGrid = grid.AddRow(myGrid, newRow)

	fmt.Println("\nUpdated Grid:")
	grid.PrintGrid(myGrid)
}
