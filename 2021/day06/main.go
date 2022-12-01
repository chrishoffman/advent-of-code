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
	// timers := parseFile()

	// for i := 0; i < 80; i++ {
	// 	for t := 0; t < len(timers); t++ {
	// 		switch timers[t] {
	// 		case 0:
	// 			timers[t] = 6
	// 			timers = append(timers, 9)
	// 		default:
	// 			timers[t]--
	// 		}
	// 	}
	// }

	// fmt.Println(len(timers))
}

func problemTwo() {
	timers := parseFile()

	newTimers := map[int]int{}
	for day := 0; day < 256; day++ {
		for {
			if days == 0 {
				newTimers[6] += count

			}
		}
	}

	total :=
		fmt.Println(len(timers))
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
