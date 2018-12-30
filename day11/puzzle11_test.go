package main

import (
	"testing"
)

type testCase struct {
	x          int
	y          int
	gridSerial int
}

func TestCalculatePowerLevel(t *testing.T) {
	test1 := testCase{3, 5, 8}
	ans1 := 4

	test2 := testCase{122, 79, 57}
	ans2 := -5

	test3 := testCase{217, 196, 39}
	ans3 := 0

	test4 := testCase{101, 153, 71}
	ans4 := 4

	tests := []struct {
		testCase testCase
		answer   int
	}{
		{test1, ans1},
		{test2, ans2},
		{test3, ans3},
		{test4, ans4},
	}

	for _, test := range tests {
		if calculatePowerLevel(test.testCase.x, test.testCase.y, test.testCase.gridSerial) != test.answer {
			t.Errorf("Error with power level answer %d, expected %d",
				calculatePowerLevel(test.testCase.x, test.testCase.y, test.testCase.gridSerial), test.answer)
		}
	}
}

func TestLargestCell3x3(t *testing.T) {
	testGridSerial1 := 18
	testSquare1 := 3
	ans1 := Coord{33, 45}

	tests := []struct {
		testGridSerial int
		testSquare     int
		answer         Coord
	}{
		{testGridSerial1, testSquare1, ans1},
	}

	for _, test := range tests {
		coord, _ := largestCellSquare(generateGrid(test.testGridSerial), test.testGridSerial, test.testSquare)
		if coord != test.answer {
			t.Errorf("Error with largest cell 3x3 answer %d, expected %d", coord, test.answer)
		}
	}
}

func TestLargestCell(t *testing.T) {
	testGridSerial1 := 18
	ansPosition1 := Coord{90, 269}
	ansSize1 := 16

	tests := []struct {
		testGridSerial int
		answerPosition Coord
		answerSize     int
	}{
		{testGridSerial1, ansPosition1, ansSize1},
	}

	for _, test := range tests {
		coord, size := largestCell(test.testGridSerial)
		if coord != test.answerPosition || size != test.answerSize {
			t.Errorf("Error with largest cell position %v, expected %v", coord, test.answerPosition)
			t.Errorf("Error with largest cell size %d, expected %d", coord, test.answerSize)
		}
	}
}
