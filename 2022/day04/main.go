package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type sectionRange struct{ start, end int }
type elfPair struct{ elf1, elf2 sectionRange }

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	pairs := parseFile()

	var count int
	for _, p := range pairs {
		if (p.elf1.start <= p.elf2.start && p.elf1.end >= p.elf2.end) ||
			(p.elf2.start <= p.elf1.start && p.elf2.end >= p.elf1.end) {
			count++
		}
	}
	fmt.Println(count)
}

func problemTwo() {
	pairs := parseFile()

	var count int
	for _, p := range pairs {
		if !(p.elf1.end < p.elf2.start || p.elf2.end < p.elf1.start) {
			count++
		}
	}
	fmt.Println(count)
}

func parseFile() []elfPair {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var elves []elfPair
	for _, r := range raw {
		ranges := strings.Split(r, ",")
		var sects []sectionRange
		for _, sect := range ranges {
			startEnd := strings.Split(sect, "-")
			start, _ := strconv.Atoi(startEnd[0])
			end, _ := strconv.Atoi(startEnd[1])
			sects = append(sects, sectionRange{start, end})
		}
		elves = append(elves, elfPair{sects[0], sects[1]})
	}
	return elves
}
