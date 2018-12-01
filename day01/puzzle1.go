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

func findTotal(freqs []int) int {
	var tot int
	for _, freq := range freqs {
		tot += freq
	}
	return tot
}

func findLoop(freqs []int) int {
	seen := make(map[int]bool)
	seen[0] = true

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

func parseFreqs(file *os.File) []int {
	var freqs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		freq, err := strconv.Atoi(scanner.Text())
		check(err)
		freqs = append(freqs, freq)
	}
	return freqs
}

func main() {
	file, err := os.Open("./input01_1.txt")
	check(err)
	defer file.Close()

	freqs := parseFreqs(file)

	fmt.Printf("puzzle 1: %d\n", findTotal(freqs))

	fmt.Printf("puzzle 2: %d\n", findLoop(freqs))
}
