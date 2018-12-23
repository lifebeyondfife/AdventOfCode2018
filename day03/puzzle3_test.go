package main

import (
	"testing"
)

func TestOverlapCount(t *testing.T) {
	test1 := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	ans1 := 4

	tests := []struct {
		testCase []string
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if overlapCount(test.testCase) != test.answer {
			t.Errorf("Error with common characters %d (expected %d) for list of strings %s",
				overlapCount(test.testCase), test.answer, test.testCase)
		}
	}
}

func TestNonOverlap(t *testing.T) {
	test1 := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	ans1 := 3

	tests := []struct {
		testCase []string
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if nonOverlapID(test.testCase) != test.answer {
			t.Errorf("Error with finding non-overlapping cut %d (expected %d) for list of strings %s",
				nonOverlapID(test.testCase), test.answer, test.testCase)
		}
	}
}
