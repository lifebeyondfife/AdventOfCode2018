package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	shift = iota
	awake = iota
	sleep = iota
)

type guardEvent struct {
	datetime time.Time
	action   int
	guard    int
}

func parseGuardEventStr(logStr string) guardEvent {
	var y, m, d, h, min int
	n, err := fmt.Sscanf(logStr, "[%d-%d-%d %d:%d]", &y, &m, &d, &h, &min)
	check(err)

	datetime := time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)

	actionStr := logStr[19:]
	var guard int
	n, err = fmt.Sscanf(actionStr, "Guard #%d begins shift", &guard)

	if n == 1 {
		return guardEvent{datetime, shift, guard}
	}
	if actionStr == "wakes up" {
		return guardEvent{datetime, awake, math.MinInt32}
	}
	return guardEvent{datetime, sleep, math.MinInt32}
}

func sleepiestGuardTime(logStrs []string) int {
	var guardEvents []guardEvent
	for _, guardEventStr := range logStrs {
		guardEvents = append(guardEvents, parseGuardEventStr(guardEventStr))
	}

	sort.Slice(guardEvents, func(i, j int) bool {
		return guardEvents[i].datetime.Before(guardEvents[j].datetime)
	})

	sleepiestGuard := findSleepiestGuard(guardEvents)
	bestMinute := findBestMinute(sleepiestGuard, guardEvents)

	return sleepiestGuard * bestMinute
}

func findSleepiestGuard(guardEvents []guardEvent) int {
	guardSleepMinutes := make(map[int]int)
	var asleep time.Time
	var guard int
	for _, guardEvent := range guardEvents {
		if guardEvent.action == shift {
			guard = guardEvent.guard
			continue
		}

		if guardEvent.action == sleep {
			asleep = guardEvent.datetime
			continue
		}

		guardSleepMinutes[guard] += int(guardEvent.datetime.Sub(asleep).Minutes())
	}

	sleepiestGuard := math.MinInt32
	for k, v := range guardSleepMinutes {
		if v > guardSleepMinutes[sleepiestGuard] {
			sleepiestGuard = k
		}
	}

	return sleepiestGuard
}

func findBestMinute(sleepiestGuard int, guardEvents []guardEvent) int {
	guardMinuteAsleep := make(map[int]int)
	var asleep time.Time
	var guard int
	for _, guardEvent := range guardEvents {
		if guardEvent.action == shift {
			guard = guardEvent.guard
			continue
		}

		if guard != sleepiestGuard {
			continue
		}

		if guardEvent.action == sleep {
			asleep = guardEvent.datetime
			continue
		}

		for i := asleep.Minute(); i < asleep.Minute()+int(guardEvent.datetime.Sub(asleep).Minutes()); i++ {
			guardMinuteAsleep[i]++
		}
	}

	sleepiestMinute := math.MinInt32
	for k, v := range guardMinuteAsleep {
		if v > guardMinuteAsleep[sleepiestMinute] {
			sleepiestMinute = k
		}
	}

	return sleepiestMinute
}

func mostFrequentGuardSleepTime(logStrs []string) int {
	var guardEvents []guardEvent
	for _, guardEventStr := range logStrs {
		guardEvents = append(guardEvents, parseGuardEventStr(guardEventStr))
	}

	sort.Slice(guardEvents, func(i, j int) bool {
		return guardEvents[i].datetime.Before(guardEvents[j].datetime)
	})

	return mostFrequentGuardSleepTimeEvents(guardEvents)
}

func mostFrequentGuardSleepTimeEvents(guardEvents []guardEvent) int {
	var guardMinuteAsleep = map[int]map[int]int{}
	var asleep time.Time
	var guard int
	for _, guardEvent := range guardEvents {
		if guardEvent.action == shift {
			guard = guardEvent.guard

			if guardMinuteAsleep[guard] == nil {
				guardMinuteAsleep[guard] = make(map[int]int)
			}

			continue
		}

		if guardEvent.action == sleep {
			asleep = guardEvent.datetime
			continue
		}

		for i := asleep.Minute(); i < asleep.Minute()+int(guardEvent.datetime.Sub(asleep).Minutes()); i++ {
			guardMinuteAsleep[guard][i]++
		}
	}

	globalSleepiestMinute := math.MinInt32
	var mostFrequestAsleepGuard int
	for k, v := range guardMinuteAsleep {
		if mostFrequestAsleepGuard == 0 {
			mostFrequestAsleepGuard = k
		}

		localSleepiestMinute := math.MinInt32
		for min, count := range v {
			if count > v[localSleepiestMinute] {
				localSleepiestMinute = min
			}
		}

		if v[localSleepiestMinute] > guardMinuteAsleep[mostFrequestAsleepGuard][globalSleepiestMinute] {
			globalSleepiestMinute = localSleepiestMinute
			mostFrequestAsleepGuard = k
		}
	}

	return globalSleepiestMinute * mostFrequestAsleepGuard
}

func parseLogs(file *os.File) []string {
	var logStrs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logStrs = append(logStrs, scanner.Text())
	}
	return logStrs
}

func main() {
	file, err := os.Open("./input04.txt")
	check(err)
	defer file.Close()

	logStrs := parseLogs(file)

	fmt.Printf("part 1: %d\n", sleepiestGuardTime(logStrs))
	fmt.Printf("part 2: %d\n", mostFrequentGuardSleepTime(logStrs))
}
