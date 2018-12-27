package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Coordinate contains the (x, y) values of a cartesian coordinate
type Coordinate struct {
	x int
	y int
}

func createGrid(nodes []Coordinate) [][]int {
	xmax, ymax := math.MinInt32, math.MinInt32
	for _, node := range nodes {
		if node.x > xmax {
			xmax = node.x
		}
		if node.y > ymax {
			ymax = node.y
		}
	}

	grid := make([][]int, xmax+1)

	for i := range grid {
		grid[i] = make([]int, ymax+1)
	}

	return grid
}

func abs(n int32) int32 {
	y := n >> 31
	return (n ^ y) - y
}

func manhattanDistance(a Coordinate, b Coordinate) int {
	return int(abs(int32(a.x-b.x)) + abs(int32(a.y-b.y)))
}

func nearestNeighbours(grid [][]int, nodes []Coordinate) {
	for i := range grid {
		for j := range grid[i] {
			var coord = Coordinate{i, j}
			var nearestNode int
			var nextNearestNode int

			for index, node := range nodes {
				if manhattanDistance(coord, node) <= manhattanDistance(coord, nodes[nearestNode]) {
					nextNearestNode = nearestNode
					nearestNode = index
				}
			}

			if manhattanDistance(coord, nodes[nearestNode]) == manhattanDistance(coord, nodes[nextNearestNode]) &&
				nearestNode != nextNearestNode {
				grid[i][j] = -1
			} else {
				grid[i][j] = nearestNode
			}
		}
	}
}

func pointsPerNode(grid [][]int) map[int]int {
	points := make(map[int]int)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == -1 {
				continue
			}

			points[grid[i][j]]++
		}
	}

	return points
}

func removeInfiniteNodes(grid [][]int, scores map[int]int) {
	for i := range grid {
		for j := range grid[i] {
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				delete(scores, grid[i][j])
			}
		}
	}
}

func largestFiniteArea(nodes []Coordinate) int {
	grid := createGrid(nodes)
	nearestNeighbours(grid, nodes)

	scores := pointsPerNode(grid)
	removeInfiniteNodes(grid, scores)

	var highestScore int
	for _, v := range scores {
		if v > highestScore {
			highestScore = v
		}
	}

	return highestScore
}

func findInnerRegionSize(nodes []Coordinate, maxCumDistance int) int {
	grid := createGrid(nodes)

	var innerRegionSize int
	for i := range grid {
		for j := range grid[i] {
			var coord = Coordinate{i, j}
			var cumSize int
			for _, node := range nodes {
				cumSize += manhattanDistance(node, coord)

				if cumSize >= maxCumDistance {
					break
				}
			}

			if cumSize < maxCumDistance {
				innerRegionSize++
			}
		}
	}

	return innerRegionSize
}

func parseCoordinates() []Coordinate {
	file, err := os.Open("./input06.txt")
	check(err)
	defer file.Close()

	var coords []Coordinate
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordStr := scanner.Text()
		var x, y int
		_, err := fmt.Sscanf(coordStr, "%d, %d", &x, &y)
		check(err)
		coords = append(coords, Coordinate{x, y})
	}
	return coords
}

func main() {
	nodes := parseCoordinates()

	fmt.Printf("part 1: %d\n", largestFiniteArea(nodes))
	fmt.Printf("part 2: %d\n", findInnerRegionSize(nodes, 10000))
}
