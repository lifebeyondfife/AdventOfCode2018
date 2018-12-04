package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type cut struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseCut(cutStr string) cut {
	var r = regexp.MustCompile(`#([0-9]*) @ ([0-9]*),([0-9]*): ([0-9]*)x([0-9]*)`)

	m := r.FindSubmatch([]byte(cutStr))

	var props []int
	for _, match := range m[1:] {
		propInt, err := strconv.Atoi(string(match))
		check(err)
		props = append(props, propInt)
	}

	return cut{props[0], props[1], props[2], props[3], props[4]}
}

func overlapCount(cutStrs []string) int {
	var cuts []cut
	for _, cutStr := range cutStrs {
		cuts = append(cuts, parseCut(cutStr))
	}

	var overlap int
	var grid [1000][1000][]int

	for _, cut := range cuts {
		for x := 0; x < cut.width; x++ {
			for y := 0; y < cut.height; y++ {
				if len(grid[x+cut.x][y+cut.y]) == 1 {
					overlap++
				}

				grid[x+cut.x][y+cut.y] = append(grid[x+cut.x][y+cut.y], cut.id)
			}
		}
	}

	return overlap
}

func nonOverlapID(cutStrs []string) int {
	var cuts []cut
	for _, cutStr := range cutStrs {
		cuts = append(cuts, parseCut(cutStr))
	}

	nonOverlap := make(map[int]bool)
	for _, cut := range cuts {
		nonOverlap[cut.id] = true
	}

	var grid [1000][1000][]int

	for _, cut := range cuts {
		for x := 0; x < cut.width; x++ {
			for y := 0; y < cut.height; y++ {
				if len(grid[x+cut.x][y+cut.y]) >= 1 {
					for _, id := range grid[x+cut.x][y+cut.y] {
						nonOverlap[id] = false
						nonOverlap[cut.id] = false
					}
				}

				grid[x+cut.x][y+cut.y] = append(grid[x+cut.x][y+cut.y], cut.id)
			}
		}
	}

	for k, v := range nonOverlap {
		if v {
			return k
		}
	}

	return 0
}

func parseCuts(file *os.File) []string {
	var cutStrs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cutStrs = append(cutStrs, scanner.Text())
	}
	return cutStrs
}

func main() {
	file, err := os.Open("./input03.txt")
	check(err)
	defer file.Close()

	cutStrs := parseCuts(file)

	fmt.Printf("puzzle 1: %d\n", overlapCount(cutStrs))
	fmt.Printf("puzzle 2: %d\n", nonOverlapID(cutStrs))
}
