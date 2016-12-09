package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const numColumns, numRows = 20, 20
const dead, alive = false, true

var grid [][]bool
var changedCells []cellCoordinates

type cellCoordinates struct {
	x int
	y int
}

func main() {

	// Clear the terminal to make way for the grid
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	initializeGrid()

	grid[8][9] = true
	grid[9][9] = true
	grid[10][9] = true
	grid[10][10] = true
	grid[7][11] = true

	displayGrid()

	// Game loop
	for {
		time.Sleep(500 * time.Millisecond)
		tick()
		for _, cell := range changedCells {
			grid[cell.x][cell.y] = !grid[cell.x][cell.y]
		}
		displayGrid()
	}
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
	fmt.Print("\033[0;0H")
	for y := numRows - 1; y >= 0; y = y - 1 {
		for x := 0; x < numColumns; x = x + 1 {
			if grid[x][y] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func tick() {
	changedCells = nil

	for x, columns := range grid {
		for y, cellState := range columns {
			if cellState != willBeAlive(x, y) {
				changedCells = append(changedCells, cellCoordinates{x, y})
			}
		}
	}
}

func willBeAlive(x int, y int) bool {
	numLiveNeighbors := getNumLiveNeighbors(x, y)
	if grid[x][y] {
		return numLiveNeighbors == 2 || numLiveNeighbors == 3
	}
	return numLiveNeighbors == 3
}

func getNumLiveNeighbors(x int, y int) int {
	neighbors := 0
	for xx := x - 1; xx <= x+1; xx = xx + 1 {
		for yy := y - 1; yy <= y+1; yy = yy + 1 {
			if !(xx == x && yy == y) && inBounds(xx, yy) && grid[xx][yy] {
				neighbors = neighbors + 1
			}
		}
	}
	return neighbors
}

func inBounds(x int, y int) bool {
	return x >= 0 && x < numColumns && y >= 0 && y < numRows
}
