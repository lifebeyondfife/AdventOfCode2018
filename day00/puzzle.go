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

func template(_ []int) int {
	return 0
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
	file, err := os.Open("./input01.txt")
	check(err)
	defer file.Close()

	fmt.Printf("part 1: \n")
	fmt.Printf("part 2: \n")
}
