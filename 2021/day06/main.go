package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	fmt.Println(elapseDays(80))
}

func problemTwo() {
	fmt.Println(elapseDays(256))
}

func elapseDays(n int) int {
	timers := parseFile()

	var count int
	for i := 0; i < n; i++ {
		count = 0
		newTimers := map[int]int{}
		for t, fish := range timers {
			switch t {
			case 0:
				newTimers[6] += fish
				newTimers[8] += fish
				count += fish
			default:
				newTimers[t-1] += fish
			}
			count += fish
		}
		timers = newTimers
	}
	return count
}

func parseFile() map[int]int {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	list := strings.Split(raw[0], ",")
	final := map[int]int{}
	for _, v := range list {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		final[val]++
	}

	return final
}
