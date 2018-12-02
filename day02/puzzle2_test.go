package main

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	test1 := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	ans1 := 12

	tests := []struct {
		testCase []string
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if checksum(test.testCase) != test.answer {
			t.Errorf("Error with checksum %d (expected %d) for list of strings %s",
				checksum(test.testCase), test.answer, test.testCase)
		}
	}
}

func TestCommonCharacters(t *testing.T) {
	test1 := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	ans1 := "fgij"

	tests := []struct {
		testCase []string
		answer   string
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if commonCharacters(test.testCase) != test.answer {
			t.Errorf("Error with common characters %s (expected %s) for list of strings %s",
				commonCharacters(test.testCase), test.answer, test.testCase)
		}
	}
}
