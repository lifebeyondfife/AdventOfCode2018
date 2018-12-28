package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node contains a recursive tree structure for child nodes, and integer metadata
type Node struct {
	children []Node
	metadata []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateTree(node *Node, nodeData []int) []int {
	node.children = make([]Node, nodeData[0])
	node.metadata = make([]int, nodeData[1])
	nodeData = nodeData[2:]

	for i := range node.children {
		nodeData = generateTree(&node.children[i], nodeData)
	}

	node.metadata = nodeData[:len(node.metadata)]
	return nodeData[len(node.metadata):]
}

func calculateScore(node Node) int {
	var score int
	for _, child := range node.children {
		score += calculateScore(child)
	}

	for _, metadatum := range node.metadata {
		score += metadatum
	}

	return score
}

func treeScore(numbers []int) int {
	root := Node{}

	_ = generateTree(&root, numbers)

	return calculateScore(root)
}

func calculateNodeValue(node Node) int {
	var score int
	if len(node.children) == 0 {
		for _, metadatum := range node.metadata {
			score += metadatum
		}

		return score
	}

	for _, metadatum := range node.metadata {
		if metadatum > 0 && metadatum <= len(node.children) {
			score += calculateNodeValue(node.children[metadatum-1])
		}
	}

	return score
}

func nodeValue(numbers []int) int {
	root := Node{}

	_ = generateTree(&root, numbers)

	return calculateNodeValue(root)
}

func parseNumbers() []int {
	file, err := os.Open("./input08.txt")
	check(err)
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for _, numberStr := range strings.Split(scanner.Text(), " ") {
		number, err := strconv.Atoi(numberStr)
		check(err)
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	numbers := parseNumbers()

	fmt.Printf("part 1: %d\n", treeScore(numbers))
	fmt.Printf("part 2: %d\n", nodeValue(numbers))
}
