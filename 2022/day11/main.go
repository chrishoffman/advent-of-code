package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	run(20, func(v int) int { return v / 3 })
}

func problemTwo() {
	run(10000, func(v int) int { return v })
}

func run(rounds int, f func(int) int) {
	monkeys, lcm := parseFile()

	inspections := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range monkey.items {
				new := f(monkey.operation(item) % lcm)
				dest := monkey.action[monkey.test(new)]

				monkeys[dest].items = append(monkeys[dest].items, new)
				monkey.items = monkey.items[1:len(monkey.items)]
				inspections[m]++
			}
		}
	}

	sort.Ints(inspections)
	fmt.Println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}

type monkey struct {
	items     []int
	test      func(int) bool
	operation func(int) int
	action    map[bool]int
}

func parseFile() ([]*monkey, int) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var monkeys []*monkey
	lcm := 1
	for i := 1; i < len(raw); i += 7 {
		m := new(monkey)

		levelsRaw := strings.Split(strings.TrimPrefix(raw[i], "  Starting items: "), ", ")
		for _, l := range levelsRaw {
			lev, _ := strconv.Atoi(l)
			m.items = append(m.items, lev)
		}

		opRaw := strings.Fields(strings.TrimPrefix(raw[i+1], "  Operation: new = old "))
		increment, _ := strconv.Atoi(opRaw[1])
		switch opRaw[0] {
		case "+":
			m.operation = func(v int) int { return v + increment }
		case "*":
			if increment == 0 {
				m.operation = func(v int) int { return v * v }
			} else {
				m.operation = func(v int) int { return v * increment }
			}
		}

		testRaw := strings.TrimPrefix(raw[i+2], "  Test: divisible by ")
		divisor, _ := strconv.Atoi(testRaw)
		lcm *= divisor
		m.test = func(v int) bool { return v%divisor == 0 }

		trueRaw := strings.TrimPrefix(raw[i+3], "    If true: throw to monkey ")
		falseRaw := strings.TrimPrefix(raw[i+4], "    If false: throw to monkey ")
		trueDest, _ := strconv.Atoi(trueRaw)
		falseDest, _ := strconv.Atoi(falseRaw)
		m.action = map[bool]int{
			true:  trueDest,
			false: falseDest,
		}

		monkeys = append(monkeys, m)
	}
	return monkeys, lcm
}
