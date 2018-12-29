package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseGame() (int, int) {
	file, err := os.Open("./input09.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	gameStr := scanner.Text()
	var playerCount, marbleCount int
	_, err = fmt.Sscanf(gameStr, "%d players; last marble is worth %d points", &playerCount, &marbleCount)
	check(err)

	return playerCount, marbleCount
}

func newMarble(value int) *ring.Ring {
	marble := ring.New(1)
	marble.Value = value
	return marble
}

func playGame(playerCount, marbleCount int) int {
	ring := ring.New(1)
	ring.Value = 0

	players := make([]int, playerCount)
	var player int
	for i := 1; i <= marbleCount; i++ {
		if i%23 != 0 {
			ring.Move(2)
			ring = ring.Link(newMarble(i))
		} else {
			ring = ring.Move(-8)
			players[player] += i + ring.Value.(int)
			ring = ring.Prev()
			ring.Unlink(1)
			ring = ring.Move(2)
		}

		player = (player + 1) % len(players)
	}

	var winningScore int
	for _, score := range players {
		if score > winningScore {
			winningScore = score
		}
	}

	return winningScore
}

func main() {
	playerCount, marbleCount := parseGame()

	fmt.Printf("part 1: %d\n", playGame(playerCount, marbleCount))
	fmt.Printf("part 2: %d\n", playGame(playerCount, marbleCount*100))
}
