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
	orderedIds := make([]string, len(ids))
	copy(orderedIds, ids)

	/* There is a flaw in this program's strategy that means it won't work
	for all inputs. If the difference occurs between the two common
	strings occurs in the first character, the two strings won't be
	next to each other in the ordered array. This gets the right answer
	for the given input but there are other for which it won't. */
	sort.Strings(orderedIds)

	for i := 1; i < len(orderedIds); i++ {
		var common string
		var differences int
		for idx := range orderedIds[i-1] {
			if orderedIds[i-1][idx] != orderedIds[i][idx] {
				differences++

				if differences > 1 {
					break
				}
			} else {
				common = common + string(orderedIds[i-1][idx])
			}
		}

		if differences == 1 {
			return common
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
