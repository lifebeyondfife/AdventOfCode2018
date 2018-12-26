package main

import (
	"testing"
)

func TestFileParsing(t *testing.T) {
	ans1 := "YvVIiMhHcw"

	tests := []struct {
		answer string
	}{
		{ans1},
	}

	for _, test := range tests {
		if parseUnits()[:10] != test.answer {
			t.Errorf("Error with parsing file: first 10 characters %s, expected %s",
				parseUnits()[:10], test.answer)
		}
	}
}
func TestExecuteChainReactions(t *testing.T) {
	test1 := "dabAcCaCBAcCcaDA"
	ans1 := len("dabCBAcaDA")

	tests := []struct {
		testCase string
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if executeChainReactions(test.testCase) != test.answer {
			t.Errorf("Error with executing chain reactions answer %d, expected %d",
				executeChainReactions(test.testCase), test.answer)
		}
	}
}

func TestFindProblemPolymer(t *testing.T) {
	test1 := "dabAcCaCBAcCcaDA"
	ans1 := len("daDA")

	tests := []struct {
		testCase string
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if findProblemPolymer(test.testCase) != test.answer {
			t.Errorf("Error with finding problematic polymer answer %d, expected %d",
				findProblemPolymer(test.testCase), test.answer)
		}
	}
}
