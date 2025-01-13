package grid

import "fmt"

// CreateGrid initializes a 2D grid (slide of slices) with 3 rows and 3 columns
func CreateGrid() [][]int {
	return [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}

// ModifyGrid updates specific element in the grid
func ModifyGrid(grid [][]int) [][]int {
	grid[0][0] = 99
	grid[1][1] = 88
	return grid
}

// AddRow adds a new row to the grid
func AddRow(grid [][]int, newRow []int) [][]int {
	return append(grid, newRow)
}

// PrintGrid prints the 2D grid
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println("row:", row)
	}
}
