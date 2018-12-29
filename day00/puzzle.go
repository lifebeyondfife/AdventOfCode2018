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

func parse() []int {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

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
	fmt.Printf("part 1: \n")
	fmt.Printf("part 2: \n")
}
