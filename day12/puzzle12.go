package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Pot describes the position of a pot, and whether or not it holds a plant
type Pot struct {
	id   int
	char string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseStateAndRules() (string, map[string]string) {
	file, err := os.Open("./input12.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var state string
	_, err = fmt.Sscanf(scanner.Text(), "initial state: %s", &state)
	check(err)
	scanner.Scan()

	rules := make(map[string]string)
	for scanner.Scan() {
		var key, value string
		_, err = fmt.Sscanf(scanner.Text(), "%s => %s", &key, &value)
		check(err)
		rules[key] = value
	}
	return state, rules
}

func potsToStr(pots []Pot) string {
	var potBuffer bytes.Buffer

	for _, pot := range pots {
		potBuffer.WriteString(pot.char)
	}

	return potBuffer.String()
}

func potsToStrTrimmed(pots []Pot) string {
	return strings.Trim(potsToStr(pots), ".")
}

func iterateState(pots []Pot, generations int, rules map[string]string) []Pot {
	for gen := 0; gen < generations; gen++ {
		nextGen := make([]Pot, len(pots))
		for index, pot := range pots {
			nextGen[index] = Pot{id: pot.id, char: pot.char}
		}

		for i := 2; i < len(pots)-2; i++ {
			potStr := potsToStr(pots[i-2 : i+3])
			if rules[potStr] == "#" {
				nextGen[i].char = "#"
			} else {
				nextGen[i].char = "."
			}
		}

		pots = padEmptyPots(nextGen)
	}
	return pots
}

func padEmptyPots(pots []Pot) []Pot {
	const pad int = 3
	var earliestPotValue int
	var earliestPotIndex int

	for index, pot := range pots {
		if pot.char == "#" {
			earliestPotValue = pot.id
			earliestPotIndex = index
			break
		}
	}

	var emptyPots []Pot
	for i := 0; i < pad-earliestPotIndex; i++ {
		emptyPots = append(emptyPots, Pot{char: ".", id: earliestPotValue - pad + i})
	}

	pots = append(emptyPots, pots...)

	var latestPotValue int
	var latestPotIndex int

	for i := len(pots) - 1; i >= 0; i-- {
		if pots[i].char == "#" {
			latestPotValue = pots[i].id
			latestPotIndex = i
			break
		}
	}

	emptyPots = nil
	for i := len(pots) - latestPotIndex; i <= pad; i++ {
		emptyPots = append(emptyPots, Pot{char: ".", id: latestPotValue + i})
	}

	pots = append(pots, emptyPots...)

	return pots
}

func createPots(stateStr string) []Pot {
	var pots []Pot

	for index, char := range stateStr {
		pots = append(pots, Pot{index, string(char)})
	}

	return padEmptyPots(pots)
}

func totalScore(pots []Pot) int {
	var score int
	for _, pot := range pots {
		if pot.char == "#" {
			score += pot.id
		}
	}
	return score
}

func findStableState(pots []Pot, rules map[string]string) (int, int) {
	previousDelta := 0
	previousScore := totalScore(pots)
	var generations int

	for {
		generations++
		pots = iterateState(pots, 1, rules)
		score := totalScore(pots)
		if score-previousScore == previousDelta && previousDelta > 0 {
			break
		}
		previousDelta = score - previousScore
		previousScore = score
	}

	return generations, previousDelta
}

func main() {
	stateStr, rules := parseStateAndRules()
	state := createPots(stateStr)
	generations, delta := findStableState(state, rules)

	fmt.Printf("part 1: %d\n", totalScore(iterateState(state, 20, rules)))
	fmt.Printf("part 2: %d\n", totalScore(iterateState(state, generations, rules))+
		(50000000000-generations)*delta)
}
