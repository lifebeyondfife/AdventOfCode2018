package main

import (
	"testing"
)

func TestParseNumbers(t *testing.T) {
	ans1 := []int{9, 11, 7}

	tests := []struct {
		answer []int
	}{
		{ans1},
	}

	for _, test := range tests {
		numbers := parseNumbers()
		if numbers[0] != test.answer[0] || numbers[1] != test.answer[1] || numbers[2] != test.answer[2] {
			t.Errorf("Error with parsed numbers %d, expected %d", numbers[:3], test.answer)
		}
	}
}

func TestTreeScore(t *testing.T) {
	test1 := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	ans1 := 138

	tests := []struct {
		testCase []int
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if treeScore(test.testCase) != test.answer {
			t.Errorf("Error with tree score %d, expected %d", treeScore(test.testCase), test.answer)
		}
	}
}

func TestNodeValue(t *testing.T) {
	test1 := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	ans1 := 66

	tests := []struct {
		testCase []int
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if nodeValue(test.testCase) != test.answer {
			t.Errorf("Error with tree score %d, expected %d", nodeValue(test.testCase), test.answer)
		}
	}
}
