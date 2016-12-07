package main

import "fmt"

const numColumns, numRows, dead, alive = 10, 10, 0, 1

var grid [][]int

func main() {
	initializeGrid()
	displayGrid()
}

func initializeGrid() {
	// Allocate the top level slice representing the grid
	grid = make([][]int, numRows)

	// Allocate one large slice to hold all of the cell states
	cellStates := make([]int, numColumns*numRows)

	// Initialize all cells to the dead state
	for i := range cellStates {
		cellStates[i] = dead
	}

	// Partition the large slice for use by the top level grid slice
	for i := range grid {
		grid[i], cellStates = cellStates[:numColumns], cellStates[numColumns:]
	}
}

func displayGrid() {
	for i := range grid {
		fmt.Println(grid[i])
	}
}
