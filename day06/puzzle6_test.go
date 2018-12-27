package main

import (
	"testing"
)

func TestParseCoordinates(t *testing.T) {
	ans1 := Coordinate{45, 315}

	tests := []struct {
		answer Coordinate
	}{
		{ans1},
	}

	for _, test := range tests {
		if parseCoordinates()[0] != test.answer {
			t.Errorf("Error with parsing coordinates answer %d, expected %d", parseCoordinates()[0], test.answer)
		}
	}
}

func TestLargestFiniteArea(t *testing.T) {
	test1 := []Coordinate{
		Coordinate{1, 1},
		Coordinate{1, 6},
		Coordinate{8, 3},
		Coordinate{3, 4},
		Coordinate{5, 5},
		Coordinate{8, 9},
	}
	ans1 := 17

	tests := []struct {
		testCase []Coordinate
		answer   int
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if largestFiniteArea(test.testCase) != test.answer {
			t.Errorf("Error with largest finite area answer %d, expected %d",
				largestFiniteArea(test.testCase), test.answer)
		}
	}
}

func TestFindInnerRegionSize(t *testing.T) {
	testGrid1 := []Coordinate{
		Coordinate{1, 1},
		Coordinate{1, 6},
		Coordinate{8, 3},
		Coordinate{3, 4},
		Coordinate{5, 5},
		Coordinate{8, 9},
	}
	testMaxCumSize1 := 32
	ans1 := 16

	tests := []struct {
		testGrid       []Coordinate
		testMaxCumSize int
		answer         int
	}{
		{testGrid1, testMaxCumSize1, ans1},
	}

	for _, test := range tests {
		if findInnerRegionSize(test.testGrid, test.testMaxCumSize) != test.answer {
			t.Errorf("Error with inner region size answer %d, expected %d",
				findInnerRegionSize(test.testGrid, test.testMaxCumSize), test.answer)
		}
	}
}
