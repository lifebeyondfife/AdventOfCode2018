package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateChainTerms() []string {
	var chainTerms []string

	for char := 'a'; char <= 'z'; char++ {
		chainTerms = append(chainTerms, string(char)+string(unicode.ToUpper(char)),
			string(unicode.ToUpper(char))+string(char))
	}

	return chainTerms
}

func findProblemPolymer(units string) int {
	var reactionLengths []int

	for char := 'a'; char <= 'z'; char++ {
		var unitsRemoved string

		for _, unit := range units {
			if unit == char || unit == unicode.ToUpper(char) {
				continue
			}

			unitsRemoved = unitsRemoved + string(unit)
		}

		reactionLengths = append(reactionLengths, executeChainReactions(unitsRemoved))
	}

	sort.Ints(reactionLengths)
	return reactionLengths[0]
}

func executeChainReactions(units string) int {
	chainTerms := generateChainTerms()

	for {
		var foundReaction bool
		for _, chainTerm := range chainTerms {
			for {
				index := strings.Index(units, chainTerm)

				if index == -1 {
					break
				}

				foundReaction = true

				units = units[:index] + units[index+2:]
			}
		}

		if !foundReaction {
			break
		}
	}

	return len(units)
}

func parseUnits() string {
	file, err := os.Open("./input05.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	units := parseUnits()

	fmt.Printf("part 1: %d\n", executeChainReactions(units))
	fmt.Printf("part 2: %d\n", findProblemPolymer(units))
}
