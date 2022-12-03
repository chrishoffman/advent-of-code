package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/chrishoffman/advent-of-code"
)

type rucksack struct{ comp1, comp2 string }

func (r rucksack) concat() string {
	return fmt.Sprintf("%s%s", r.comp1, r.comp2)
}

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	rucksacks := parseFile("./input.txt")
	fmt.Println(totalPriority(rucksacks))
}

func problemTwo() {
	rucksacks := parseFile("./input.txt")
	priorities := generatePriorityMap()

	var totalPriority int
	for i := 0; i < len(rucksacks); i += 3 {
		unique := make(map[rune]map[int]bool)
		for j := 0; j < 3; j++ {
			for _, r := range rucksacks[j+i].concat() {
				if _, ok := unique[r]; !ok {
					unique[r] = make(map[int]bool)
				}
				unique[r][j] = true
			}
		}
		for r := range unique {
			if len(unique[r]) == 3 {
				totalPriority += priorities[r]
			}
		}
	}
	fmt.Println(totalPriority)
}

func totalPriority(rucksacks []rucksack) int {
	var totalPriority int
	priorities := generatePriorityMap()
	for _, r := range rucksacks {
		common := commonItems(r)
		for _, c := range common {
			totalPriority += priorities[c]
		}
	}
	return totalPriority
}

func commonItems(r rucksack) []rune {
	unique := make(map[rune]int)
	for _, l := range r.comp1 {
		unique[l]++
	}

	var common []rune
	found := make(map[rune]int)
	for _, l := range r.comp2 {
		if _, ok := unique[l]; ok {
			if _, ok := found[l]; ok {
				continue
			}
			common = append(common, l)
			found[l]++
		}
	}

	return common
}

func generatePriorityMap() map[rune]int {
	currPriority := 1
	priorities := make(map[rune]int)
	for r := 'a'; r <= 'z'; r++ {
		priorities[r] = currPriority

		upper := unicode.ToUpper(r)
		priorities[upper] = currPriority + 26

		currPriority++
	}
	return priorities
}

func parseFile(file string) []rucksack {
	raw, err := advent.ParseFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	var rucksacks []rucksack
	for _, l := range raw {
		midpoint := len(l) / 2
		rucksacks = append(rucksacks, rucksack{l[:midpoint], l[midpoint:]})
	}
	return rucksacks
}
