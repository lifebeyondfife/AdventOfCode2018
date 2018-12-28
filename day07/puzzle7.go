package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type worker struct {
	nextFree int
	part     string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func allDone(done map[string]bool) bool {
	for _, v := range done {
		if !v {
			return false
		}
	}

	return true
}

func allDependenciesDone(done map[string]bool, deps []string) bool {
	for _, dep := range deps {
		if !done[dep] {
			return false
		}
	}

	return true
}

func generateDoneMap(graph map[string][]string) map[string]bool {
	done := make(map[string]bool)

	var latest string
	for k, v := range graph {
		if k > latest {
			latest = k
		}
		for _, part := range v {
			if part > latest {
				latest = part
			}
		}
	}

	if len(latest) > 1 {
		panic(errors.New("Keys should be single character"))
	}

	for char := 'A'; char <= rune(latest[0]); char++ {
		done[string(char)] = false
	}

	return done
}

func assemblyOrder(graph map[string][]string) string {
	done := generateDoneMap(graph)
	var order string

	for !allDone(done) {
		var available []string
		for part := range done {
			if !done[part] && allDependenciesDone(done, graph[part]) {
				available = append(available, part)
			}
		}

		earliest := available[0]
		for _, a := range available[1:] {
			if a < earliest {
				earliest = a
			}
		}

		done[earliest] = true
		order += earliest
	}

	return order
}

func workedOn(workers []worker, part string) bool {
	for _, worker := range workers {
		if part == worker.part {
			return true
		}
	}

	return false
}

func multipleWorkersAssemblyOrder(graph map[string][]string, workerCount int, timePerPart int) int {
	done := generateDoneMap(graph)
	workers := make([]worker, workerCount)
	var time int

	for !allDone(done) {
		var availableWorkers []int
		for index, worker := range workers {
			if worker.nextFree <= time {
				availableWorkers = append(availableWorkers, index)
				if worker.part != "" {
					done[worker.part] = true
					worker.part = ""
				}
			}
		}

		if len(availableWorkers) == 0 {
			time++
			continue
		}

		var availableParts []string
		for part := range done {
			if !done[part] && allDependenciesDone(done, graph[part]) && !workedOn(workers, part) {
				availableParts = append(availableParts, part)
			}
		}

		sort.Strings(availableParts)

		for _, index := range availableWorkers {
			if len(availableParts) == 0 {
				break
			}

			part := availableParts[0]
			availableParts = availableParts[1:]

			workers[index].part = part
			workers[index].nextFree = time + timePerPart + int(rune(part[0])-'A') + 1
		}

		time++
	}

	return time - 1
}

func parseDependencies() map[string][]string {
	file, err := os.Open("./input07.txt")
	check(err)
	defer file.Close()

	graph := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dependency := scanner.Text()
		var node, dep string
		_, err := fmt.Sscanf(dependency, "Step %s must be finished before step %s can begin.", &dep, &node)
		check(err)

		graph[node] = append(graph[node], dep)
	}
	return graph
}

func main() {
	graph := parseDependencies()

	fmt.Printf("part 1: %s\n", assemblyOrder(graph))
	fmt.Printf("part 2: %d\n", multipleWorkersAssemblyOrder(graph, 4, 60))
}
