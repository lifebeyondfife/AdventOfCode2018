package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checksum(ids []string) int {
	var twoCount, threeCount int

	for _, s := range ids {
		counts := make(map[rune]int)

		for _, c := range s {
			counts[c]++
		}

		var twoSet, threeSet bool
		for _, v := range counts {
			if v == 2 {
				twoSet = true
			} else if v == 3 {
				threeSet = true
			}
		}

		if twoSet {
			twoCount++
		}
		if threeSet {
			threeCount++
		}
	}

	return twoCount * threeCount
}

func commonCharacters(ids []string) string {
	for i := 0; i < len(ids[0]); i++ {
		var shortIds []string
		for _, id := range ids {
			shortIds = append(shortIds, id[:i]+id[i+1:])
		}

		sort.Strings(shortIds)

		for i := 1; i < len(shortIds); i++ {
			if shortIds[i-1] == shortIds[i] {
				return shortIds[i]
			}
		}
	}

	return ""
}

func parseIds(file *os.File) []string {
	var ids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	return ids
}

func main() {
	file, err := os.Open("./input02.txt")
	check(err)
	defer file.Close()

	ids := parseIds(file)

	fmt.Printf("puzzle 1: %d\n", checksum(ids))
	fmt.Printf("puzzle 2: %s\n", commonCharacters(ids))
}
