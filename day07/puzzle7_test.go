package main

import (
	"testing"
)

func TestParseDependencies(t *testing.T) {
	test1 := "E"
	ans1 := "V"

	test2 := "R"
	ans2 := "L"

	tests := []struct {
		testCase string
		answer   string
	}{
		{test1, ans1},
		{test2, ans2},
	}

	for _, test := range tests {

		if parseDependencies()[test.testCase][0] != test.answer {
			t.Errorf("Error with answer parsing dependencies %s, expected %s",
				parseDependencies()[test.testCase][0], test.answer)
		}
	}
}

func TestAssemblyOrder(t *testing.T) {
	test1 := map[string][]string{
		"A": []string{"C"},
		"B": []string{"A"},
		"D": []string{"A"},
		"E": []string{"B", "D", "F"},
		"F": []string{"C"},
	}
	ans1 := "CABDFE"

	tests := []struct {
		testCase map[string][]string
		answer   string
	}{
		{test1, ans1},
	}

	for _, test := range tests {
		if assemblyOrder(test.testCase) != test.answer {
			t.Errorf("Error with answer %s, expected %s", assemblyOrder(test.testCase), test.answer)
		}
	}
}

func TestMultipleWorkersAssemblyOrder(t *testing.T) {
	testParts1 := map[string][]string{
		"A": []string{"C"},
		"B": []string{"A"},
		"D": []string{"A"},
		"E": []string{"B", "D", "F"},
		"F": []string{"C"},
	}
	testWorkerCount1 := 2
	testTimePerPart1 := 0
	ans1 := 15

	tests := []struct {
		testParts       map[string][]string
		testWorkerCount int
		testTimePerPart int
		answer          int
	}{
		{testParts1, testWorkerCount1, testTimePerPart1, ans1},
	}

	for _, test := range tests {
		if multipleWorkersAssemblyOrder(test.testParts, test.testWorkerCount, test.testTimePerPart) != test.answer {
			t.Errorf("Error with answer %d, expected %d",
				multipleWorkersAssemblyOrder(test.testParts, test.testWorkerCount, test.testTimePerPart), test.answer)
		}
	}
}
