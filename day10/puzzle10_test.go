package main

import (
	"testing"
)

func TestParseStars(t *testing.T) {
	test1 := "./input10.txt"
	ans1 := Star{-9767, 50146, 1, -5}

	tests := []struct {
		testCase string
		answer   Star
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if parseStars(test.testCase)[0] != test.answer {
			t.Errorf("Error with first parsed Star %#v, expected %#v", parseStars(test.testCase)[0], test.answer)
		}
	}
}

func TestTemplate(t *testing.T) {
	test1 := []Star{
		Star{9, 1, 0, 2},
		Star{7, 0, -1, 0},
		Star{3, -2, -1, 1},
		Star{6, 10, -2, -1},
		Star{2, -4, 2, 2},
		Star{-6, 10, 2, -2},
		Star{1, 8, 1, -1},
		Star{1, 7, 1, 0},
		Star{-3, 11, 1, -2},
		Star{7, 6, -1, -1},
		Star{-2, 3, 1, 0},
		Star{-4, 3, 2, 0},
		Star{10, -3, -1, 1},
		Star{5, 11, 1, -2},
		Star{4, 7, 0, -1},
		Star{8, -2, 0, 1},
		Star{15, 0, -2, 0},
		Star{1, 6, 1, 0},
		Star{8, 9, 0, -1},
		Star{3, 3, -1, 1},
		Star{0, 5, 0, -1},
		Star{-2, 2, 2, 0},
		Star{5, -2, 1, 2},
		Star{1, 4, 2, 1},
		Star{-2, 7, 2, -2},
		Star{3, 6, -1, -1},
		Star{5, 0, 1, 0},
		Star{-6, 0, 2, 0},
		Star{5, 9, 1, -2},
		Star{14, 7, -2, 0},
		Star{-3, 6, 2, -1},
	}
	ansMessage1 := `......................
......................
......................
......................
......#...#..###......
......#...#...#.......
......#...#...#.......
......#####...#.......
......#...#...#.......
......#...#...#.......
......#...#...#.......
......#...#..###......
......................
......................
......................
......................
`
	ansSeconds1 := 3

	tests := []struct {
		testCase      []Star
		answerMessage string
		answerSeconds int
	}{
		{test1, ansMessage1, ansSeconds1},
	}

	for _, test := range tests {
		convergedStars, seconds := getConvergedStars(test.testCase)
		if starsToString(convergedStars) != test.answerMessage {
			t.Errorf("Error with answer message %s, expected %s", starsToString(convergedStars), test.answerMessage)
		}
		if seconds != test.answerSeconds {
			t.Errorf("Error with answer seconds %d, expected %d", seconds, test.answerSeconds)
		}
	}
}
