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

var cycleOps = map[string]int{
	"addx": 2,
	"noop": 1,
}

func problemOne() {
	insts := parseFile()

	var sigSum, cycle int
	x := 1
	for _, i := range insts {
		for c := 0; c < cycleOps[i.op]; c++ {
			cycle++

			if cycle <= 220 && (cycle-20)%40 == 0 {
				sigSum += cycle * x
			}
		}

		if i.op == "addx" {
			x += i.v
		}
	}
	fmt.Println(sigSum)
}

func problemTwo() {
	insts := parseFile()

	var cycle int
	x := 1
	for _, i := range insts {
		for c := 0; c < cycleOps[i.op]; c++ {
			cycle++

			timing := x + 2 - cycle
			if timing >= 0 && timing < 3 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			if cycle%40 == 0 {
				cycle -= 40
				fmt.Print("\n")
			}
		}

		if i.op == "addx" {
			x += i.v
		}
	}
}

type instruction struct {
	op string
	v  int
}

func parseFile() []instruction {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var insts []instruction
	for _, r := range raw {
		fields := strings.Fields(r)
		inst := instruction{op: fields[0]}
		if len(fields) > 1 {
			inst.v, _ = strconv.Atoi(fields[1])
		}
		insts = append(insts, inst)
	}
	return insts
}
