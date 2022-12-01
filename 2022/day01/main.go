package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	data := parseFile()

	var maxCals int
	for _, d := range data {
		var totalCals int
		for _, c := range d {
			totalCals += c
		}
		if totalCals > maxCals {
			maxCals = totalCals
		}
	}

	fmt.Println(maxCals)
}

func problemTwo() {
	data := parseFile()

	var total []int
	for _, d := range data {
		var totalCals int
		for _, c := range d {
			totalCals += c
		}
		total = append(total, totalCals)
	}
	sort.Ints(total)
	topThree := total[len(total)-3:]

	var totalCals int
	for _, x := range topThree {
		totalCals += x
	}
	fmt.Println(totalCals)
}

func parseFile() [][]int {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var data [][]int
	for i := 0; i < len(raw); i++ {
		var cals []int
		for raw[i] != "" {
			val, err := strconv.Atoi(raw[i])
			if err != nil {
				log.Fatalln(err)
			}
			cals = append(cals, val)
			i++
		}
		data = append(data, cals)
	}
	return data
}
