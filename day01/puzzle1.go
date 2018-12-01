package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findLoop(freqs []int) int {
	seen := make(map[int]bool)

	var tot int
	for {
		for _, freq := range freqs {
			tot += freq

			if seen[tot] {
				return tot

			}

			seen[tot] = true
		}
	}
}

func main() {
	file, err := os.Open("./input01_1.txt")
	check(err)
	defer file.Close()

	var freqs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		freq, err := strconv.Atoi(scanner.Text())
		check(err)
		freqs = append(freqs, freq)
	}

	var tot int
	for _, freq := range freqs {
		tot += freq
	}

	fmt.Printf("puzzle 1: %d\n", tot)

	fmt.Printf("puzzle 2: %d\n", findLoop(freqs))
}
