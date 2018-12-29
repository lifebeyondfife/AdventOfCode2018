package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Star contains the position and velocity of a star at a given moment in time
type Star struct {
	x  int
	y  int
	dx int
	dy int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseStars(filename string) []Star {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var stars []Star
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x, y, dx, dy int
		starStr := scanner.Text()
		_, err := fmt.Sscanf(starStr, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy)
		check(err)
		stars = append(stars, Star{x, y, dx, dy})
	}
	return stars
}

func getConvergedStars(stars []Star) ([]Star, int) {
	prevMaxMin := math.MaxInt32
	var seconds int
	for {
		xMax, yMax, xMin, yMin := math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
		for index := range stars {
			stars[index].x += stars[index].dx
			stars[index].y += stars[index].dy

			if stars[index].x > xMax {
				xMax = stars[index].x
			}
			if stars[index].x < xMin {
				xMin = stars[index].x
			}
			if stars[index].y > yMax {
				yMax = stars[index].y
			}
			if stars[index].y < yMin {
				yMin = stars[index].y
			}
		}

		if prevMaxMin < (xMax-xMin)+(yMax-yMin) {
			break
		} else {
			prevMaxMin = (xMax - xMin) + (yMax - yMin)
		}

		seconds++
	}

	for index := range stars {
		stars[index].x -= stars[index].dx
		stars[index].y -= stars[index].dy
	}

	return stars, seconds
}

func calculateBounds(stars []Star) (int, int, int, int) {
	xMax, yMax, xMin, yMin := math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for index := range stars {
		if stars[index].x > xMax {
			xMax = stars[index].x
		}
		if stars[index].x < xMin {
			xMin = stars[index].x
		}
		if stars[index].y > yMax {
			yMax = stars[index].y
		}
		if stars[index].y < yMin {
			yMin = stars[index].y
		}
	}
	return xMax - xMin + 13, yMax - yMin + 9, 6 - xMin, 4 - yMin
}

func starsToString(stars []Star) string {
	width, height, xOffset, yOffset := calculateBounds(stars)
	grid := make([][]rune, height)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = '.'
		}
	}

	for _, star := range stars {
		if star.x+xOffset >= 0 && star.x+xOffset < len(grid[0]) &&
			star.y+yOffset >= 0 && star.y+yOffset < len(grid) {
			grid[star.y+yOffset][star.x+xOffset] = '#'
		}
	}

	var gridStr string
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			gridStr += string(grid[i][j])
		}
		gridStr += "\n"
	}

	return gridStr
}

func main() {
	stars := parseStars("./input10.txt")
	convergedStars, seconds := getConvergedStars(stars)

	fmt.Printf("part 1: \n%s\n", starsToString(convergedStars))
	fmt.Printf("part 2: %d\n", seconds)
}
