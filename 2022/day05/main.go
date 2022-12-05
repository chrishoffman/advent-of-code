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
	stacks, ops := parseFile()

	for _, op := range ops {
		sIndex := op.start - 1
		eIndex := op.end - 1

		var items []string
		for i := op.count - 1; i >= 0; i-- {
			items = append(items, stacks[sIndex][i])
		}
		stacks[eIndex] = append(items, stacks[eIndex]...)
		stacks[sIndex] = stacks[sIndex][op.count:]
	}

	var final string
	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			final += stacks[i][0]
		}
	}
	fmt.Println(final)
}

func problemTwo() {
	stacks, ops := parseFile()

	for _, op := range ops {
		sIndex := op.start - 1
		eIndex := op.end - 1

		newStart := make([]string, len(stacks[sIndex]))
		copy(newStart, stacks[sIndex][op.count:])

		stacks[eIndex] = append(stacks[sIndex][0:op.count], stacks[eIndex]...)
		stacks[sIndex] = newStart
	}

	var final string
	for i := 0; i < len(stacks); i++ {
		if len(stacks[i]) > 0 {
			final += stacks[i][0]
		}
	}
	fmt.Println(final)
}

type operation struct{ count, start, end int }

func parseFile() ([][]string, []operation) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	stacks := make([][]string, 9)
	opStart := 0
	for _, r := range raw {
		opStart++
		if r == "" {
			break
		}

		stack := 0
		for i := 1; i < len(r); i += 4 {
			item := string(r[i])
			if item == "1" {
				break
			}
			if item != " " {
				stacks[stack] = append(stacks[stack], item)
			}
			stack++
		}
	}

	var ops []operation
	for r := opStart; r < len(raw); r++ {
		parsed := strings.Fields(raw[r])
		count, _ := strconv.Atoi(parsed[1])
		start, _ := strconv.Atoi(parsed[3])
		end, _ := strconv.Atoi(parsed[5])
		ops = append(ops, operation{count, start, end})
	}

	return stacks, ops
}
