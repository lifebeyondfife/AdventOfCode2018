package main

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	test1 := []int{1}
	ans1 := 0

	tests := []struct {
		testCase []int
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if template(test.testCase) != test.answer {
			t.Errorf("Error with answer %d, expected %d", template(test.testCase), test.answer)
		}
	}
}
