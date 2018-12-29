package main

import (
	"testing"
)

func TestPlayGame(t *testing.T) {
	testPlayerCount1 := 10
	testMarbleCount1 := 1618
	ans1 := 8317

	testPlayerCount2 := 13
	testMarbleCount2 := 7999
	ans2 := 146373

	testPlayerCount3 := 17
	testMarbleCount3 := 1104
	ans3 := 2764

	testPlayerCount4 := 21
	testMarbleCount4 := 6111
	ans4 := 54718

	testPlayerCount5 := 30
	testMarbleCount5 := 5807
	ans5 := 37305

	tests := []struct {
		testPlayerCount int
		testMarbleCount int
		answer          int
	}{
		{testPlayerCount1, testMarbleCount1, ans1},
		{testPlayerCount2, testMarbleCount2, ans2},
		{testPlayerCount3, testMarbleCount3, ans3},
		{testPlayerCount4, testMarbleCount4, ans4},
		{testPlayerCount5, testMarbleCount5, ans5},
	}

	for _, test := range tests {
		if playGame(test.testPlayerCount, test.testMarbleCount) != test.answer {
			t.Errorf("Error with play game answer %d, expected %d",
				playGame(test.testPlayerCount, test.testMarbleCount), test.answer)
		}
	}
}
