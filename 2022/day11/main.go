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
	monkeys, _ := parseFile()

	inspections := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range monkey.items {
				new := monkey.operatiion(item) / uint64(3)
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

func problemTwo() {
	monkeys, lcm := parseFile()

	inspections := make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range monkey.items {
				new := monkey.operatiion(item) % uint64(lcm)
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
	items      []uint64
	test       func(uint64) bool
	operatiion func(uint64) uint64
	action     map[bool]int
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
			m.items = append(m.items, uint64(lev))
		}

		opRaw := strings.Fields(strings.TrimPrefix(raw[i+1], "  Operation: new = old "))
		increment, _ := strconv.Atoi(opRaw[1])
		switch opRaw[0] {
		case "+":
			m.operatiion = func(v uint64) uint64 { return v + uint64(increment) }
		case "*":
			if increment == 0 {
				m.operatiion = func(v uint64) uint64 { return v * v }
			} else {
				m.operatiion = func(v uint64) uint64 { return v * uint64(increment) }
			}
		}

		testRaw := strings.TrimPrefix(raw[i+2], "  Test: divisible by ")
		divisor, _ := strconv.Atoi(testRaw)
		lcm *= divisor
		m.test = func(v uint64) bool { return v%uint64(divisor) == 0 }

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
