package main

import "fmt"

const numColumns, numRows = 10, 10
const dead, alive = false, true

var grid [][]bool

func main() {
	initializeGrid()

	grid[5][5] = true
	grid[4][4] = true
	grid[6][6] = true

	displayGrid()
}

func initializeGrid() {
	// Allocate the top level slice representing the grid
	grid = make([][]bool, numRows)

	// Allocate one large slice to hold all of the cell states
	cellStates := make([]bool, numColumns*numRows)

	// Partition the large slice for use by the top level grid slice
	for i := range grid {
		grid[i], cellStates = cellStates[:numColumns], cellStates[numColumns:]
	}
}

func displayGrid() {
	for _, columns := range grid {
		for _, cellState := range columns {
			if cellState {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
