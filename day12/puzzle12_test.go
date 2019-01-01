package main

import (
	"testing"
)

func TestParsePotsAndRules(t *testing.T) {
	testKey1 := "##..."
	ansState1 := "##..##..#.##.###....###.###.#.#.######.#.#.#.#.##.###.####..#.###...#######.####.##...#######.##..#"
	ans1 := "."

	tests := []struct {
		testKey     string
		answerState string
		answer      string
	}{
		{testKey1, ansState1, ans1},
	}

	for _, test := range tests {
		state, rules := parseStateAndRules()
		if state != test.answerState {
			t.Errorf("Error with state %s, expected %s", state, test.answerState)
		}
		if rules[test.testKey] != test.answer {
			t.Errorf("Error with rule for %s, expected %s", test.testKey, test.answer)
		}
	}
}

func TestGeneration(t *testing.T) {
	testState := createPots("#..#.#..##......###...###")
	testRules := map[string]string{
		"...##": "#",
		"..#..": "#",
		".#...": "#",
		".#.#.": "#",
		".#.##": "#",
		".##..": "#",
		".####": "#",
		"#.#.#": "#",
		"#.###": "#",
		"##.#.": "#",
		"##.##": "#",
		"###..": "#",
		"###.#": "#",
		"####.": "#",
	}

	ansGeneration1 := "#...#....#.....#..#..#..#"
	ansGeneration10 := "#.#..#...#.##....##..##..##..##"

	tests := []struct {
		testState      []Pot
		testRules      map[string]string
		testGeneration int
		answer         string
	}{
		{testState, testRules, 1, ansGeneration1},
		{testState, testRules, 10, ansGeneration10},
	}

	for _, test := range tests {
		pots := iterateState(test.testState, test.testGeneration, test.testRules)
		if potsToStrTrimmed(pots) != test.answer {
			t.Errorf("Error with pots generation %d: %s, expected %s",
				test.testGeneration, potsToStrTrimmed(pots), test.answer)
		}
	}
}

func TestTotalScore(t *testing.T) {
	testState := createPots("#..#.#..##......###...###")
	testRules := map[string]string{
		"...##": "#",
		"..#..": "#",
		".#...": "#",
		".#.#.": "#",
		".#.##": "#",
		".##..": "#",
		".####": "#",
		"#.#.#": "#",
		"#.###": "#",
		"##.#.": "#",
		"##.##": "#",
		"###..": "#",
		"###.#": "#",
		"####.": "#",
	}
	ansScore := 325

	tests := []struct {
		testState []Pot
		testRules map[string]string
		answer    int
	}{
		{testState, testRules, ansScore},
	}

	for _, test := range tests {
		score := totalScore(iterateState(test.testState, 20, test.testRules))
		if score != test.answer {
			t.Errorf("Error with pots score after 20 generations %d, expected %d", score, test.answer)
		}
	}
}
