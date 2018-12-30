package main

import (
	"fmt"
	"math"
)

// Coord holds a cartesian coordinate
type Coord struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculatePowerLevel(x, y, gridSerial int) int {
	rackID := x + 10
	powerLevel := rackID * y
	powerLevel += gridSerial
	powerLevel *= rackID
	powerLevel = (powerLevel / 100) % 10
	powerLevel -= 5
	return powerLevel
}

func generateGrid(gridSerial int) [][]int {
	grid := make([][]int, 300)

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 300)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = calculatePowerLevel(i, j, gridSerial)
		}
	}

	return grid
}

func largestCellSquare(grid [][]int, gridSerial, square int) (Coord, int) {
	maxCellPosition := Coord{0, 0}
	maxCell := math.MinInt32

	for i := 0; i < len(grid)-square+1; i++ {
		for j := 0; j < len(grid[i])-square+1; j++ {
			var cell int
			for a := 0; a < square; a++ {
				for b := 0; b < square; b++ {
					cell += grid[i+a][j+b]
				}
			}
			if cell > maxCell {
				maxCell = cell
				maxCellPosition = Coord{i, j}
			}
		}
	}

	return maxCellPosition, maxCell
}

func largestCell(gridSerial int) (Coord, int) {
	grid := generateGrid(gridSerial)
	maxCellSize := 1
	maxCell := math.MinInt32
	maxCellPosition := Coord{0, 0}

	// Using an upper bound of 20 is a heuristic (guess). I cannot put this as high as 300
	// because I didn't use the Summed Area Table technique to avoid recomputation of
	// larger grid sizes - see https://en.wikipedia.org/wiki/Summed-area_table
	for i := 1; i <= 20; i++ {
		coord, cell := largestCellSquare(grid, gridSerial, i)

		if cell > maxCell {
			maxCellSize = i
			maxCellPosition = coord
			maxCell = cell
		}
	}

	return maxCellPosition, maxCellSize
}

func main() {
	gridSerial := 5093
	cellPosition3x3, _ := largestCellSquare(generateGrid(gridSerial), gridSerial, 3)
	cellPosition, size := largestCell(gridSerial)

	fmt.Printf("part 1: %v\n", cellPosition3x3)
	fmt.Printf("part 2: %v, %d\n", cellPosition, size)
}
