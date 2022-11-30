package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	values := parseFile()

	increaseCount := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
}

func problemTwo() {
	values := parseFile()

	increaseCount := 0
	for i := 1; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		prevSum := values[i-1] + values[i] + values[i+1]
		if sum > prevSum {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
}

func parseFile() []int {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var values []int
	for _, v := range raw {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		values = append(values, value)
	}

	return values
}
